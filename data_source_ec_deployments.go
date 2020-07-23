package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
)

/*

data "ec_deployments" "this" {}

output "deployment_map" {
  value = data.ec_deployments.this.ids
}


*/

func resourceDeployments() *schema.Resource {
	return &schema.Resource{
		Read: resourceDeploymentsRead,
		Schema: map[string]*schema.Schema{
			"ids": &schema.Schema{
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func resourceDeploymentsRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ECClient)
	log.Printf("[DEBUG] Listing deployments data\n")

	resp, err := client.ListDeployments()
	if err != nil {
		return err
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] EC response body: %v\n", string(respBytes))
	var deployinfo DeploymentsListResponse

	err = json.Unmarshal(respBytes, &deployinfo)

	if err != nil {
		return err
	}

	deployMap := make(map[string]string)
	for _, p := range deployinfo.Deployments {
		deployMap[p.Name] = p.ID
	}
	err = d.Set("ids", deployMap)
	d.SetId(fmt.Sprintf("Deployments_%s", time.Now().UTC().String()))

	return nil
}

