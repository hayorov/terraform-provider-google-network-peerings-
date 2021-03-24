# Terraform Provider google-network-peerings

This repo is a companion repo to the [Call APIs with Terraform Providers](https://learn.hashicorp.com/collections/terraform/providers) Learn collection. 

In the collection, you will use the google-network-peerings provider as a bridge between Terraform and the google-network-peerings API. Then, extend Terraform by recreating the google-network-peerings provider. By the end of this collection, you will be able to take these intuitions to create your own custom Terraform provider. 

## Build provider

Run the following command to build the provider

```shell
$ go build -o terraform-provider-google-network-peerings
```

## Test sample configuration

First, build and install the provider.

```shell
$ make install
```

Then, navigate to the `examples` directory. 

```shell
$ cd examples
```

Run the following command to initialize the workspace and apply the sample configuration.

```shell
$ terraform init && terraform apply
```
