/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	servicebus "github.com/Azure/azure-service-operator/hack/generated/_apis/microsoft.servicebus/v1alpha1api20180101preview"
	"github.com/Azure/azure-service-operator/hack/generated/pkg/testcommon"
)

func Test_ServiceBus_Standard_CRUD(t *testing.T) {
	t.Parallel()

	testContext := testContext.ForTest(t)

	rg := testContext.CreateNewTestResourceGroupAndWait()

	zoneRedundant := false
	namespace := &servicebus.Namespace{
		ObjectMeta: testContext.MakeObjectMetaWithName(testContext.Namer.GenerateName("sbstandard")),
		Spec: servicebus.Namespaces_Spec{
			Location: testContext.AzureRegion,
			Owner:    testcommon.AsOwner(rg.ObjectMeta),
			Sku: &servicebus.SBSku{
				Name: servicebus.SBSkuNameStandard,
			},
			Properties: servicebus.SBNamespaceProperties{
				ZoneRedundant: &zoneRedundant,
			},
		},
	}

	testContext.CreateAndWait(namespace)

	testContext.Expect(namespace.Status.Id).ToNot(BeNil())
	armId := *namespace.Status.Id

	testContext.RunParallelSubtests(
		testcommon.Subtest{
			Name: "Queue CRUD",
			Test: func(t testcommon.KubePerTestContext) { ServiceBus_Queue_CRUD(testContext, namespace.ObjectMeta) },
		},
		testcommon.Subtest{
			Name: "Topic CRUD",
			Test: func(t testcommon.KubePerTestContext) { ServiceBus_Topic_CRUD(testContext, namespace.ObjectMeta) },
		},
	)

	testContext.DeleteAndWait(namespace)

	// Ensure that the resource was really deleted in Azure
	exists, retryAfter, err := testContext.AzureClient.HeadResource(testContext.Ctx, armId, "2018-01-01-preview")
	testContext.Expect(err).ToNot(HaveOccurred())
	testContext.Expect(retryAfter).To(BeZero())
	testContext.Expect(exists).To(BeFalse())
}

// Topics can only be created in Standard or Premium SKUs
func ServiceBus_Topic_CRUD(testContext testcommon.KubePerTestContext, sbNamespace metav1.ObjectMeta) {

	topic := &servicebus.NamespacesTopic{
		ObjectMeta: testContext.MakeObjectMeta("topic"),
		Spec: servicebus.NamespacesTopics_Spec{
			Location: &testContext.AzureRegion,
			Owner:    testcommon.AsOwner(sbNamespace),
		},
	}

	testContext.CreateAndWait(topic)
	defer testContext.DeleteAndWait(topic)

	testContext.Expect(topic.Status.Id).ToNot(BeNil())

	// a basic assertion on a property
	testContext.Expect(topic.Status.Properties.SizeInBytes).ToNot(BeNil())
	testContext.Expect(*topic.Status.Properties.SizeInBytes).To(Equal(0))
}
