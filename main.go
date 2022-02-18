package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"portfoleon/api"
	"strconv"
	"strings"
	"time"

	strip "github.com/grokify/html-strip-tags-go"
)

var Version = "1.0.1"

//The filename use as output
var fileName = ""

//The filename use as template for output
var templateName = ""

//The JsonFile use to test the template
var jsonFile = ""

//Should we run as webeserver
var webServer bool = false

//The bind Address
var bindAddress = ":8080"

//Variavble stroing the ApiKey to login to portfoleon
var apiKey = ""

//The default organization
var organization = ""

//The default workspace
var workspace = ""

//The default view
var viewName = ""

//The default number of Status counts to include
var statusCount = -1

//The actions witch should be performed
var action = "View"

//Veriable storing the active global token
var token = ""

//Should we do Fields lookup
var doFieldsLookup bool = true

//Should we only Name for lookup values
var onlyLookupName bool = true

//Should we use drafts
var useDrafts bool = true

//The number of gray list days
var grayDays int = 45

//The gray list status
var grayStatus string = ""

//Read os flags: -u <baseUrl> , -k <apiKey>, -ip <bindAddress> ....
//Or use Environment flags PORTFOLEON_APIKEY, PORTFOLEON_BASEURL, PORTFOLEON_BINDADDRESS
func Init() {
	// Read OS variables
	var s = os.Getenv("PORTFOLEON_APIKEY")
	if s != "" {
		apiKey = s
	}
	s = os.Getenv("PORTFOLEON_BASEURL")
	if s != "" {
		api.BaseUrl = s
	}
	s = os.Getenv("PORTFOLEON_BINDADDRESS")
	if s != "" {
		bindAddress = s
	}
	organization = os.Getenv("PORTFOLEON_ORGANIZATION")
	workspace = os.Getenv("PORTFOLEON_WORKSPACE")
	viewName = os.Getenv("PORTFOLEON_VIEWNAME")
	s = os.Getenv("PORTFOLEON_STATUSCOUNT")
	if s != "" {
		statusCount, _ = strconv.Atoi(s)
	}

	// flags declaration using flag package
	version := flag.Bool("version", false, "prints current version ("+Version+")")
	flag.StringVar(&api.BaseUrl, "u", api.BaseUrl, "Specify baseuUrl towards protfoleon")
	flag.StringVar(&apiKey, "k", apiKey, "Specify apiKey.")
	flag.StringVar(&bindAddress, "b", bindAddress, "Specify bindAdress.")
	flag.StringVar(&fileName, "f", fileName, "Write output to file.")
	flag.BoolVar(&webServer, "serve", webServer, "Use if we should run a webserver.")

	flag.StringVar(&organization, "o", organization, "Name of Portfoleon organization to use.")
	flag.StringVar(&workspace, "w", workspace, "Name of Portfoleon workspace to use.")
	flag.StringVar(&viewName, "v", viewName, "Name of Portfoleon view to dump.")
	flag.IntVar(&statusCount, "c", statusCount, "The number of statuses to include in dump.")
	flag.StringVar(&action, "a", action, "The action(s) that should be performed.")
	flag.BoolVar(&doFieldsLookup, "l", doFieldsLookup, "Should we do field lookups.")
	flag.BoolVar(&useDrafts, "d", useDrafts, "Should we use drafts.")
	flag.BoolVar(&onlyLookupName, "compact", onlyLookupName, "Should we only use the values of the lookup fields only.")
	flag.StringVar(&templateName, "t", templateName, "The name of template to use.")
	flag.StringVar(&jsonFile, "tJson", jsonFile, "The name of jsonfile to test the template with.")
	flag.IntVar(&grayDays, "days", grayDays, "The number of days use to graylist a status update .")
	flag.StringVar(&grayStatus, "status", grayStatus, "The status used for graylisting (requires writeable token).")
	flag.Parse() // after declaring flags we need to call it
	if *version {
		fmt.Println("Version ", Version)
		os.Exit(0)
	}
}

//Apply the data to the template
func toTemplate(tplName string, data *string) (string, error) {
	t, err := template.New(filepath.Base(tplName)).Funcs(template.FuncMap{
		"now": time.Now,
		"inc": func(n int) int {
			return n + 1
		},
		"strip": func(html string) string {
			return strip.StripTags(html)
		},
		"slice": func(args ...interface{}) []interface{} {
			return args
		},
	}).ParseFiles(tplName)
	if err != nil {
		return "", err
	}
	tplData := "{\"data\" :" + *data + "}"
	var m map[string]interface{}
	if err := json.Unmarshal([]byte(tplData), &m); err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, m); err != nil {
		return "", err
	}
	return tpl.String(), nil
}

