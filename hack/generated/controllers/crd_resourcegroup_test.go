/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"
	"sigs.k8s.io/controller-runtime/pkg/client"

	resources "github.com/Azure/azure-service-operator/hack/generated/_apis/microsoft.resources/v1alpha1api20200601"
	"github.com/Azure/azure-service-operator/hack/generated/pkg/armclient"
)

func Test_ResourceGroup_CRUD(t *testing.T) {
	t.Parallel()

	g := NewGomegaWithT(t)
	testContext := testContext.ForTest(t)

	// Create a resource group
	rg := testContext.NewTestResourceGroup()
	testContext.CreateAndWait(rg)

	// check properties
	g.Expect(rg.Status.Location).To(Equal(testContext.AzureRegion))
	g.Expect(rg.Status.Properties.ProvisioningState).To(Equal(string(armclient.SucceededProvisioningState)))
	g.Expect(rg.Status.ID).ToNot(BeNil())
	armId := rg.Status.ID

	// Update the tags
	rg.Spec.Tags["tag1"] = "value1"
	testContext.Update(rg)

	objectKey, err := client.ObjectKeyFromObject(rg)
	g.Expect(err).ToNot(HaveOccurred())

	// ensure they get updated
	testContext.Eventually(func() map[string]string {
		newRG := &resources.ResourceGroup{}
		testContext.Get(objectKey, newRG)
		return newRG.Status.Tags
	}).Should(HaveKeyWithValue("tag1", "value1"))

	testContext.DeleteAndWait(rg)

	// Ensure that the resource group was really deleted in Azure
	// TODO: Do we want to just use an SDK here? This process is quite icky as is...
	exists, _, err := testContext.AzureClient.HeadResource(
		testContext.Ctx,
		armId,
		"2020-06-01")
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(exists).To(BeFalse())
}
