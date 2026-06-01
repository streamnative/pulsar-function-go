#!/usr/bin/env bash
#
# Licensed to the Apache Software Foundation (ASF) under one
# or more contributor license agreements.  See the NOTICE file
# distributed with this work for additional information
# regarding copyright ownership.  The ASF licenses this file
# to you under the Apache License, Version 2.0 (the
# "License"); you may not use this file except in compliance
# with the License.  You may obtain a copy of the License at
#
#   http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing,
# software distributed under the License is distributed on an
# "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
# KIND, either express or implied.  See the License for the
# specific language governing permissions and limitations
# under the License.
#

set -ex

BINDIR=`dirname "$0"`
PULSAR_HOME=`cd ${BINDIR}/..;pwd`
FUNCTION_MESH_HOME=${PULSAR_HOME}
OUTPUT_BIN=${FUNCTION_MESH_HOME}/output/bin
KIND_BIN=$OUTPUT_BIN/kind
HELM=${OUTPUT_BIN}/helm
KUBECTL=${OUTPUT_BIN}/kubectl
NAMESPACE=default
CLUSTER=sn-platform
CLUSTER_ID=$(uuidgen | tr "[:upper:]" "[:lower:]")

function ci::ensure_topic_with_schema() {
    topic=$1
    schema=$2
    if [[ "$schema" == "json" ]]; then
      MESSAGE=$(kubectl exec -n ${NAMESPACE} ${CLUSTER}-pulsar-broker-0 -- python3 examples/consume_json.py "${topic}")
    fi

    if [[ "$schema" == "avro" ]]; then
      MESSAGE=$(kubectl exec -n ${NAMESPACE} ${CLUSTER}-pulsar-broker-0 -- python3 examples/consume_avro.py "${topic}")
    fi
}

function ci::verify_function_mesh() {
    FUNCTION_NAME=$1

    num=0
    while [[ ${num} -lt 1 ]]; do
        sleep 5
        kubectl get pods
        num=$(kubectl get pods -l compute.functionmesh.io/name="${FUNCTION_NAME}" | wc -l)
    done

    kubectl wait -l compute.functionmesh.io/name="${FUNCTION_NAME}" --for=condition=Ready pod --timeout=2m && true

    num=0
    while [[ ${num} -lt 1 ]]; do
        sleep 5
        kubectl get pods -l compute.functionmesh.io/name="${FUNCTION_NAME}"
        kubectl logs -l compute.functionmesh.io/name="${FUNCTION_NAME}" --tail=50 || true
        num=$(kubectl logs -l compute.functionmesh.io/name="${FUNCTION_NAME}" --tail=-1 | grep "to pulsar" | wc -l)
    done
}

function ci::produce_message() {
    topic=$1
    message=$2
    kubectl exec -n ${NAMESPACE} ${CLUSTER}-pulsar-broker-0 -- bin/pulsar-client produce -m "${message}" -n 1 "${topic}"
}

function ci::verify_exclamation_function() {
    inputtopic=$1
    outputtopic=$2
    inputmessage=$3
    outputmessage=$4
    timesleep=$5
    ci::produce_message "${inputtopic}" "${inputmessage}"
    sleep "$timesleep"
    MESSAGE=$(kubectl exec -n ${NAMESPACE} ${CLUSTER}-pulsar-broker-0 -- bin/pulsar-client consume -n 1 -s "sub" --subscription-position Earliest "${outputtopic}")
    echo "$MESSAGE"
    if [[ "$MESSAGE" == *"$outputmessage"* ]]; then
        return 0
    fi
    return 1
}