//The webhandler
func webHandlerResponse(response *string, w http.ResponseWriter, r *http.Request, _action string) error {
	var _apiKey = ""
	var _token = token
	var err error
	//Get a valid token for portfoleon on every request
	reqToken := r.Header.Get("Authorization")
	if reqToken != "" {
		splitToken := strings.Split(reqToken, "Bearer ")
		_apiKey = splitToken[1]
		_token = ""
	}
	if _apiKey == "" {
		_apiKey = apiKey
	}
	//Run refresh of token
	if _token != "" && api.RefreshToken(&_token) != nil {
		_token = ""
	}
	//If we don't have a token the create a new one
	if _token == "" {
		//Create an new token
		_token, err = api.GetToken(_apiKey)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintf(w, `{"error", "%s"}`, err)
			return err
		}
		//Store token global if
		if reqToken == "" {
			token = _token
		}
	}
	if _action == "" {
		if r.URL.Path[1:] != "" {
			_action = strings.ReplaceAll(r.URL.Path[1:], "/", ",")
		} else {
			_action = action
		}
	}
	_organization := r.URL.Query().Get("organization")
	if _organization == "" {
		_organization = organization
	}
	_workspace := r.URL.Query().Get("workspace")
	if _workspace == "" {
		_workspace = workspace
	}
	_viewName := r.URL.Query().Get("name")
	if _viewName == "" {
		_viewName = viewName
	}
	var _statusCount = statusCount
	s := r.URL.Query().Get("count")
	if s != "" {
		_statusCount, _ = strconv.Atoi(s)
	}
	var _doFieldsLookup = doFieldsLookup
	s = r.URL.Query().Get("lookup")
	if s != "" {
		_doFieldsLookup, _ = strconv.ParseBool(s)
	}
	var _onlyLookupName = onlyLookupName
	s = r.URL.Query().Get("compact")
	if s != "" {
		_onlyLookupName, _ = strconv.ParseBool(s)
	}
	var _drafts = useDrafts
	s = r.URL.Query().Get("drafts")
	if s != "" {
		_drafts, _ = strconv.ParseBool(s)
	}
	err = api.GetAction(response, _token, _action, _organization, _workspace,
		_viewName, _statusCount, _doFieldsLookup, _onlyLookupName, _drafts)
	return err
}

//The handler for web requests
func webHandler(w http.ResponseWriter, r *http.Request) {
	var _tplName = filepath.Base(r.URL.Query().Get("template"))
	if _tplName == "" {
		_tplName = templateName
	}
	var response = ""
	err := webHandlerResponse(&response, w, r, r.URL.Query().Get("action"))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, `{"error", "%s"}`, err)
	} else {
		if _tplName != "" {
			response, err = toTemplate(_tplName, &response)
			if err != nil {
				log.Fatal("Template processing Error", err)
			}
		} else {
			//When we don't use a template output is json
			w.Header().Set("Content-Type", "application/json")
		}
		//Write out the  data
		fmt.Fprint(w, response)
	}
}

func main() {
	Init()
	if webServer {
		//WebServer mode
		//Simple webserver to reponse on request with the requested workspace items
		http.HandleFunc("/", webHandler)
		log.Println("Starting API servering on", bindAddress)
		log.Fatal(http.ListenAndServe(bindAddress, nil))
	} else {
		//Output to console or file
		token, err := api.GetToken(apiKey)
		if err != nil || token == "" {
			log.Fatal("Login failed", err)
		}
		var response string = ""
		if grayDays != 0 && grayStatus != "" {
			r, err := api.DoGrayListing(token, action, organization, workspace, viewName, grayStatus, grayDays)
			if err != nil {
				log.Fatal(err)
			}
			j, _ := json.Marshal(r)
			response = string(j)
		} else if jsonFile != "" {
			b, err := ioutil.ReadFile(jsonFile) // just pass the file name
			if err != nil {
				log.Fatal(err)
			}
			response = string(b) // convert content to a 'string'
		} else {
			err = api.GetAction(&response, token, action, organization, workspace, viewName,
				statusCount, doFieldsLookup, onlyLookupName, useDrafts)
			if err != nil {
				log.Fatal(err)
			}
		}

		//Should we run the response trough a template
		if templateName != "" {
			response, err = toTemplate(templateName, &response)
			if err != nil {
				log.Fatal(err)
			}
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
