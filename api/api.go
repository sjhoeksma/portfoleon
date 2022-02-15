//The package gives access to portfoleon use the api stack.
//We using the Serve mode you can read all workitems
package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

//The Token Interface structure for a token request of portfoleon
type TokenInterface struct {
	B64token  string `json:"b64token"`
	Email     string `json:"email"`
	Firstname string `json:"firstName"`
	Fullname  string `json:"fullName"`
	Id        int    `json:"id"`
	Token     string `json:"token"`
}

type OrganizationInterface struct {
	BillingCustomerCode   interface{} `json:"billing_customer_code"`
	BillingPlanCode       string      `json:"billing_plan_code"`
	BillingPlanExpiryDate interface{} `json:"billing_plan_expiry_date"`
	BillingPlanStartDate  string      `json:"billing_plan_start_date"`
	CallStatusCode        interface{} `json:"call_status_code"`
	FinAdminEmail         interface{} `json:"fin_admin_email"`
	ID                    int         `json:"id"`
	IsTest                bool        `json:"is_test"`
	LiveAvailable         bool        `json:"live_available"`
	LiveOrganizationID    interface{} `json:"live_organization_id"`
	Name                  string      `json:"name"`
	RoleTypeCode          string      `json:"role_type_code"`
	ShortName             interface{} `json:"short_name"`
	TestOrganizationID    int         `json:"test_organization_id"`
	URL                   interface{} `json:"url"`
}

type WorkspaceInterface struct {
	ID             int         `json:"id"`
	Name           string      `json:"name"`
	OrganizationID int         `json:"organization_id"`
	RoleTypeCode   string      `json:"role_type_code"`
	Settings       interface{} `json:"settings"`
}

//The StatusReportInterface
type StatusReportInterface struct {
	DtReport                 string      `json:"dt_report"`
	DtSubmitted              string      `json:"dt_submitted"`
	ID                       int         `json:"id"`
	Latest                   bool        `json:"latest"`
	PercentComplete          interface{} `json:"percent_complete"`
	PercentCompleteUnchanged float64     `json:"percent_complete_unchanged"`
	Report                   string      `json:"report"`
	StatusColor              string      `json:"status_color"`
	StatusID                 int         `json:"status_id"`
	StatusIDUnchanged        bool        `json:"status_id_unchanged"`
	StatusName               string      `json:"status_name"`
	UserID                   int         `json:"user_id"`
	UserName                 string      `json:"user_name"`
}

//The StatusInterface
type StatusInterface struct {
	Color          string `json:"color"`
	ID             int    `json:"id"`
	Name           string `json:"name"`
	OrganizationID int    `json:"organization_id"`
	ValueOrder     int    `json:"value_order"`
}

//The lookup talbe for status
type StatusLookupInterface map[int]StatusInterface

//The WorkItem interface
type WorkItemInterface struct {
	AvgFte float64 `json:"avg_fte"`
	//	ChildCount                int                    `json:"child_count"`
	//	ChildOpenCount            int                    `json:"child_open_count"`
	Code    int    `json:"code"`
	Draft   bool   `json:"draft"`
	DtEnd   string `json:"dt_end"`
	DtStart string `json:"dt_start"`
	//	ExternalSystemConnectorID interface{}            `json:"external_system_connector_id"`
	//	ExternalSystemItemID      interface{}            `json:"external_system_item_id"`
	//	ExternalSystemItemTypeID  interface{}            `json:"external_system_item_type_id"`
	Fields map[string]interface{} `json:"fields"`
	ID     int                    `json:"id"`
	//	LatestRevisionID int                    `json:"latest_revision_id"`
	//	Level                     int                    `json:"level"`
	//	Links                     []interface{}          `json:"links"`
	Name             string      `json:"name"`
	ParentWorkItemID interface{} `json:"parent_work_item_id"`
	//	Path                interface{}   `json:"path"`
	PercentComplete     interface{}   `json:"percent_complete"`
	Phases              []interface{} `json:"phases"`
	ResourceIds         []interface{} `json:"resource_ids"`
	ResourceLocationIds []interface{} `json:"resource_location_ids"`
	ResourceRoleIds     []interface{} `json:"resource_role_ids"`
	ResourceTeamIds     []interface{} `json:"resource_team_ids"`
	DtReport            string        `json:"dt_report"`
	StatusID            int           `json:"status_id"`
	Status              interface{}
	StatusReport        string
	StatusReports       []StatusReportInterface
	Tags                []interface{} `json:"tags"`
	TotalEffort         float64       `json:"total_effort"`
	TrackedHours        int           `json:"tracked_hours"`
	WorkItemTypeID      int           `json:"work_item_type_id"`
	WorkItemType        interface{}
	//	WorkspaceID               int         `json:"workspace_id"`
	//	WorkzoneID                interface{} `json:"workzone_id"`
}

