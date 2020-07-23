package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/url"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

/*

provider "ec" {
	url 	 = "https://api.elastic-cloud.com"
	api_key  = "Qwerty12345678"
	insecure = false
}

Minimal usage:
export EC_API_KEY = "Qwerty12345678"

provider "ec" {}

*/

// Provider for ECE cluster management using Terraform.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"url": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "https://api.elastic-cloud.com",
				Description: "Elastic Cloud API endpoint",
			},
			"api_key": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("EC_API_KEY", nil),
				Description: "Elastic Cloud API key",
			},
			"insecure": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Disable certificate verification of API calls",
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"ec_deployments": resourceDeployments(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"ec_deployment": resourceElasticsearchDeployment(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	rawURL := d.Get("url").(string)
	log.Printf("[DEBUG] Connecting to EC: %s\n", rawURL)

	_, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}

	key := d.Get("api_key").(string)
	httpClient := getHTTPClient(d)
	ecClient := &ECClient{
		HTTPClient: httpClient,
		BaseURL:    rawURL,
		Key:        key,
	}

	return ecClient, nil
}

func getHTTPClient(d *schema.ResourceData) *http.Client {
	insecure := d.Get("insecure").(bool)

	tlsConfig := &tls.Config{}

	if insecure {
		tlsConfig.InsecureSkipVerify = true
	}

	transport := &http.Transport{TLSClientConfig: tlsConfig}

	client := &http.Client{Transport: transport}

	log.Printf("[DEBUG] HTTP client timeout: %v\n", client.Timeout)

	return client
}

