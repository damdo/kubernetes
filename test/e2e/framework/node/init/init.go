/*
Copyright 2022 The Kubernetes Authors.

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

// Package init registers node.AllNodesReady.
package init

import (
	"time"

	"github.com/onsi/ginkgo/v2"

	"k8s.io/kubernetes/test/e2e/framework"
	e2enode "k8s.io/kubernetes/test/e2e/framework/node"
)

func init() {
	framework.NewFrameworkExtensions = append(framework.NewFrameworkExtensions,
		func(f *framework.Framework) {
			ginkgo.AfterEach(func() {
				if f.ClientSet == nil {
					// Test didn't reach f.BeforeEach, most
					// likely because the test got
					// skipped. Nothing to check...
					return
				}
				e2enode.AllNodesReady(f.ClientSet, 7*time.Minute)
			})
		},
	)
}