//The ViewInterface
type ViewInterface struct {
	ID           int    `json:"id"`
	IsPrivate    bool   `json:"is_private"`
	Name         string `json:"name"`
	ViewSettings struct {
		TableSettings struct {
			ColumnSettings []struct {
				FieldName string `json:"field_name"`
				IsVisible bool   `json:"is_visible"`
				Width     int    `json:"width"`
			} `json:"column_settings"`
		} `json:"table_settings"`
	} `json:"view_settings,omitempty"`
}

type ViewLookupInterface map[string]ViewInterface

type FieldValueInterface struct {
	FieldID    int    `json:"field_id"`
	ID         int    `json:"id"`
	IsEnabled  bool   `json:"is_enabled"`
	Name       string `json:"name"`
	ValueOrder int    `json:"value_order"`
}

//The fieldsInterface
type FieldsInterface struct {
	Ascending         bool                  `json:"ascending"`
	Caption           string                `json:"caption"`
	DataTypeCode      string                `json:"data_type_code"`
	ID                int                   `json:"id"`
	IsEnabled         bool                  `json:"is_enabled"`
	Name              string                `json:"name"`
	SelectValues      []FieldValueInterface `json:"selectValues"`
	UnitOfMeasurement interface{}           `json:"unit_of_measurement"`
	WorkspaceID       int                   `json:"workspace_id"`
}

type FieldsLookupInterface map[string]FieldsInterface

//Resources interface
type ResourcesInterface struct {
	DtEnd                interface{} `json:"dt_end"`
	DtStart              string      `json:"dt_start"`
	Email                interface{} `json:"email"`
	FteLimit             float64     `json:"fte_limit"`
	HasGaps              interface{} `json:"has_gaps"`
	HasOverloads         interface{} `json:"has_overloads"`
	ID                   int         `json:"id"`
	IsEnabled            bool        `json:"is_enabled"`
	ManagerUserID        interface{} `json:"manager_user_id"`
	Name                 string      `json:"name"`
	OrganizationID       int         `json:"organization_id"`
	PersonnelBudgetSlots string      `json:"personnel_budget_slots"`
	PersonnelCode        interface{} `json:"personnel_code"`
	ResourceLocationID   int         `json:"resource_location_id"`
	ResourceRoleIds      []int       `json:"resource_role_ids"`
	ResourceTeamID       int         `json:"resource_team_id"`
	UserID               interface{} `json:"user_id"`
}

type ResourcesLookupInterface map[int]ResourcesInterface

