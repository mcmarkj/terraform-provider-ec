provider "ec" {}

data "ec_deployments" "this" {}

output "deployment_map" {
  value = data.ec_deployments.this.ids
}
