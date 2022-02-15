package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"portfoleon/api"
	"strconv"
)

var fileName = ""
var webServer bool = false

//The bind Address
var BindAddress = ":8080"

//Read os flags: -u <baseUrl> , -k <apiKey>, -ip <bindAddress> ....
//Or use Environment flags PORTFOLEON_APIKEY, PORTFOLEON_BASEURL, PORTFOLEON_BINDADDRESS
func Init() {
	// Read OS variables
	var s = os.Getenv("PORTFOLEON_APIKEY")
	if s != "" {
		api.ApiKey = s
	}
	s = os.Getenv("PORTFOLEON_BASEURL")
	if s != "" {
		api.BaseUrl = s
	}
	s = os.Getenv("PORTFOLEON_BINDADDRESS")
	if s != "" {
		BindAddress = s
	}
	api.Organization = os.Getenv("PORTFOLEON_ORGANIZATION")
	api.Workspace = os.Getenv("PORTFOLEON_WORKSPACE")
	api.ViewName = os.Getenv("PORTFOLEON_VIEWNAME")
	s = os.Getenv("PORTFOLEON_STATUSCOUNT")
	if s != "" {
		api.StatusCount, _ = strconv.Atoi(s)
	}

	// flags declaration using flag package
	flag.StringVar(&api.BaseUrl, "u", api.BaseUrl, "Specify baseuUrl towards protfoleon")
	flag.StringVar(&api.ApiKey, "k", api.ApiKey, "Specify apiKey.")
	flag.StringVar(&BindAddress, "b", BindAddress, "Specify bindAdress.")
	flag.StringVar(&fileName, "f", fileName, "Write output to file.")
	flag.BoolVar(&webServer, "serve", webServer, "Use if we should run a webserver.")

	flag.StringVar(&api.Organization, "o", api.Organization, "Name of Portfoleon organization to use.")
	flag.StringVar(&api.Workspace, "w", api.Workspace, "Name of Portfoleon workspace to use.")
	flag.StringVar(&api.ViewName, "v", api.ViewName, "Name of Portfoleon view to dump.")
	flag.IntVar(&api.StatusCount, "c", api.StatusCount, "The number of statuses to include in dump.")
	flag.StringVar(&api.Action, "a", api.Action, "The action(s) that should be performed.")
	flag.BoolVar(&api.DoFieldsLookup, "l", api.DoFieldsLookup, "Should we do field lookups.")
	flag.BoolVar(&api.UseDrafts, "d", api.UseDrafts, "Should we use drafts.")
	flag.BoolVar(&api.OnlyLookupName, "compact", api.OnlyLookupName, "Should we only use the values of the lookup fields only.")

	flag.Parse() // after declaring flags we need to call it
}

func main() {
	Init()
	if webServer {
		//WebServer mode
		//Simple webserver to reponse on request with the requested workspace items
		http.HandleFunc("/", api.WebHandler)
		log.Println("Starting API servering on", BindAddress)
		log.Fatal(http.ListenAndServe(BindAddress, nil))
	} else {
		//Output to console or file
		token, err := api.GetToken(api.ApiKey)
		if err != nil || token == "" {
			log.Fatal("Login failed", err)
		}
		var response string = ""
		err = api.GetAction(&response, token, api.Action, api.Organization, api.Workspace, api.ViewName,
			api.StatusCount, api.DoFieldsLookup, api.OnlyLookupName)
		if err != nil {
			log.Fatal(err)
		}

		if fileName != "" {
			f, err := os.Create(fileName)
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
			_, err = f.WriteString(response + "\n")
			if err != nil {
				log.Fatal(err)
			}
		} else {
			fmt.Println(response)
		}
	}
}
