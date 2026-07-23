package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// For this lab, use your assigned Azure subscription.
var subscriptionID = "d8fb93ba-c4c0-4034-bb7b-ad54c281d7bc"

func TestAzureLinuxVMCreation(t *testing.T) {
	terraformOptions := &terraform.Options{
		TerraformDir: "../",

		Vars: map[string]interface{}{
			"labelPrefix": "yang0371a06",
		},
	}

	// Automatically remove Azure resources after testing.
	defer terraform.Destroy(t, terraformOptions)

	// Deploy the Terraform infrastructure.
	terraform.InitAndApply(t, terraformOptions)

	// Retrieve Terraform output values.
	vmName := terraform.Output(t, terraformOptions, "vm_name")
	resourceGroupName := terraform.Output(
		t,
		terraformOptions,
		"resource_group_name",
	)
	nicName := terraform.Output(t, terraformOptions, "nic_name")

	t.Run("Confirm VM exists", func(t *testing.T) {
		vmExists := azure.VirtualMachineExists(
			t,
			vmName,
			resourceGroupName,
			subscriptionID,
		)

		assert.True(
			t,
			vmExists,
			"Expected VM %s to exist",
			vmName,
		)
	})

	t.Run("Confirm NIC exists and is connected to VM", func(t *testing.T) {
		// Confirm the NIC resource exists.
		nic, err := azure.GetNetworkInterfaceE(
			nicName,
			resourceGroupName,
			subscriptionID,
		)

		require.NoError(
			t,
			err,
			"Expected NIC %s to exist",
			nicName,
		)

		require.NotNil(
			t,
			nic,
			"Expected NIC information to be returned",
		)

		// Confirm that the NIC is attached to the VM.
		attachedNICs := azure.GetVirtualMachineNics(
			t,
			vmName,
			resourceGroupName,
			subscriptionID,
		)

		assert.Contains(
			t,
			attachedNICs,
			nicName,
			"Expected NIC %s to be attached to VM %s",
			nicName,
			vmName,
		)
	})

	t.Run("Confirm VM uses Ubuntu 22.04", func(t *testing.T) {
		image := azure.GetVirtualMachineImage(
			t,
			vmName,
			resourceGroupName,
			subscriptionID,
		)

		assert.Equal(
			t,
			"Canonical",
			image.Publisher,
			"Expected Canonical as the image publisher",
		)

		assert.Equal(
			t,
			"0001-com-ubuntu-server-jammy",
			image.Offer,
			"Expected Ubuntu Jammy image offer",
		)

		assert.Equal(
			t,
			"22_04-lts-gen2",
			image.SKU,
			"Expected Ubuntu 22.04 LTS Gen2",
		)
	})
}
