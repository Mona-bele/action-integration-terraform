# action-integration-terraform
Sets up Terraform API in your GitHub Actions workflow.

[![Go CI](https://github.com/Mona-bele/action-integration-terraform/actions/workflows/go_ci.yml/badge.svg)](https://github.com/Mona-bele/action-integration-terraform/actions/workflows/go_ci.yml) [![test integration](https://github.com/Mona-bele/action-integration-terraform/actions/workflows/ci-integration.yml/badge.svg)](https://github.com/Mona-bele/action-integration-terraform/actions/workflows/ci-integration.yml)

## Usage

The `Mona-bele/action-integration-terraform` action is used to set up Terraform API in your GitHub Actions workflow.

### Example workflow

```yaml
name: Terraform API
on:
  push:
    branches:
      - main
        
jobs:
  terraform-apply:
    name: 'Terraform'
    runs-on: ubuntu-latest
    environment: production
    steps:
      - name: Checkout
        uses: actions/checkout@v2
      
      - name: Terraform API Integration Create Variable
        uses: Mona-bele/action-integration-terraform@v1
        with:
          tf_api_token: ${{ secrets.TF_API_TOKEN }}
          tf_workspace: ${{ env.workspace }}
          tf_organization: ${{ env.organization }}
          tf_run_type: 'variable_create'
          tf_variable_key: 'test'
          tf_variable_value: 'test'
          tf_variable_category: 'terraform'
          tf_variable_sensitive: 'false'
          tf_variable_hcl: 'false'
          tf_variable_description: 'test'
          
      - name: Terraform API Integration Delete Variable
        uses: Mona-bele/action-integration-terraform@v1
        with:
          tf_api_token: ${{ secrets.TF_API_TOKEN }}
          tf_workspace: ${{ env.workspace }}
          tf_organization: ${{ env.organization }}
          tf_run_type: 'variable_delete'
          tf_variable_key: 'test'
      
      

```

# License
action-integration-terraform is licensed under the [MIT License](./LICENSE).