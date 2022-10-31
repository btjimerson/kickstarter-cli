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

//Gateway - structure for creating a Gateway
type Gateway struct {
	Directory      string `json:"-"`
	Name           string `json:"name,omitempty"`
	Namespace      string `json:"namespace,omitempty"`
	Port           int64  `json:"port,omitempty"`
	PortName       string `json:"portName,omitempty"`
	Protocol       string `json:"protocol,omitempty"`
	Host           string `json:"host,omitempty"`
	TLSMode        string `json:"tlsMode,omitempty"`
	CredentialName string `json:"credentialName,omitempty"`
}

//GetGateway - gets a gateway object
func GetGateway(gw Gateway) {

	util.CheckAPIURL()

	reqBody, err := json.Marshal(gw)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	url, err := util.GetAPIURL()
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", url+"/gateway", bytes.NewBuffer([]byte(reqBody)))
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
	if len(gw.Directory) == 0 {
		directory, _ = os.UserHomeDir()
	} else {
		directory = gw.Directory
	}
	filename := directory + "/" + gw.Name + "-gateway.yaml"
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	file.Write(respBody)

	fmt.Println("Saved gateway to", filename)

}
