package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
)

const baseUpdateURL = "https://admin.gratisdns.com/ddns.php"

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

func main() {

	username := flag.String("u", "username", "GratisDNS Username")
	password := flag.String("p", "password", "GratisDNS DynamicDNS Password")
	domainName := flag.String("d", "example.com", "Domain where the hostname is located")
	hostName := flag.String("h", "ddns.example.com", "Dynamic DNS hostname relativ")
	flag.Parse()

	if checkIfUpdateIsNeeded(*hostName) {

		updateUriConstruct := fmt.Sprintf("%s?u=%s&p=%s&d=%s&h=%s", baseUpdateURL, *username, *password, *domainName, *hostName)

		resp, err := http.Get(updateUriConstruct)
		if err != nil {
			log.Fatal(err)
		}

		defer resp.Body.Close()
	}
}
