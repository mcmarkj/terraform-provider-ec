package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
)

const (
	jsonContentType    = "application/json"
	deploymentResource = "/api/v1/deployments"
)

// ECClient struct contains HTTPClient, BaseURL, Key
type ECClient struct {
	HTTPClient *http.Client
	BaseURL    string
	Key        string
}

// ListDeployments function get
func (c *ECClient) ListDeployments() (resp *http.Response, err error) {
	log.Printf("[DEBUG] ListDeployments\n")

	resourceURL := c.BaseURL + deploymentResource
	authString := "ApiKey " + c.Key
	log.Printf("[DEBUG] ListDeployments Resource URL: %s\n", resourceURL)
	req, err := http.NewRequest("GET", resourceURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", jsonContentType)
	req.Header.Set("Authorization", authString)

	resp, err = c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	log.Printf("[DEBUG] ListDeployments response: %v\n", resp)

	return resp, nil
}

// GetDeployment returns deployment by id
func (c *ECClient) GetDeployment(id string) (resp *http.Response, err error) {
	log.Printf("[DEBUG] GetDeployment ID: %s\n", id)

	resourceURL := c.BaseURL + deploymentResource + "/" + id
	authString := "ApiKey " + c.Key
	log.Printf("[DEBUG] GetDeployment Resource URL: %s\n", resourceURL)
	req, err := http.NewRequest("GET", resourceURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", jsonContentType)
	req.Header.Set("Authorization", authString)

	resp, err = c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	log.Printf("[DEBUG] GetDeployment response: %v\n", resp)

	if resp.StatusCode != 200 && resp.StatusCode != 404 {
		respBytes, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("%q: deployment could not be retrieved: %v", id, string(respBytes))
	}

	return resp, nil

}

// CreateDeployment function
func (c *ECClient) CreateDeployment(createDeploymentRequest DeploymentCreateRequest) (createResponse *DeploymentCreateResponse, err error) {
	log.Printf("[DEBUG] CreateDeployment: %v\n", createDeploymentRequest)
	jsonData, err := json.Marshal(createDeploymentRequest)

	if err != nil {
		return nil, err
	}

	jsonString := string(jsonData)
	log.Printf("[DEBUG] CreateDeployment request body: %s\n", jsonString)

	body := strings.NewReader(jsonString)
	resourceURL := c.BaseURL + deploymentResource
	authString := "ApiKey " + c.Key

	log.Printf("[DEBUG] CreateDeployment Resource URL: %s\n", resourceURL)
	req, err := http.NewRequest("POST", resourceURL, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", jsonContentType)
	req.Header.Set("Authorization", authString)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	log.Printf("[DEBUG] CreateDeployment response: %v\n", resp)

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	log.Printf("[DEBUG] CreateDeployment response body: %v\n", string(respBytes))

	err = json.Unmarshal(respBytes, &createResponse)
	if err != nil {
		return nil, err
	}

	return createResponse, nil

}

// DeleteDeployment function
func (c *ECClient) DeleteDeployment(id string) (resp *http.Response, err error) {
	log.Printf("[DEBUG] DeleteDeployment ID: %s\n", id)

	resourceURL := c.BaseURL + deploymentResource + "/" + id + "/_shutdown"
	authString := "ApiKey " + c.Key
	log.Printf("[DEBUG] DeleteDeployment Resource URL: %s\n", resourceURL)
	req, err := http.NewRequest("POST", resourceURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", jsonContentType)
	req.Header.Set("Authorization", authString)

	resp, err = c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	log.Printf("[DEBUG] DeleteDeployment response: %v\n", resp)

	if resp.StatusCode != 200 {
		respBytes, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("%q: Deployment could not be deleted: %v", id, string(respBytes))
	}

	return resp, nil
}

// WaitForElasticsearchDeploymentStatus function to wait required deployment status
func (c *ECClient) WaitForElasticsearchDeploymentStatus(id string, status string, allowMissing bool) error {
	timeoutSeconds := time.Second * time.Duration(600)
	log.Printf("[DEBUG] WaitForElasticsearchDeploymentStatus will wait for %v seconds for '%s' status for deployment ID: %s\n", timeoutSeconds, status, id)

	return resource.Retry(timeoutSeconds, func() *resource.RetryError {
		resp, err := c.GetDeployment(id)

		if err != nil {
			return resource.NonRetryableError(err)
		}

		if resp.StatusCode == 404 && allowMissing {
			return nil
		} else if resp.StatusCode == 200 {

			var clusterInfo DeploymentGetResponse
			err = json.NewDecoder(resp.Body).Decode(&clusterInfo)
			if err != nil {
				return resource.NonRetryableError(err)
			}

			if clusterInfo.Resources.Elasticsearch[0].Info.Status == status {
				log.Printf("[DEBUG] WaitForElasticsearchDeploymentStatus desired deployment status reached: %s\n", clusterInfo.Resources.Elasticsearch[0].Info.Status)
				return nil
			}

			log.Printf("[DEBUG] WaitForElasticsearchDeploymentStatus current deployment status: %s. Desired status: %s\n", clusterInfo.Resources.Elasticsearch[0].Info.Status, status)
		}

		return resource.RetryableError(
			fmt.Errorf("%q: timeout while waiting for the elasticsearch deployment to reach %s status", id, status))
	})

}

// UpdateDeployment function to update existing deployment
func (c *ECClient) UpdateDeployment(id string, updateDeploymentRequest DeploymentUpdateRequest) (updateResponse *DeploymentUpdateResponse, err error) {
	log.Printf("[DEBUG] UpdateDeployment: %v\n", updateDeploymentRequest)
	jsonData, err := json.Marshal(updateDeploymentRequest)

	if err != nil {
		return nil, err
	}

	jsonString := string(jsonData)
	log.Printf("[DEBUG] UpdateDeployment request body: %s\n", jsonString)

	body := strings.NewReader(jsonString)
	resourceURL := c.BaseURL + deploymentResource + "/" + id
	authString := "ApiKey " + c.Key

	log.Printf("[DEBUG] UpdateDeployment Resource URL: %s\n", resourceURL)
	req, err := http.NewRequest("PUT", resourceURL, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", jsonContentType)
	req.Header.Set("Authorization", authString)

	resp, err := c.HTTPClient.Do(req)

	if err != nil {
		return nil, err
	}

	log.Printf("[DEBUG] UpdateDeployment response: %v\n", resp)

	if resp.StatusCode != 200 {
		respBytes, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("%q: Deployment could not be updated: %v", id, string(respBytes))
	}

	respBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	log.Printf("[DEBUG] UpdateDeployment response body: %v\n", string(respBytes))

	err = json.Unmarshal(respBytes, &updateResponse)
	if err != nil {
		return nil, err
	}

	return updateResponse, nil

}