type WorkItemTypeInterface struct {
	Code string `json:"code"`
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type WorkItemTypeLookupInterface map[int]WorkItemTypeInterface

//The bind Address
var BindAddress = ":8080"

//Variable storing the baseURL to access portfoleon api
var BaseUrl = "https://portfoleon.herokuapp.com/api/v1"

//Variavble stroing the ApiKey to login to portfoleon
var ApiKey = ""

//The default organization
var Organization = ""

//The default workspace
var Workspace = ""

//The default view
var ViewName = ""

//The default number of Status counts to include
var StatusCount = -1

//The actions witch should be performed
var Action = "View"

//Veriable storing the active global token
var token = ""

//Should we do Fields lookup
var DoFieldsLookup bool = true

//Should we only Name for lookup values
var OnlyLookupName bool = true

//Should we use drafts
var UseDrafts bool = true

//Create a bearer token by logging in useing the apiKey
func GetToken(apiKey string) (string, error) {
	var jsonData = []byte(`{
		"api_key": "` + apiKey + `"
	}`)
	request, err := http.NewRequest("POST", BaseUrl+"/security/token", bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	if err != nil {
		return "", err
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	var responseObject TokenInterface
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return "", err
	}
	if len(responseObject.B64token) == 0 {
		return "", errors.New("Invalid loading: " + string(responseData))
	}
	return responseObject.B64token, nil
}

//Refresh the bearer token
func RefreshToken(token *string) error {
	if *token == "" {
		return nil
	}
	request, err := http.NewRequest("GET", BaseUrl+"/security/refresh_token", nil)
	request.Header.Add("Authorization", "Bearer "+*token)
	if err != nil {
		return err
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	var responseObject TokenInterface
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return err
	}
	if len(responseObject.B64token) == 0 {
		return errors.New("Invalid loading: " + string(responseData))
	}
	*token = responseObject.B64token
	return nil
}

//Get the organisation beloging by name
func GetOrganization(token string, name string) (int, error) {
	request, err := http.NewRequest("GET", BaseUrl+"/organizations", nil)
	request.Header.Add("Authorization", "Bearer "+token)
	if err != nil {
		return 0, err
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, err
	}
	type OrganizationResponse struct {
		Data []OrganizationInterface `json:"data"`
	}
	var responseObject OrganizationResponse
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return 0, err
	}
	for _, s := range responseObject.Data {
		if s.Name == name || s.ShortName == name || name == "" {
			return s.ID, nil
		}
	}
	return 0, errors.New("Organization " + name + " not found")
}

//Get the workspace for the given organization
func GetWorkspace(token string, organization int, name string) (int, error) {
	request, err := http.NewRequest("GET", BaseUrl+"/workspaces?organization_id="+strconv.Itoa(organization), nil)
	request.Header.Add("Authorization", "Bearer "+token)
	if err != nil {
		return 0, err
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, err
	}
	type WorkspaceResponse struct {
		Data []WorkspaceInterface `json:"data"`
	}
	var responseObject WorkspaceResponse
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return 0, err
	}
	for _, s := range responseObject.Data {
		if s.Name == name || name == "" {
			return s.ID, nil
		}
	}
	return 0, errors.New("Workspace " + name + " not found")
}

//Create a lookup table for status
func GetStatusLookUp(token string, organization int, workspace int) (StatusLookupInterface, error) {
	var lookup StatusLookupInterface = make(StatusLookupInterface)

	request, err := http.NewRequest("GET", BaseUrl+"/organizations/"+strconv.Itoa(organization)+"/statuses", nil)
	request.Header.Add("Authorization", "Bearer "+token)
	if err != nil {
		return lookup, err
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return lookup, err
	}
	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return lookup, err
	}
	type StatusResponse struct {
		Data []StatusInterface `json:"data"`
	}
	var responseObject StatusResponse
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return lookup, err
	}
	for _, s := range responseObject.Data {
		lookup[s.ID] = s
	}
	return lookup, nil
}

//Create a lookup table for status
func GetWorkItemTypeLookUp(token string, organization int, workspace int) (WorkItemTypeLookupInterface, error) {
	var lookup WorkItemTypeLookupInterface = make(WorkItemTypeLookupInterface)

	request, err := http.NewRequest("GET", BaseUrl+"/work_item_types?workspace_id="+strconv.Itoa(workspace), nil)
	request.Header.Add("Authorization", "Bearer "+token)
	if err != nil {
		return lookup, err
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return lookup, err
	}
	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return lookup, err
	}
	type WorkItemTypeResponse struct {
		Data []WorkItemTypeInterface `json:"data"`
	}
	var responseObject WorkItemTypeResponse
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return lookup, err
	}
	for _, s := range responseObject.Data {
		lookup[s.ID] = s
	}
	return lookup, nil
}

