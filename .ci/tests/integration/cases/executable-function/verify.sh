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

manifests_file="${BASE_DIR}"/.ci/tests/integration/cases/executable-function/manifests.yaml

kubectl apply -f "${manifests_file}" > /dev/null 2>&1

verify_fm_result=$(ci::verify_function_mesh exec-function-excla-generic-sample 2>&1)
if [ $? -ne 0 ]; then
  echo "$verify_fm_result"
  kubectl delete -f "${manifests_file}" > /dev/null 2>&1 || true
  exit 1
fi

verify_executable_result=$(NAMESPACE=${PULSAR_NAMESPACE} CLUSTER=${PULSAR_RELEASE_NAME} ci::verify_exclamation_function "persistent://public/default/exec-generic-excla-input" "persistent://public/default/exec-generic-excla-output" "test-message" "test-message!" 10 2>&1)
if [ $? -ne 0 ]; then
  echo "$verify_executable_result"
  kubectl delete -f "${manifests_file}" > /dev/null 2>&1 || true
  exit 1
fi

# verify total processed
verify_total_processed=$(ci::verify_total_processed exec-function-excla-generic-sample 1)
if [ $? -ne 0 ]; then
  echo "$verify_total_processed"
  kubectl delete -f "${manifests_file}" > /dev/null 2>&1 || true
  exit 1
fi

# verify log topic and state store
verify_log_topic=$(ci::verify_topic "persistent://public/default/exec-generic-excla-logs" "got word: test-message for 1 times" 2>&1)
if [ $? -ne 0 ]; then
  echo "$verify_log_topic"
  kubectl delete -f "${manifests_file}" > /dev/null 2>&1 || true
  exit 1
fi

# verify publish
verify_extra_out=$(ci::verify_topic "persistent://public/default/test-exec-package-serde-extra" "config: cfg-value, secret: sec-value" 2>&1)
if [ $? -ne 0 ]; then
  echo "$verify_extra_out"
  kubectl delete -f "${manifests_file}" > /dev/null 2>&1 || true
  exit 1
fi

echo "e2e-test: ok" | yq eval -
kubectl delete -f "${manifests_file}" > /dev/null 2>&1 || true
