/*
Copyright 2020 The Kubernetes Authors.

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

package diskclient

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2020-06-30/compute"

	"sigs.k8s.io/cloud-provider-azure/pkg/retry"
)

const (
	// APIVersion is the API version for compute.
	APIVersion = "2019-11-01"
	// AzureStackCloudAPIVersion is the API version for Azure Stack
	AzureStackCloudAPIVersion = "2019-07-01"
	// AzureStackCloudName is the cloud name of Azure Stack
	AzureStackCloudName = "AZURESTACKCLOUD"
)

// Interface is the client interface for Disks.
// Don't forget to run the following command to generate the mock client:
// mockgen -source=$GOPATH/src/sigs.k8s.io/cloud-provider-azure/pkg/azureclients/diskclient/interface.go -package=mockdiskclient Interface > $GOPATH/src/sigs.k8s.io/cloud-provider-azure/pkg/azureclients/diskclient/mockdiskclient/interface.go
type Interface interface {
	// Get gets a Disk.
	Get(ctx context.Context, resourceGroupName string, diskName string) (result compute.Disk, rerr *retry.Error)

	// CreateOrUpdate creates or updates a Disk.
	CreateOrUpdate(ctx context.Context, resourceGroupName string, diskName string, diskParameter compute.Disk) *retry.Error

	// Update updates a Disk.
	Update(ctx context.Context, resourceGroupName string, diskName string, diskParameter compute.DiskUpdate) *retry.Error

	// Delete deletes a Disk by name.
	Delete(ctx context.Context, resourceGroupName string, diskName string) *retry.Error

	// ListByResourceGroup lists all the disks under a resource group.
	ListByResourceGroup(ctx context.Context, resourceGroupName string) ([]compute.Disk, *retry.Error)
}
