# Terraform Provider for Ivanti Service Manager

This repository holds the source code for a third party Ivanti Service Manager provider.  At original release, all objects will be treated as custom objects.  Common objects built into the product are welcome and will be added in as time allows.

## Project on hold ##

Please note, this project has been largeley abandoned due to competing priorities and lack of a testing instance.  Next date for update of this project is currently unknown.

## Requirements

-	[Terraform](https://www.terraform.io/downloads.html) >= 0.13.x, Terraform Enterprise, or Terraform Cloud set to >= 0.13.x
-   [Ivanti Service Manager](https://www.ivanti.com/products/enterprise-service-management) >= 2019.3.0.2019122901

## Using The Provider

1. Add the provider to your list of existing providers
1. Define the input parameters
1. Create resource and data declarations in your code