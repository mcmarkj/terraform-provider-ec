## Elastic Cloud provider

This repo contains source code for terraform provider: __terraform-provider-ec__.
https://www.elastic.co/guide/en/cloud/current/ec-restful-api.html

## Build

- Build with pre-installed golang:

```sh
$ go build -v
```
- Build with docker (make sure Docker is installed). It will pull golang image, build artifact and place it in current directory. You don't need to care about dependencies:

```sh
$ ./build_docker.sh

```

At the end, you'll get binary file `terraform-provider-ec`.

## Usage

- To start using compiled plugin `terraform-provider-ec` you have to place this file into your working directory with Terrafrom code or use [other path](https://www.terraform.io/docs/configuration/providers.html#third-party-plugins).
  - command to download binary file `terraform-provider-ec` for current OS in `$HOME/.terraform.d/plugins` directory:
```sh
$ ./download_binary.sh {{release_version}}

```
- Elasticsearch Service supports only API key-based [authentication](https://www.elastic.co/guide/en/cloud/current/ec-api-authentication.html).

## Environment variables

Set API key as an environment variable `EC_API_KEY`:
```sh
export EC_API_KEY = "Qwerty12345678"
```

Obviously, you can set up it in your terrafrom code, but it's not recommended to store secrets in control version systems.

## Provider

```sh

provider "ec" {
  url      = "https://api.elastic-cloud.com"
  api_key  = "Qwerty12345678"
  insecure = false
}

# Min usage:
provider "ec" {}
```
  - `url` - Elastic Cloud endpoint. (Default: https://api.elastic-cloud.com)
  - `api_key` - API Key. Set it with environment variable EC_API_KEY.
  - `insecure` - TLS setting for EC client. (Default: false - don't trust self-signed certificates)

## Data Source: ec_deployments

Use this data source to get names and IDs of existing deployments.
### Example usage

```sh

provider "ec" {}

data "ec_deployments" "this" {}
```
### Attributes Reference

  - `ids` - mapping `deployment name`: `id`

## Resource: ec_deployment

Manage ElasticCloud deployment

### Example usage

```sh

provider "ec" {}

resource "ec_deployment" "this" {
  name                = "my-first-api-deployment"
  region              = "gcp-europe-west1"
  version             = "7.8.0"
  template_id         = "gcp-io-optimized"

  data_node           = true
  master_node         = true
  ingest_node         = true
  ml_node             = false

  elastic_instance_id = "gcp.data.highio.1"
  elastic_zone_count  = 1
  elastic_node_memory = 8192
}

```
### Argument Reference

  - `name` - (Required) Deployment name.
  - `region` - (Required) Geographic area where the data center of the cloud provider that hosts your deployment is located. [List of available regions](https://www.elastic.co/guide/en/cloud/current/ec-regions-templates-instances.html).
  - `version` - (Required) Elastic stack version.
  - `template_id` - (Required) Unique identifier of the [deployment template](https://www.elastic.co/guide/en/cloud/current/ec-getting-started-templates.html) - pre-configure the components of the Elastic Stack.
  - `data_node` - (Bool) Controls the combinations of Elasticsearch node types. Defines whether this node can hold data.
  - `master_node` - (Bool) Controls the combinations of Elasticsearch node types. Defines whether this node can be elected master.
  - `ingest_node` - (Bool) Controls the combinations of Elasticsearch node types. Defines whether this node can run an ingest pipeline.
  - `ml_node` - (Bool) Controls the combinations of Elasticsearch node types. Defines whether this node can run ml jobs.
  - `elastic_instance_id` - Instance ID for elastic node. Depends on cloud provider.
  - `elastic_zone_count` - Number of availability zones in selected region.
  - `elastic_node_memory` - Amount of memory (MB) per elastic node.
  
### Attributes Reference

  - `username` - elasticsearch username.
  - `password` - elasticsearch password.
  - `endpoint` - elasticsearch cluster ednpoint.
  - `cluster_id` - elasticsearch cluster ID.
  - `deployment_id` - elasticcloud deployment ID.

## Debugging

By default, provider log messages are not written to standard out during provider execution. To enable verbose output of Terraform and provider log messages, set the `TF_LOG` environment variable to `DEBUG`.
