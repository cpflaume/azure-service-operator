/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	network "github.com/Azure/azure-service-operator/hack/generated/_apis/microsoft.network/v1alpha1api20171001"
	"github.com/Azure/azure-service-operator/hack/generated/pkg/testcommon"
)

func Test_VirtualNetwork_CRUD(t *testing.T) {
	t.Parallel()

	testContext := testContext.ForTest(t)

	rg := testContext.CreateNewTestResourceGroupAndWait()

	vn := &network.VirtualNetwork{
		ObjectMeta: testContext.MakeObjectMetaWithName(testContext.Namer.GenerateName("vn")),
		Spec: network.VirtualNetworks_Spec{
			Owner:    testcommon.AsOwner(rg.ObjectMeta),
			Location: &testcommon.DefaultTestRegion,
			Properties: network.VirtualNetworkPropertiesFormat{
				AddressSpace: &network.AddressSpace{
					AddressPrefixes: []string{"10.0.0.0/8"},
				},
			},
		},
	}

	testContext.CreateAndWait(vn)

	testContext.Expect(vn.Status.Id).ToNot(BeNil())
	armId := *vn.Status.Id

	testContext.DeleteAndWait(vn)

	// Ensure that the resource was really deleted in Azure
	exists, retryAfter, err := testContext.AzureClient.HeadResource(testContext.Ctx, armId, "2017-10-01")
	testContext.Expect(err).ToNot(HaveOccurred())
	testContext.Expect(retryAfter).To(BeZero())
	testContext.Expect(exists).To(BeFalse())
}
