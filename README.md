#### CST8918 – Lab L-A06: Terraform Infrastructure Testing with Terratest


---

#### Overview

This lab demonstrates how to provision Azure infrastructure using **Terraform** and validate the deployment using **Terratest** written in Go.

The infrastructure deploys an Ubuntu Linux Virtual Machine and its supporting networking resources. Automated tests verify that the deployment completed successfully and that the VM configuration matches the expected requirements.

---

#### Technologies Used

- Terraform
- Azure Resource Manager (AzureRM) Provider
- Go
- Terratest
- Azure CLI
- Git
- GitHub

---

#### Azure Resources Created

Terraform provisions the following Azure resources:

- Resource Group
- Virtual Network (VNet)
- Subnet
- Network Security Group (NSG)
- Public IP Address
- Network Interface (NIC)
- Ubuntu 22.04 Linux Virtual Machine

---

#### Terraform Outputs

The following outputs are provided by Terraform:

| Output | Description |
|---------|-------------|
| resource_group_name | Azure Resource Group |
| vm_name | Virtual Machine Name |
| nic_name | Network Interface Name |
| public_ip | Public IP Address |

---

#### Terratest Validation

The Terratest suite automatically performs the following validation steps:

#### Test 1 – Deploy Infrastructure

- Terraform Init
- Terraform Apply

#### Test 2 – Verify Virtual Machine Exists

Confirms that the Azure Linux Virtual Machine was successfully created.

#### Test 3 – Verify Network Interface

Confirms:

- Network Interface exists
- Network Interface is attached to the Virtual Machine

#### Test 4 – Verify Ubuntu Image

Confirms the VM is running:

- Publisher: Canonical
- Offer: Ubuntu Server Jammy
- SKU: Ubuntu 22.04 LTS Gen2

#### Test 5 – Cleanup

Automatically destroys all Azure resources after testing.

---

# Project Structure

```text
.
├── main.tf
├── variables.tf
├── outputs.tf
├── providers.tf
├── README.md
├── test
│   └── azure_webserver_test.go
├── Screenshot
└── Testresult
```

---

#### Running the Project

#### Prerequisites

- Terraform
- Go
- Azure CLI
- Azure Subscription
- Terratest dependencies

Login to Azure:

```bash
az login
```

Initialize Terraform:

```bash
terraform init
```

---

#### Run Terratest

From the **test** directory:

```bash
go test -v -timeout 30m azure_webserver_test.go
```

---

#### Sample Successful Output

```text
=== RUN   TestAzureLinuxVMCreation
=== RUN   TestAzureLinuxVMCreation/Confirm_VM_exists
=== RUN   TestAzureLinuxVMCreation/Confirm_NIC_exists_and_is_connected_to_VM
=== RUN   TestAzureLinuxVMCreation/Confirm_VM_uses_Ubuntu_22.04

--- PASS: TestAzureLinuxVMCreation
PASS
```

---

#### Screenshots

The repository includes screenshots demonstrating:

- Successful Terraform deployment
- Azure Linux VM creation
- Terratest execution
- Successful test results

Screenshots are located in:

```text
Screenshot/
```

---

#### Test Results

The Terratest execution logs are included in:

```text
Testresult/
```

---

#### Cleanup

Terratest automatically executes:

```go
defer terraform.Destroy(t, terraformOptions)
```

to remove all Azure resources after the tests complete, preventing unnecessary Azure costs.

---

#### Learning Outcomes

This lab demonstrates:

- Infrastructure as Code (IaC)
- Terraform resource provisioning
- Azure infrastructure deployment
- Infrastructure validation using Terratest
- Automated testing with Go
- Continuous validation of cloud infrastructure

---

