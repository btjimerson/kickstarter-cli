package types

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"net/http"

	"aspenmesh/kickstarter/util"
)

//Deployment - structure for creating a Deployment
type Deployment struct {
	Directory          string `json:"-"`
	Name               string `json:"name,omitempty"`
	Namespace          string `json:"namespace,omitempty"`
	ContainerImage     string `json:"containerImage,omitempty"`
	ContainerName      string `json:"containerName,omitempty"`
	ImagePullPolicy    string `json:"imagePullPolicy,omitempty"`
	DeploymentStrategy string `json:"deploymentStrategy,omitempty"`
	ServiceAccount     string `json:"serviceAccount,omitempty"`
	Replicas           int64  `json:"replicas,omitempty"`
	ContainerPort      int64  `json:"containerPort,omitempty"`
}

//GetDeployment - gets a deployment object
func GetDeployment(d Deployment) {

	util.CheckAPIURL()

	reqBody, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	url, err := util.GetAPIURL()
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", url+"/deployment", bytes.NewBuffer([]byte(reqBody)))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var directory string
	if len(d.Directory) == 0 {
		directory, _ = os.UserHomeDir()
	} else {
		directory = d.Directory
	}

	filename := directory + "/" + d.Name + "-deployment.yaml"
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	file.Write(respBody)

	fmt.Println("Saved deployment to", filename)

}
