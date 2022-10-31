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

//Service - structure for creating a Service
type Service struct {
	Directory string `json:"-"`
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	Port      int64  `json:"port,omitempty"`
}

//GetService - gets a service object
func GetService(s Service) {

	util.CheckAPIURL()

	reqBody, err := json.Marshal(s)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	url, err := util.GetAPIURL()
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", url+"/service", bytes.NewBuffer([]byte(reqBody)))
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
	if len(s.Directory) == 0 {
		directory, _ = os.UserHomeDir()
	} else {
		directory = s.Directory
	}
	filename := directory + "/" + s.Name + "-service.yaml"
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	file.Write(respBody)

	fmt.Println("Saved service to", filename)

}
