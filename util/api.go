package util

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

//SetAPIURL sets the url of the api
func SetAPIURL(url string) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	ksDir := homeDir + string(os.PathSeparator) + ".kickstarter"
	if _, err = os.Stat(ksDir); os.IsNotExist(err) {
		os.Mkdir(ksDir, 0755)
	}

	ksConfig := ksDir + string(os.PathSeparator) + "config"

	apiConfig := "api=" + url
	os.WriteFile(ksConfig, []byte(apiConfig), 0755)

	fmt.Println("Saved API URL as", url)

}

//GetAPIURL gets the url of the api
func GetAPIURL() (string, error) {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	ksDir := homeDir + string(os.PathSeparator) + ".kickstarter"
	if _, err = os.Stat(ksDir); os.IsNotExist(err) {
		return "", err
	}

	ksConfig := ksDir + string(os.PathSeparator) + "config"
	if _, err = os.Stat(ksConfig); os.IsNotExist(err) {
		return "", err
	}

	file, err := os.Open(ksConfig)
	if err != nil {
		return "", err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var config []string
	for scanner.Scan() {
		config = append(config, scanner.Text())
	}

	var url string
	for _, value := range config {
		if strings.Contains(value, "api=") {
			cutString := strings.FieldsFunc(value, func(r rune) bool {
				return r == '='
			})
			url = cutString[1]
		}
	}

	return url, nil

}

func isAPIURLSet() bool {
	url, err := GetAPIURL()
	if err != nil || len(url) == 0 {
		return false
	}

	return true

}

//CheckAPIURL checks that the url of the api has been set
func CheckAPIURL() {

	if !isAPIURLSet() {
		fmt.Println("No API URL has been set.  Set the URL with the kickstarter api command first.")
		os.Exit(1)
	}
}
