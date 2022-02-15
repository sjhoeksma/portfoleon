# Portfoleon
This is a smal tool to extract view data from https://portfoleon.com in json format so it can be used in other tools

## Steps performed by this tool
* Login using the APIKEY and convert it into Bearer token
* Select the organization by name, if no nama is specified the first organization available is used
* Select the workspace by name, if no name is specified the first workspace available is used
* Load all lookup tables needed to resolve workitems fields and resources
* Select the view by name or all data if not set and convert data in to a single JSON file. In view data you can specify the number (default 0) of last status reports should be added. 
* If DoFieldsLookup is set we will replace all lookup fields IDs with all the data records

## Comandline Options
Use the following commands line agruments
```
portfoleon
  -a string
    	The action(s) that should be performed. (default "View")
  -b string
    	Specify bindAdress. (default ":8080")
  -c int
    	The number of status reports to include in dump. (default -1 )
  -compact
    	Should we only use the values of the lookup fields only. (default true)
  -d	Should we use drafts. (default true)
  -f string
    	Write output to file.
  -k string
    	Specify apiKey. (default "")
  -l	Should we do field lookups. (default true)
  -o string
    	Name of Portfoleon organization to use.
  -serve
    	Use if we should run a webserver.
  -u string
    	Specify baseuUrl towards protfoleon (default "https://portfoleon.herokuapp.com/api/v1")
  -v string
    	Name of Portfoleon view to dump.
  -w string
    	Name of Portfoleon workspace to use.
```

## Environment Variables
You can also use the folowing envrionment variables instead of the commandline arguments
* PORTFOLEON_APIKEY
* PORTFOLEON_BASEURL
* PORTFOLEON_BINDADDRESS
* PORTFOLEON_ORGANIZATION
* PORTFOLEON_WORKSPACE
* PORTFOLEON_VIEWNAME
* PORTFOLEON_STATUSCOUNT

## WebServer
When running a webserver using the -serve option, you should set your APIKEY in the Autorization : bearer APITOKEN. Use the following query values, if defaults should be overwriten
* organization=The organization name
* workspace=The workspace name
* name=The view name or empty for all data
* count=The number of status reports to include
* compact=Return lookup values only
* lookup=Lookup the all fields
