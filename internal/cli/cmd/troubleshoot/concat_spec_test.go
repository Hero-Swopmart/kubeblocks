/*
Copyright ApeCloud, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package troubleshoot

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	troubleshootv1beta2 "github.com/replicatedhq/troubleshoot/pkg/apis/troubleshoot/v1beta2"
	"k8s.io/apimachinery/pkg/util/yaml"
)

var _ = Describe("Concat Spec Test", func() {

	It("ConcatPreflightSpec Test", func() {
		targetByte := `
apiVersion: troubleshoot.sh/v1beta2
kind: Preflight
metadata:
  name: sample
spec:
  analyzers:
    - nodeResources:
        checkName: Must have at least 3 nodes in the cluster
        outcomes:
          - fail:
              when: "< 3"
              message: This application requires at least 3 nodes
          - warn:
              when: "< 5"
              message: This application recommends at last 5 nodes.
          - pass:
              message: This cluster has enough nodes.`
		sourceByte := `
apiVersion: troubleshoot.sh/v1beta2
kind: Preflight
metadata:
  name: sample
spec:
  collectors:
    - redis:
        collectorName: my-redis
        uri: rediss://default:replicated@server:6380
        tls:
          skipVerify: true
  analyzers:
    - redis:
        checkName: Must be redis 5.x or later
        collectorName: my-redis
        outcomes:
          - fail:
              when: "connected == false"
              message: Cannot connect to redis server
          - fail:
              when: "version < 5.0.0"
              message: The redis server must be at least version 5
          - pass:
              message: The redis connection checks out.`
		targetSpec := new(troubleshootv1beta2.Preflight)
		sourceSpec := new(troubleshootv1beta2.Preflight)
		Expect(yaml.Unmarshal([]byte(targetByte), targetSpec)).Should(Succeed())
		Expect(yaml.Unmarshal([]byte(sourceByte), sourceSpec)).Should(Succeed())
		var newSpec = ConcatPreflightSpec(nil, sourceSpec)
		Expect(newSpec).Should(Equal(sourceSpec))
		newSpec = ConcatPreflightSpec(targetSpec, nil)
		Expect(newSpec).Should(Equal(targetSpec))
		newSpec = ConcatPreflightSpec(targetSpec, sourceSpec)
		Expect(len(newSpec.Spec.Analyzers)).Should(Equal(2))
	})
	It("ConcatHostPreflightSpec Test", func() {
		targetByte := `
apiVersion: troubleshoot.sh/v1beta2
kind: HostPreflight
metadata:
  name: cpu
spec:
  collectors:
    - cpu: {}
  analyzers:
    - cpu:
        outcomes:
          - fail:
              when: "physical < 4"
              message: At least 4 physical CPU cores are required
          - fail:
              when: "logical < 8"
              message: At least 8 CPU cores are required
          - warn:
              when: "count < 16"
              message: At least 16 CPU cores preferred
          - pass:
              message: This server has sufficient CPU cores.`
		sourceByte := `
apiVersion: troubleshoot.sh/v1beta2
kind: HostPreflight
metadata:
  name: http
spec:
  collectors:
    - http:
        collectorName: registry
        get:
          url: https://registry.replicated.com
  analyzers:
    - http:
        collectorName: registry
        outcomes:
          - fail:
              when: "error"
              message: Error connecting to registry
          - pass:
              when: "statusCode == 404"
              message: Connected to registry
          - fail:
              message: "Unexpected response"`
		targetSpec := new(troubleshootv1beta2.HostPreflight)
		sourceSpec := new(troubleshootv1beta2.HostPreflight)
		Expect(yaml.Unmarshal([]byte(targetByte), targetSpec)).Should(Succeed())
		Expect(yaml.Unmarshal([]byte(sourceByte), sourceSpec)).Should(Succeed())
		var newSpec = ConcatHostPreflightSpec(nil, sourceSpec)
		Expect(newSpec).Should(Equal(sourceSpec))
		newSpec = ConcatHostPreflightSpec(targetSpec, nil)
		Expect(newSpec).Should(Equal(targetSpec))
		newSpec = ConcatHostPreflightSpec(targetSpec, sourceSpec)
		Expect(len(newSpec.Spec.Analyzers)).Should(Equal(2))
	})
})
