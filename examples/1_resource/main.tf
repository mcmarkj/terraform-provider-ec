provider "ec" {}

resource "ec_deployment" "this" {
  name                = "hello-kitty"
  region              = "gcp-europe-west1"
  version             = "7.8.1"
  template_id         = "gcp-io-optimized"
  data_node           = true
  master_node         = true
  ingest_node         = true
  ml_node             = false
  elastic_instance_id = "gcp.data.highio.1"
  elastic_zone_count  = 1
  elastic_node_memory = 2048
}

output "cluster_user" {
  description = "Cluster user"
  value       = ec_deployment.this.username
}

output "cluster_password" {
  description = "Cluster password"
  value       = ec_deployment.this.password
}

output "cluster_id" {
  description = "Cluster ID"
  value       = ec_deployment.this.cluster_id
}

output "cluster_endpoint" {
  description = "Elasticsearch ednpoint"
  value       = ec_deployment.this.endpoint
}

output "deployment_id" {
  description = "Deployment ID"
  value       = ec_deployment.this.deployment_id
}
