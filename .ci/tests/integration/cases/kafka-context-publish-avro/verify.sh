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
E2E_KUBECONFIG=${E2E_KUBECONFIG:-"/tmp/e2e-k8s.config"}

source "${BASE_DIR}"/.ci/helm.sh

if [ ! "$KUBECONFIG" ]; then
  export KUBECONFIG=${E2E_KUBECONFIG}
fi

manifests_file="${BASE_DIR}"/.ci/tests/integration/cases/kafka-context-publish-avro/manifests.yaml
avro_schema='{"type":"record","name":"Student","fields":[{"name":"name","type":["null","string"]},{"name":"age","type":["null","int"]},{"name":"grade","type":["null","int"]}]}'

if ! reset_topics_result=$(ci::kafka_reset_topics kafka-context-publish-avro-input kafka-context-publish-avro-output 2>&1); then
  echo "$reset_topics_result"
  exit 1
fi

if ! apply_result=$(ci::apply_manifest_or_debug "${manifests_file}" 2>&1); then
  echo "$apply_result"
  exit 1
fi

if ! verify_fm_result=$(ci::verify_kafka_function_mesh kafka-context-publish-avro-generic-sample 2>&1); then
  echo "$verify_fm_result"
  kubectl delete -f "${manifests_file}" > /dev/null 2>&1 || true
  exit 1
fi

produce_command=$(cat <<EOF
cat > /tmp/value.avsc <<'AVSC'
${avro_schema}
AVSC
cat > /tmp/message.json <<'JSON'
{"name":{"string":"jack"},"age":{"int":21},"grade":{"int":80}}
JSON
kafka-avro-console-producer --bootstrap-server my-kafka:9092 --topic kafka-context-publish-avro-input --property schema.registry.url=http://my-schema-registry:8081 --property value.schema.file=/tmp/value.avsc < /tmp/message.json
EOF
)
if ! produce_result=$(ci::kafka_client "${produce_command}" 2>&1); then
  echo "$produce_result"
  kubectl delete -f "${manifests_file}" > /dev/null 2>&1 || true
  exit 1
fi

consume_command='timeout 60 kafka-avro-console-consumer --bootstrap-server my-kafka:9092 --topic kafka-context-publish-avro-output --from-beginning --max-messages 1 --property schema.registry.url=http://my-schema-registry:8081'
if ! consume_result=$(ci::kafka_client "${consume_command}" 2>&1); then
  echo "$consume_result"
  kubectl delete -f "${manifests_file}" > /dev/null 2>&1 || true
  exit 1
fi
if [[ "$consume_result" != *'"name":{"string":"jack"}'* ]] || [[ "$consume_result" != *'"age":{"int":21}'* ]] || [[ "$consume_result" != *'"grade":{"int":81}'* ]]; then
  echo "$consume_result"
  kubectl delete -f "${manifests_file}" > /dev/null 2>&1 || true
  exit 1
fi

if ! schema_result=$(ci::schema_registry_get_subject "kafka-context-publish-avro-output-value" 2>&1); then
  echo "$schema_result"
  kubectl delete -f "${manifests_file}" > /dev/null 2>&1 || true
  exit 1
fi

if ! verify_total_processed=$(ci::verify_total_processed kafka-context-publish-avro-generic-sample 1); then
  echo "$verify_total_processed"
  kubectl delete -f "${manifests_file}" > /dev/null 2>&1 || true
  exit 1
fi

echo "e2e-test: ok" | yq eval -
kubectl delete -f "${manifests_file}" > /dev/null 2>&1 || true
