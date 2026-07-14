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

set -e

E2E_DIR=$(dirname "$0")
BASE_DIR=$(cd "${E2E_DIR}"/../../../../..;pwd)
PULSAR_NAMESPACE=${PULSAR_NAMESPACE:-"default"}
PULSAR_RELEASE_NAME=${PULSAR_RELEASE_NAME:-"sn-platform"}
E2E_KUBECONFIG=${E2E_KUBECONFIG:-"/tmp/e2e-k8s.config"}

source "${BASE_DIR}"/.ci/helm.sh

if [ ! "$KUBECONFIG" ]; then
  export KUBECONFIG=${E2E_KUBECONFIG}
fi

manifests_file="${BASE_DIR}"/.ci/tests/integration/cases/pulsar-context-publish-avro/manifests.yaml

kubectl delete -f "${manifests_file}" > /dev/null 2>&1 || true
ci::pulsar_reset_topics \
  "persistent://public/default/pulsar-context-publish-avro-input" \
  "persistent://public/default/pulsar-context-publish-avro-output" \
  "persistent://public/default/pulsar-context-publish-avro-unused-output"

kubectl apply -f "${manifests_file}" > /dev/null 2>&1

if ! verify_fm_result=$(ci::verify_function_mesh pulsar-context-publish-avro-generic-sample 2>&1); then
  echo "$verify_fm_result"
  kubectl delete -f "${manifests_file}" > /dev/null 2>&1 || true
  exit 1
fi

binary_safe_name=$'jack/slash\nline\rcarriage'
if ! verify_go_result=$(NAMESPACE=${PULSAR_NAMESPACE} CLUSTER=${PULSAR_RELEASE_NAME} ci::verify_avro_function "persistent://public/default/pulsar-context-publish-avro-input" "persistent://public/default/pulsar-context-publish-avro-output" "${binary_safe_name}" 21 80 "schema message matches" 10 2>&1); then
  echo "$verify_go_result"
  kubectl delete -f "${manifests_file}" > /dev/null 2>&1 || true
  exit 1
fi

if ! verify_output_schema=$(ci::verify_topic_schema "persistent://public/default/pulsar-context-publish-avro-output" "avro" 2>&1); then
  echo "$verify_output_schema"
  kubectl delete -f "${manifests_file}" > /dev/null 2>&1 || true
  exit 1
fi

if ! verify_total_processed=$(ci::verify_total_processed pulsar-context-publish-avro-generic-sample 1); then
  echo "$verify_total_processed"
  kubectl delete -f "${manifests_file}" > /dev/null 2>&1 || true
  exit 1
fi

echo "e2e-test: ok" | yq eval -
kubectl delete -f "${manifests_file}" > /dev/null 2>&1 || true
