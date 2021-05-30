/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	documentdb "github.com/Azure/azure-service-operator/hack/generated/_apis/microsoft.documentdb/v1alpha1api20150408"
	"github.com/Azure/azure-service-operator/hack/generated/pkg/testcommon"
)

func Test_CosmosDB_CRUD(t *testing.T) {
	t.Parallel()

	testContext := testContext.ForTest(t)

	rg := testContext.CreateNewTestResourceGroupAndWait()

	// Custom namer because storage accounts have strict names
	namer := testContext.Namer.WithSeparator("")

	// Create a Cosmos DB account
	kind := documentdb.DatabaseAccountsSpecKindGlobalDocumentDB
	acct := &documentdb.DatabaseAccount{
		ObjectMeta: testContext.MakeObjectMetaWithName(namer.GenerateName("db")),
		Spec: documentdb.DatabaseAccounts_Spec{
			Location: &testContext.AzureRegion,
			Owner:    testcommon.AsOwner(rg.ObjectMeta),
			Kind:     &kind,
			Properties: documentdb.DatabaseAccountCreateUpdateProperties{
				DatabaseAccountOfferType: documentdb.DatabaseAccountCreateUpdatePropertiesDatabaseAccountOfferTypeStandard,
				Locations: []documentdb.Location{
					{
						LocationName: &testContext.AzureRegion,
					},
				},
			},
		},
	}

	testContext.CreateAndWait(acct)

	expectedKind := documentdb.DatabaseAccountStatusKindGlobalDocumentDB
	testContext.Expect(*acct.Status.Kind).To(Equal(expectedKind))

	testContext.Expect(acct.Status.Id).ToNot(BeNil())
	armId := *acct.Status.Id

	// Run sub-tests
	/*
		t.Run("Blob Services CRUD", func(t *testing.T) {
			StorageAccount_BlobServices_CRUD(t, testContext, acct.ObjectMeta)
		})
	*/

	testContext.DeleteAndWait(acct)

	// Ensure that the resource group was really deleted in Azure
	exists, retryAfter, err := testContext.AzureClient.HeadResource(testContext.Ctx, armId, "2015-04-08")
	testContext.Expect(err).ToNot(HaveOccurred())
	testContext.Expect(retryAfter).To(BeZero())
	testContext.Expect(exists).To(BeFalse())
}