function ci::verify_exclamation_function_with_auth() {
    inputtopic=$1
    outputtopic=$2
    inputmessage=$3
    outputmessage=$4
    timesleep=$5
    command="kubectl exec -n ${NAMESPACE} ${CLUSTER}-pulsar-broker-0 -- sh -c 'bin/pulsar-client --auth-plugin \$brokerClientAuthenticationPlugin --auth-params \$brokerClientAuthenticationParameters produce -m \"${inputmessage}\" -n 1 \"${inputtopic}\"'"
    sh -c "$command"
    sleep "$timesleep"
    consumeCommand="kubectl exec -n ${NAMESPACE} ${CLUSTER}-pulsar-broker-0 -- sh -c 'bin/pulsar-client --auth-plugin \$brokerClientAuthenticationPlugin --auth-params \$brokerClientAuthenticationParameters consume -n 1 -s "sub" --subscription-position Earliest \"${outputtopic}\"'"
    MESSAGE=$(sh -c "$consumeCommand")
    echo "$MESSAGE"
    if [[ "$MESSAGE" == *"$outputmessage"* ]]; then
        return 0
    fi
    return 1
}

function ci::verify_json_function() {
    inputtopic=$1
    outputtopic=$2
    name=$3
    age=$4
    grade=$5
    outputmessage=$6
    timesleep=$7
    kubectl exec -n ${NAMESPACE} ${CLUSTER}-pulsar-broker-0 -- python3 examples/produce_json.py "${inputtopic}" "${name}" "${age}" "${grade}"
    sleep "$timesleep"
    MESSAGE=$(kubectl exec -n ${NAMESPACE} ${CLUSTER}-pulsar-broker-0 -- python3 examples/consume_json.py "${outputtopic}")
    echo "$MESSAGE"
    if [[ "$MESSAGE" == *"$outputmessage"* ]]; then
        return 0
    fi
    return 1
}

function ci::verify_avro_function() {
    inputtopic=$1
    outputtopic=$2
    name=$3
    age=$4
    grade=$5
    outputmessage=$6
    timesleep=$7
    kubectl exec -n ${NAMESPACE} ${CLUSTER}-pulsar-broker-0 -- python3 examples/produce_avro.py "${inputtopic}" "${name}" "${age}" "${grade}"
    sleep "$timesleep"
    MESSAGE=$(kubectl exec -n ${NAMESPACE} ${CLUSTER}-pulsar-broker-0 -- python3 examples/consume_avro.py "${outputtopic}")
    echo "$MESSAGE"
    if [[ "$MESSAGE" == *"$outputmessage"* ]]; then
        return 0
    fi
    return 1
}

function ci::verify_topic() {
    topic=$1
    message=$2
    MESSAGE=$(kubectl exec -n ${NAMESPACE} ${CLUSTER}-pulsar-broker-0 -- bin/pulsar-client consume -n 1 -s "sub" --subscription-position Earliest "${topic}")
    echo "$MESSAGE"
    if [[ "$MESSAGE" == *"$message"* ]]; then
        return 0
    fi
    return 1
}

function ci::verify_topic_with_auth() {
    topic=$1
    message=$2
    consumeCommand="kubectl exec -n ${NAMESPACE} ${CLUSTER}-pulsar-broker-0 -- sh -c 'bin/pulsar-client --auth-plugin \$brokerClientAuthenticationPlugin --auth-params \$brokerClientAuthenticationParameters consume -n 1 -s "sub" --subscription-position Earliest \"${topic}\"'"
    MESSAGE=$(sh -c "$consumeCommand")
    echo "$MESSAGE"
    if [[ "$MESSAGE" == *"$message"* ]]; then
        return 0
    fi
    return 1
}

function ci::verify_total_processed() {
  path="$1-function-headless:9094/metrics"
  total=$2
  curlCommand="kubectl exec -n ${NAMESPACE} ${CLUSTER}-pulsar-broker-0 -- sh -c 'curl \"${path}\"' | grep pulsar_function_processed_successfully_total | grep -v \"#\""
  MESSAGE=$(sh -c "$curlCommand")
  echo "$MESSAGE"
  if [[ "$MESSAGE" == *"$total" ]]; then
      return 0
  fi
  return 1
}

function ci::create_topic() {
  topic=$1
  partition=${2:-"1"}
  kubectl exec -n ${NAMESPACE} ${CLUSTER}-pulsar-broker-0 -- bin/pulsar-admin topics create-partitioned-topic ${topic} -p ${partition} || true
}