//Create a lookup table for status
func GetFieldsLookUp(token string, organization int, workspace int) (FieldsLookupInterface, error) {
	var lookup FieldsLookupInterface = make(FieldsLookupInterface)

	request, err := http.NewRequest("GET", BaseUrl+"/fields?workspace_id="+strconv.Itoa(workspace), nil)
	request.Header.Add("Authorization", "Bearer "+token)
	if err != nil {
		return lookup, err
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return lookup, err
	}
	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return lookup, err
	}
	type FieldsResponse struct {
		Data []FieldsInterface `json:"data"`
	}
	var responseObject FieldsResponse
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return lookup, err
	}
	for _, s := range responseObject.Data {
		lookup[s.Name] = s
	}
	return lookup, nil
}

//Create a lookup table for status
func GetResourcesLookUp(token string, organization int, workspace int) (ResourcesLookupInterface, error) {
	var lookup ResourcesLookupInterface = make(ResourcesLookupInterface)

	request, err := http.NewRequest("GET", BaseUrl+"/resources?organization_id="+strconv.Itoa(organization), nil)
	request.Header.Add("Authorization", "Bearer "+token)
	if err != nil {
		return lookup, err
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return lookup, err
	}
	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return lookup, err
	}
	type ResourcesResponse struct {
		Data []ResourcesInterface `json:"data"`
	}
	var responseObject ResourcesResponse
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return lookup, err
	}
	for _, s := range responseObject.Data {
		lookup[s.ID] = s
	}
	return lookup, nil
}

//Get the view lookup for the given organization
func GetViewLookup(token string, organization int, workspace int) (ViewLookupInterface, error) {
	var lookup ViewLookupInterface = make(ViewLookupInterface)
	request, err := http.NewRequest("GET", BaseUrl+"/views?workspace_id="+strconv.Itoa(workspace), nil)
	request.Header.Add("Authorization", "Bearer "+token)
	if err != nil {
		return lookup, err
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return lookup, err
	}
	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return lookup, err
	}
	type ViewResponse struct {
		Data []ViewInterface `json:"data"`
	}
	var responseObject ViewResponse
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return lookup, err
	}
	for _, s := range responseObject.Data {
		lookup[s.Name] = s
	}
	return lookup, nil
}

//Get the status reports
func GetStatusReports(token string, workitem int, count int) ([]StatusReportInterface, error) {
	var ret []StatusReportInterface

	request, err := http.NewRequest("GET", BaseUrl+"/work_items/"+strconv.Itoa(workitem)+"/status_reports", nil)
	request.Header.Add("Authorization", "Bearer "+token)
	if err != nil {
		return ret, err
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return ret, err
	}
	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return ret, err
	}
	type reportResponse struct {
		Data []StatusReportInterface `json:"data"`
	}
	var responseObject reportResponse
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return ret, err
	}

	for i, o := range responseObject.Data {
		//Limit the list if needed
		if count != 0 && i > count {
			break
		}
		ret = append(ret, o)
	}
	return ret, nil
}

