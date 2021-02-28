package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
)

const baseUpdateURL = "https://admin.gratisdns.com/ddns.php"

type Config struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Domain   string `json:"domain"`
	Hostname string `json:"hostname"`
}

func checkIfUpdateIsNeeded(hostname string) bool {

	updateIsNeeded := false

	resp, err := http.Get("https://api.ipify.org")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// If there is an error looking up the hostname, assume that we need update
	addr, err := net.LookupHost(hostname)
	if err != nil {
		updateIsNeeded = true
	} else {

		if string(body) != addr[0] {
			updateIsNeeded = true
		}
	}

	return updateIsNeeded
}

func genConfigFile(configFile string) {

	var (
		username string
		password string
		domain   string
		host     string
	)

	fmt.Print("Username: ")
	fmt.Scanf("%s\n", &username)

	fmt.Print("DynamicDNS Password: ")
	fmt.Scanf("%s\n", &password)

	fmt.Print("Domain: ")
	fmt.Scanf("%s\n", &domain)

	fmt.Print("Hostname: ")
	fmt.Scanf("%s\n", &host)

	m := Config{username, password, domain, host}
	config, _ := json.Marshal(m)

	_ = ioutil.WriteFile(configFile, config, 0644)
}

func getCurrentDir() string {

	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	return filepath.Dir(ex)
}

func main() {

	configFile := fmt.Sprintf("%s/config.json", getCurrentDir())

	_, err := os.Stat(configFile)
	if err != nil {
		genConfigFile(configFile)
		return
	}

	var config Config

	configFileRead, err := os.Open(configFile)
	if err != nil {
		log.Fatal(err)
	}

	defer configFileRead.Close()

	configByteValue, _ := ioutil.ReadAll(configFileRead)
	json.Unmarshal(configByteValue, &config)

	if checkIfUpdateIsNeeded(*&config.Hostname) {

		updateUriConstruct := fmt.Sprintf("%s?u=%s&p=%s&d=%s&h=%s", baseUpdateURL, *&config.Username, *&config.Password, *&config.Domain, *&config.Hostname)

		resp, err := http.Get(updateUriConstruct)
		if err != nil {
			log.Fatal(err)
		}

		defer resp.Body.Close()
	}
}
