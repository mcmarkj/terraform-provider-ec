package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
)

const (
	elasticDomain = "cloud.es.io"
	elasticPort   = "9243"
	elasticSchema = "https"
)

/*

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

*/
func resourceElasticsearchDeployment() *schema.Resource {
	return &schema.Resource{
		Create: resourceElasticsearchDeploymentCreate,
		Read:   resourceElasticsearchDeploymentRead,
		Update: resourceElasticsearchDeploymentUpdate,
		Delete: resourceElasticsearchDeploymentDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"elastic_instance_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Instance ID for elastic node. Depends on cloud provider",
			},
			"elastic_zone_count": &schema.Schema{
				Type:        schema.TypeInt,
				Description: "Number of availability zones in selected region",
				Required:    true,
			},
			"elastic_node_memory": &schema.Schema{
				Type:        schema.TypeInt,
				Description: "Amount of memory (MB) per elasticsearch node",
				Required:    true,
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "A name for the deployment; otherwise this will be the generated deployment id",
				ForceNew:    false,
				Required:    true,
			},
			"region": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Geographic area where the data center of the cloud provider that hosts your deployment is located",
				ForceNew:    false,
				Required:    true,
			},
			"version": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Elastic stack version",
				ForceNew:    false,
				Required:    true,
			},
			"template_id": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Unique identifier of the deployment template - pre-configure the components of the Elastic Stack",
				ForceNew:    false,
				Required:    true,
			},
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Elasticsearch username",
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Elasticsearch password",
			},
			"endpoint": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Elasticsearch cluster endpoint",
			},
			"cluster_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Cluster ID",
			},
			"deployment_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Elastic Cloud deployment ID",
			},
			"data_node": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "Defines whether elasticsearch node can hold data",
				ForceNew:    false,
				Optional:    true,
				Default:     true,
			},
			"master_node": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "Defines whether elasticsearch node can be elected master",
				ForceNew:    false,
				Optional:    true,
				Default:     true,
			},
			"ingest_node": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "Defines whether elasticsearch node can run an ingest pipeline",
				ForceNew:    false,
				Optional:    true,
				Default:     true,
			},
			"ml_node": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "Defines whether elasticsearch node can run Machine Learning jobs",
				ForceNew:    false,
				Optional:    true,
				Default:     false,
			},
		},
	}
}

func resourceElasticsearchDeploymentCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ECClient)

	deploymentName := d.Get("name").(string)
	region := d.Get("region").(string)
	version := d.Get("version").(string)
	templateID := d.Get("template_id").(string)

	elasticInstanceID := d.Get("elastic_instance_id").(string)
	elasticZoneCount := d.Get("elastic_zone_count").(int)
	elasticNodeMemory := d.Get("elastic_node_memory").(int)

	dataNode := d.Get("data_node").(bool)
	masterNode := d.Get("master_node").(bool)
	ingestNode := d.Get("ingest_node").(bool)
	mlNode := d.Get("ml_node").(bool)

	log.Printf("[DEBUG] Creating deployment with name: %s\n", deploymentName)

	elasticsearchNodeType := ElasticsearchNodeType{
		Data:   dataNode,
		ML:     mlNode,
		Ingest: ingestNode,
		Master: masterNode,
	}
	topologySize := TopologySize{
		Resource: "memory",
		Value:    elasticNodeMemory,
	}

	clusterTopology := ElasticsearchClusterTopologyElement{
		NodeType:                elasticsearchNodeType,
		InstanceConfigurationID: elasticInstanceID,
		ZoneCount:               elasticZoneCount,
		Size:                    topologySize,
	}
	elasticsearchConfiguration := ElasticsearchConfiguration{
		Version: version,
	}
	deploymentTemplate := DeploymentTemplateReference{
		ID: templateID,
	}

	elasticsearchClusterPlan := ElasticsearchClusterPlan{
		ClusterTopology:    []ElasticsearchClusterTopologyElement{clusterTopology},
		Elasticsearch:      elasticsearchConfiguration,
		DeploymentTemplate: deploymentTemplate,
	}

	elasticsearchPayload := ElasticsearchPayload{
		Region: region,
		RefID:  "elasticsearch",
		Plan:   elasticsearchClusterPlan,
	}

	resources := DeploymentCreateResources{
		Elasticsearch: []ElasticsearchPayload{elasticsearchPayload},
	}

	createDeploymentRequest := DeploymentCreateRequest{
		Name:      deploymentName,
		Resources: resources,
	}

	createResponse, err := client.CreateDeployment(createDeploymentRequest)
	if err != nil {
		return err
	}

	deploymentID := createResponse.ID

	log.Printf("[DEBUG] Created deployment ID: %s\n", deploymentID)

	clusterEndpoint := getElasticSearchFQDN(region, createResponse.Resources[0].ID)

	err = client.WaitForElasticsearchDeploymentStatus(deploymentID, "started", false)
	if err != nil {
		return err
	}

	d.SetId(deploymentID)
	d.Set("deployment_id", deploymentID)
	d.Set("username", createResponse.Resources[0].Credentials.Username)
	d.Set("password", createResponse.Resources[0].Credentials.Password)
	d.Set("cluster_id", createResponse.Resources[0].ID)
	d.Set("endpoint", clusterEndpoint)

	return resourceElasticsearchDeploymentRead(d, meta)
}