//Get all the workitems for a given view within a workspace
func GetWorkItems(token string, organization int, workspace int, viewName string, statusCount int, drafts bool, doFieldsLookup bool, onlyLookupName bool) (string, error) {
	var addDrafts = ""
	if drafts {
		addDrafts = "drafts=true"
	}
	var request *http.Request = nil
	var err error = nil

	lookupView, err := GetViewLookup(token, organization, workspace)
	if err != nil {
		return "", err
	}
	var view = lookupView[viewName]
	if view.ID == 0 {
		request, err = http.NewRequest("GET", BaseUrl+"/work_items?workspace_id="+strconv.Itoa(workspace)+"&"+addDrafts, nil)
	} else {
		request, err = http.NewRequest("GET", BaseUrl+"/views/"+strconv.Itoa(view.ID)+"/work_items?"+addDrafts, nil)
	}
	request.Header.Add("Authorization", "Bearer "+token)
	if err != nil {
		return "", err
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	type WorkItemsResponse struct {
		Data     []WorkItemInterface `json:"data"`
		Page     interface{}         `json:"page"`
		PageSize interface{}         `json:"page_size"`
		Pages    interface{}         `json:"pages"`
		Total    interface{}         `json:"total"`
	}
	var responseObject WorkItemsResponse
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return "", err
	}

	//Create the lookup WorkItemType table
	lookupWorkItemType, err := GetWorkItemTypeLookUp(token, organization, workspace)
	if err != nil {
		return "", err
	}

	//Create the lookupStatus table
	lookupStatus, err := GetStatusLookUp(token, organization, workspace)
	if err != nil {
		return "", err
	}
	//Create the lookupFields table
	lookupFields, err := GetFieldsLookUp(token, organization, workspace)
	if err != nil {
		return "", err
	}
	//Create the lookupStatus table
	lookupResources, err := GetResourcesLookUp(token, organization, workspace)
	if err != nil {
		return "", err
	}

	for i, o := range responseObject.Data {
		//If we need to add the status reports then add them
		if statusCount >= 0 {
			//Get the statusReports for the object
			var statusReports, err = GetStatusReports(token, o.ID, statusCount)
			if err != nil {
				return "", err
			}
			if !onlyLookupName || statusCount != 1 {
				responseObject.Data[i].StatusReports = statusReports
			}
			if len(statusReports) > 0 {
				responseObject.Data[i].StatusReport = statusReports[0].Report
			}
		}
		//Fill the status field with the result of the status_id
		if onlyLookupName {
			responseObject.Data[i].Status = lookupStatus[responseObject.Data[i].StatusID].Name
			responseObject.Data[i].WorkItemType = lookupWorkItemType[responseObject.Data[i].WorkItemTypeID].Name
		} else {
			responseObject.Data[i].Status = lookupStatus[responseObject.Data[i].StatusID]
			responseObject.Data[i].WorkItemType = lookupWorkItemType[responseObject.Data[i].WorkItemTypeID]
		}

		//Removed fields if they are not visible
		if view.ID != 0 {
			for _, v := range view.ViewSettings.TableSettings.ColumnSettings {
				for nf := range o.Fields {
					if nf == v.FieldName && !v.IsVisible {
						delete(responseObject.Data[i].Fields, nf)
						break
					}
				}
			}
		}

		//Replace all fields lookup values
		if doFieldsLookup {
			for n, v := range o.Fields {
				if v == nil {
					continue
				}
				lookup := lookupFields[n]
				if lookup.DataTypeCode == "enum" {
					//For Enum do lookup of Value in the selected values
					for _, lv := range lookup.SelectValues {
						if lv.ID == int(v.(float64)) {
							if onlyLookupName {
								responseObject.Data[i].Fields[n] = lv.Name
							} else {
								responseObject.Data[i].Fields[n] = lv
							}
						}
					}
				} else

				//resource
				if lookup.DataTypeCode == "resource" {
					var resource ResourcesInterface = lookupResources[int(v.(float64))]
					if onlyLookupName {
						responseObject.Data[i].Fields[n] = resource.Name
					} else {
						responseObject.Data[i].Fields[n] = resource
					}
				}

				//tags
				if lookup.DataTypeCode == "tag" && len(v.([]interface{})) != 0 {
					var tags []interface{}
					for _, tv := range v.([]interface{}) {
						for _, fv := range lookup.SelectValues {
							if fv.ID == int(tv.(float64)) {
								if onlyLookupName {
									tags = append(tags, fv.Name)
								} else {
									tags = append(tags, fv)
								}
								break
							}
						}
					}
					responseObject.Data[i].Fields[n] = tags
				}
			}
		}
	}

	//TODO Check for page sizes

	ret, err := json.Marshal(responseObject.Data)
	return string(ret), err
}

