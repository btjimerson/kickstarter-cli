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

//VirtualService - structure for creating a Virtual Service
type VirtualService struct {
	Directory             string `json:"-"`
	Name                  string `json:"name,omitempty"`
	Namespace             string `json:"namespace,omitempty"`
	Host                  string `json:"host,omitempty"`
	GatewayName           string `json:"gatewayName,omitempty"`
	DestinationHost       string `json:"destinationHost,omitempty"`
	DestinationPortNumber int64  `json:"destinationPortNumber,omitempty"`
}

//GetVirtualService - gets a virtual service object
func GetVirtualService(vs VirtualService) {

	util.CheckAPIURL()

	reqBody, err := json.Marshal(vs)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	url, err := util.GetAPIURL()
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", url+"/virtualService", bytes.NewBuffer([]byte(reqBody)))
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
	if len(vs.Directory) == 0 {
		directory, _ = os.UserHomeDir()
	} else {
		directory = vs.Directory
	}
	filename := directory + "/" + vs.Name + "-virtual-service.yaml"
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	file.Write(respBody)

	fmt.Println("Saved virtual service to", filename)

}
