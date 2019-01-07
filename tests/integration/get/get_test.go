// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

package ints

import (
	"os"
	"testing"

	"github.com/pulumi/pulumi-kubernetes/pkg/openapi"
	"github.com/pulumi/pulumi-kubernetes/tests"
	"github.com/pulumi/pulumi/pkg/resource"
	"github.com/pulumi/pulumi/pkg/resource/deploy/providers"
	"github.com/pulumi/pulumi/pkg/testing/integration"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	kubectx := os.Getenv("KUBERNETES_CONTEXT")

	if kubectx == "" {
		t.Skipf("Skipping test due to missing KUBERNETES_CONTEXT variable")
	}

	integration.ProgramTest(t, &integration.ProgramTestOptions{
		Dir:                  "step1",
		Dependencies:         []string{"@pulumi/kubernetes"},
		Quick:                true,
		ExpectRefreshChanges: true,
		ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
			assert.NotNil(t, stackInfo.Deployment)
			assert.Equal(t, 6, len(stackInfo.Deployment.Resources))

			tests.SortResourcesByURN(stackInfo)

			stackRes := stackInfo.Deployment.Resources[5]
			assert.Equal(t, resource.RootStackType, stackRes.URN.Type())

			provRes := stackInfo.Deployment.Resources[4]
			assert.True(t, providers.IsProviderType(provRes.URN.Type()))

			//
			// Assert we can use .get to retrieve the Kubernetes dashboard Service.
			//

			pod := stackInfo.Deployment.Resources[1]
			assert.Equal(t, "kube-dashboard", string(pod.URN.Name()))
			step1Name, _ := openapi.Pluck(pod.Outputs, "metadata", "name")
			assert.Equal(t, "kubernetes", step1Name.(string))

			//
			// Assert we can use .get to retrieve CRDs.
			//

			crd := stackInfo.Deployment.Resources[0]
			assert.Equal(t, "crontab", string(crd.URN.Name()))

			ct1 := stackInfo.Deployment.Resources[2]
			assert.Equal(t, "my-new-cron-object", string(ct1.URN.Name()))

			ct2 := stackInfo.Deployment.Resources[3]
			assert.Equal(t, "my-new-cron-object-get", string(ct2.URN.Name()))

		},
	})
}