func GetAction(response *string, token string, action string, organization string, workspace string,
	viewName string, statusCount int, doFieldsLookup bool, onlyLookupName bool) error {
	action = strings.ToUpper(action)
	type actionReponseInterface struct {
		action   string
		response string
	}
	var responseRec []actionReponseInterface

	if token == "" {
		return errors.New("token is not valid for action")
	}
	orgId, err := GetOrganization(token, organization)
	if err != nil {
		return err
	}
	spaceId, err := GetWorkspace(token, orgId, workspace)
	if err != nil {
		return err
	}

	if strings.Contains(action, "VIEW") {
		var r actionReponseInterface
		r.action = "VIEW"
		r.response, err = GetWorkItems(token, orgId, spaceId, viewName, statusCount, UseDrafts,
			doFieldsLookup, onlyLookupName)
		if err != nil || token == "" {
			return err
		}
		responseRec = append(responseRec, r)
	}

	if strings.Contains(action, "USERS") {
		var r actionReponseInterface
		r.action = "USERS"
		lookup, _ := GetResourcesLookUp(token, orgId, spaceId)
		ret, _ := json.Marshal(lookup)
		r.response = string(ret)
		responseRec = append(responseRec, r)
	}

	if strings.Contains(action, "STATUS") {
		var r actionReponseInterface
		r.action = "STATUS"
		lookup, _ := GetStatusLookUp(token, orgId, spaceId)
		ret, _ := json.Marshal(lookup)
		r.response = string(ret)
		responseRec = append(responseRec, r)
	}

	if strings.Contains(action, "FIELDS") {
		var r actionReponseInterface
		r.action = "FIELDS"
		lookup, _ := GetFieldsLookUp(token, orgId, spaceId)
		ret, _ := json.Marshal(lookup)
		r.response = string(ret)
		responseRec = append(responseRec, r)
	}

	if len(responseRec) == 0 {
		return errors.New("No actions found in " + action)
	}

	if len(responseRec) == 1 {
		*response = responseRec[0].response
	} else {
		*response = "{"
		for i, s := range responseRec {
			if i > 0 {
				*response += ","
			}
			*response += "\"" + s.action + "\" : \"" + s.response + "\""
		}
		*response += "}"
	}
	return nil
}

//The handler for web requests
func apiHandler(w http.ResponseWriter, r *http.Request) {
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
		_apiKey = ApiKey
	}
	//Run refresh of token
	if _token != "" && RefreshToken(&_token) != nil {
		_token = ""
	}
	//If we don't have a token the create a new one
	if _token == "" {
		//Create an new token
		_token, err = GetToken(_apiKey)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintf(w, `{"error", "%s"}`, err)
			return
		}
		//Store token global if
		if reqToken == "" {
			token = _token
		}
	}
	action := r.URL.Query().Get("action")
	if action == "" {
		if r.URL.Path[1:] != "" {
			action = strings.ReplaceAll(r.URL.Path[1:], "/", ",")
		} else {
			action = Action
		}
	}
	organization := r.URL.Query().Get("organization")
	if organization == "" {
		organization = Organization
	}
	workspace := r.URL.Query().Get("workspace")
	if workspace == "" {
		workspace = Workspace
	}
	viewName := r.URL.Query().Get("name")
	if viewName == "" {
		viewName = ViewName
	}
	var statusCount = StatusCount
	s := r.URL.Query().Get("count")
	if s != "" {
		statusCount, _ = strconv.Atoi(s)
	}
	var doFieldsLookup = DoFieldsLookup
	s = r.URL.Query().Get("lookup")
	if s != "" {
		doFieldsLookup, _ = strconv.ParseBool(s)
	}
	var onlyLookupName = OnlyLookupName
	s = r.URL.Query().Get("compact")
	if s != "" {
		onlyLookupName, _ = strconv.ParseBool(s)
	}
	var response = ""
	err = GetAction(&response, token, action, organization, workspace,
		viewName, statusCount, doFieldsLookup, onlyLookupName)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, `{"error", "%s"}`, err)
	} else {
		//Write out the  data
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, response)
	}
}

//Simple webserver to reponse on request with the requested workspace items
func Serve() {
	http.HandleFunc("/", apiHandler)
	log.Println("Starting API servering on", BindAddress)
	log.Fatal(http.ListenAndServe(BindAddress, nil))
}