func getElasticSearchFQDN(region string, clusterID string) string {
	cloudProvider := strings.Split(region, "-")[0]
	shortRegion := strings.TrimPrefix(region, cloudProvider+"-")
	fqdn := elasticSchema + "://" + clusterID + "." + shortRegion + "." + cloudProvider + "." + elasticDomain + ":" + elasticPort
	return fqdn
}

func resourceElasticsearchDeploymentRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ECClient)

	deploymentID := d.Id()
	log.Printf("[DEBUG] Reading deployment information for ID: %s\n", deploymentID)

	resp, err := client.GetDeployment(deploymentID)
	if err != nil {
		return err
	}

	if resp.StatusCode == 404 {
		log.Printf("[DEBUG] Deployment ID not found: %s\n", deploymentID)
		d.SetId("")
		return nil
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Deployment response body: %v\n", string(respBytes))

	var deployInfo DeploymentGetResponse
	err = json.Unmarshal(respBytes, &deployInfo)
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] This is deployinfo: %v\n", deployInfo)

	region := deployInfo.Resources.Elasticsearch[0].Region
	clusterID := deployInfo.Resources.Elasticsearch[0].Info.ClusterID

	d.Set("name", deployInfo.Name)
	d.Set("region", region)

	elasticPlan := deployInfo.Resources.Elasticsearch[0].Info.PlanInfo.Current.Plan
	d.Set("version", elasticPlan.Elasticsearch.Version)
	d.Set("template_id", elasticPlan.DeploymentTemplate.ID)
	d.Set("elastic_instance_id", elasticPlan.ClusterTopology[0].InstanceConfigurationID)
	d.Set("elastic_zone_count", elasticPlan.ClusterTopology[0].ZoneCount)
	d.Set("elastic_node_memory", elasticPlan.ClusterTopology[0].Size.Value)
	d.Set("ingest_node", elasticPlan.ClusterTopology[0].NodeType.Ingest)
	d.Set("master_node", elasticPlan.ClusterTopology[0].NodeType.Master)
	d.Set("data_node", elasticPlan.ClusterTopology[0].NodeType.Data)
	d.Set("ml_node", elasticPlan.ClusterTopology[0].NodeType.ML)
	d.Set("cluster_id", clusterID)
	d.Set("deployment_id", deployInfo.Resources.Elasticsearch[0].Info.DeploymentID)
	d.Set("endpoint", getElasticSearchFQDN(region, clusterID))

	return nil
}

func resourceElasticsearchDeploymentDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ECClient)
	deploymentID := d.Id()
	log.Printf("[DEBUG] Deleting deployment ID: %s\n", deploymentID)
	_, err := client.DeleteDeployment(deploymentID)

	if err != nil {
		return err
	}

	return nil
}

func resourceElasticsearchDeploymentUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ECClient)
	d.Partial(false)
	deploymentID := d.Id()
	log.Printf("[DEBUG] Updating elasticsearch deployment ID: %s\n", deploymentID)

	resp, err := client.GetDeployment(deploymentID)
	if err != nil {
		return err
	}

	if resp.StatusCode == 404 {
		return fmt.Errorf("%q: deployment ID was not found for update", deploymentID)
	}

	deploymentName := d.Get("name").(string)
	region := d.Get("region").(string)
	version := d.Get("version").(string)
	templateID := d.Get("template_id").(string)

	elasticInstanceID := d.Get("elastic_instance_id").(string)
	elasticZoneCount := d.Get("elastic_zone_count").(int)
	elasticNodeMemory := d.Get("elastic_node_memory").(int)

	dataNode := d.Get("data_node").(bool)
	masterNode := d.Get("master_node").(bool)
	ingestNode := d.Get("ingest_node").(bool)
	mlNode := d.Get("ml_node").(bool)

	elasticsearchNodeType := ElasticsearchNodeType{
		Data:   dataNode,
		ML:     mlNode,
		Ingest: ingestNode,
		Master: masterNode,
	}
	topologySize := TopologySize{
		Resource: "memory",
		Value:    elasticNodeMemory,
	}

	clusterTopology := ElasticsearchClusterTopologyElement{
		NodeType:                elasticsearchNodeType,
		InstanceConfigurationID: elasticInstanceID,
		ZoneCount:               elasticZoneCount,
		Size:                    topologySize,
	}
	elasticsearchConfiguration := ElasticsearchConfiguration{
		Version: version,
	}
	deploymentTemplate := DeploymentTemplateReference{
		ID: templateID,
	}

	elasticsearchClusterPlan := ElasticsearchClusterPlan{
		ClusterTopology:    []ElasticsearchClusterTopologyElement{clusterTopology},
		Elasticsearch:      elasticsearchConfiguration,
		DeploymentTemplate: deploymentTemplate,
	}

	elasticsearchPayload := ElasticsearchPayload{
		Region: region,
		RefID:  "elasticsearch",
		Plan:   elasticsearchClusterPlan,
	}

	resources := DeploymentUpdateResources{
		Elasticsearch: []ElasticsearchPayload{elasticsearchPayload},
	}

	updateDeploymentRequest := DeploymentUpdateRequest{
		Name:         deploymentName,
		Resources:    resources,
		PruneOrphans: false,
	}

	_, err = client.UpdateDeployment(deploymentID, updateDeploymentRequest)

	if err != nil {
		return err
	}

	err = client.WaitForElasticsearchDeploymentStatus(deploymentID, "started", false)
	if err != nil {
		return err
	}

	return resourceElasticsearchDeploymentRead(d, meta)
}

