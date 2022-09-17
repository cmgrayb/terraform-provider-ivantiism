# Terraform Provider for Ivanti Service Manager

![Code Quality](https://github.com/cmgrayb/terraform-provider-ivantiism/actions/workflows/codeql-analysis.yml/badge.svg)
![Build Tests](https://github.com/cmgrayb/terraform-provider-ivantiism/actions/workflows/test.yml/badge.svg)
![Release](https://github.com/cmgrayb/terraform-provider-ivantiism/actions/workflows/release.yml/badge.svg)

This repository holds the source code for a third party Ivanti Service Manager provider.  At original release, all objects will be treated as custom objects.  Common objects built into the product are welcome and will be added in as time allows.

## Project on hold

Please note, this project has been largeley abandoned due to competing priorities and lack of a testing instance.  Next date for update of this project is currently unknown.

If you have a functional test environment and can contribute, please have a look at the [project](https://github.com/users/cmgrayb/projects/1) for current status of the provider and next steps.

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) >= 0.13.x, Terraform Enterprise, or Terraform Cloud set to >= 0.13.x
- [Ivanti Service Manager](https://www.ivanti.com/products/enterprise-service-management) >= 2019.3.0.2019122901

## Using The Provider

1. Add the provider to your list of existing providers
1. Define the input parameters
1. Create resource and data declarations in your code
