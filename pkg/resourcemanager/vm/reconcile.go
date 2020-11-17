// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package vm

import (
	"context"
	"fmt"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

func (c *AzureVirtualMachineClient) Ensure(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {

	instance, err := c.convert(obj)
	if err != nil {
		return true, err
	}

	client := getVirtualMachineClient(c.Creds)

	location := instance.Spec.Location
	resourceGroup := instance.Spec.ResourceGroup
	resourceName := instance.Name
	vmSize := instance.Spec.VMSize
	osType := instance.Spec.OSType
	adminUserName := instance.Spec.AdminUserName
	sshPublicKeyData := instance.Spec.SSHPublicKeyData
	nicName := instance.Spec.NetworkInterfaceName
	imageURN := instance.Spec.PlatformImageURN

	const SucceededProvisioningState = "Succeeded"

	// Check to see if secret exists and if yes retrieve the admin login and password
	secret, err := c.GetOrPrepareSecret(ctx, instance)
	if err != nil {
		return false, err
	}
	// Update secret
	err = c.AddVirtualMachineCredsToSecrets(ctx, instance.Name, secret, instance)
	if err != nil {
		return false, err
	}

	adminPassword := string(secret["password"])

	instance.Status.Provisioning = true
	// Check if this item already exists. This is required
	// to overcome the issue with the lack of idempotence of the Create call
	item, err := c.GetVirtualMachine(ctx, resourceGroup, resourceName)
	if err == nil {

		if *item.ProvisioningState == SucceededProvisioningState {

			instance.Status.Provisioned = true
			instance.Status.Provisioning = false
			instance.Status.Message = resourcemanager.SuccessMsg
			instance.Status.ResourceId = *item.ID
			instance.Status.State = *item.ProvisioningState
			return true, nil
		}

		instance.Status.Provisioned = false
		instance.Status.State = *item.ProvisioningState
		instance.Status.Message = "Requested resource has been requested but is not ready yet"
		return false, nil

	}
	future, err := c.CreateVirtualMachine(
		ctx,
		location,
		resourceGroup,
		resourceName,
		vmSize,
		string(osType),
		adminUserName,
		adminPassword,
		sshPublicKeyData,
		nicName,
		imageURN,
	)
	if err != nil {
		// let the user know what happened
		instance.Status.Message = err.Error()
		instance.Status.Provisioning = false
		// errors we expect might happen that we are ok with waiting for
		catch := []string{
			errhelp.ResourceGroupNotFoundErrorCode,
			errhelp.ParentNotFoundErrorCode,
			errhelp.NotFoundErrorCode,
			errhelp.AsyncOpIncompleteError,
			errhelp.ResourceNotFound,
			errhelp.InvalidResourceReference,
		}

		azerr := errhelp.NewAzureError(err)
		if helpers.ContainsString(catch, azerr.Type) {
			// most of these error technically mean the resource is actually not provisioning
			switch azerr.Type {
			case errhelp.AsyncOpIncompleteError:
				instance.Status.Provisioning = true

			}
			// reconciliation is not done but error is acceptable
			return false, nil
		}
		//change to spec is required
		if future.Response().StatusCode == 400 {
			instance.Status.FailedProvisioning = true
			return false, nil
		}

		// reconciliation not done and we don't know what happened
		return false, err
	}

	_, err = future.Result(client)
	if err != nil {
		// let the user know what happened
		instance.Status.Message = err.Error()
		instance.Status.Provisioning = false
		// errors we expect might happen that we are ok with waiting for
		catch := []string{
			errhelp.ResourceGroupNotFoundErrorCode,
			errhelp.ParentNotFoundErrorCode,
			errhelp.NotFoundErrorCode,
			errhelp.AsyncOpIncompleteError,
			errhelp.SubscriptionDoesNotHaveServer,
		}

		azerr := errhelp.NewAzureError(err)
		if helpers.ContainsString(catch, azerr.Type) {
			// most of these error technically mean the resource is actually not provisioning
			switch azerr.Type {
			case errhelp.AsyncOpIncompleteError:
				instance.Status.Provisioning = true
			}
			// reconciliation is not done but error is acceptable
			return false, nil
		}

		// reconciliation not done and we don't know what happened
		return false, err
	}
	return false, nil
}

func (c *AzureVirtualMachineClient) Delete(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {

	instance, err := c.convert(obj)
	if err != nil {
		return true, err
	}

	resourceGroup := instance.Spec.ResourceGroup
	resourceName := instance.Name

	status, err := c.DeleteVirtualMachine(
		ctx,
		resourceName,
		resourceGroup,
	)
	if err != nil {
		catch := []string{
			errhelp.AsyncOpIncompleteError,
		}
		gone := []string{
			errhelp.ResourceGroupNotFoundErrorCode,
			errhelp.ParentNotFoundErrorCode,
			errhelp.NotFoundErrorCode,
			errhelp.ResourceNotFound,
		}
		azerr := errhelp.NewAzureError(err)
		if helpers.ContainsString(catch, azerr.Type) {
			return true, nil
		} else if helpers.ContainsString(gone, azerr.Type) {
			return false, nil
		}
		return true, err
	}

	if err == nil {
		if status != "InProgress" {
			return false, nil
		}
	}

	return true, nil
}
func (g *AzureVirtualMachineClient) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {

	instance, err := g.convert(obj)
	if err != nil {
		return nil, err
	}

	return []resourcemanager.KubeParent{
		{
			Key: types.NamespacedName{
				Namespace: instance.Namespace,
				Name:      instance.Spec.ResourceGroup,
			},
			Target: &azurev1alpha1.ResourceGroup{},
		},
	}, nil
}

func (g *AzureVirtualMachineClient) GetStatus(obj runtime.Object) (*azurev1alpha1.ASOStatus, error) {

	instance, err := g.convert(obj)
	if err != nil {
		return nil, err
	}
	return &instance.Status, nil
}

func (g *AzureVirtualMachineClient) convert(obj runtime.Object) (*azurev1alpha1.AzureVirtualMachine, error) {
	local, ok := obj.(*azurev1alpha1.AzureVirtualMachine)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}
