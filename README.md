# Portfoleon
This is a smal tool to extract view data from https://portfoleon.com in json format so it can be used in other tools. 
You can download the latest stable version for you platform from the dist directroy and used gzip -d `protfoleon-...` to
get you nativefile

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
  -days int
    	The number of days use to graylist a status update . (default 45)
  -f string
    	Write output to file.
  -k string
    	Specify apiKey. (default "")
  -l	Should we do field lookups. (default true)
  -o string
    	Name of Portfoleon organization to use.
  -serve
    	Use if we should run a webserver.
  -status string
    	The status used for graylisting (requires writeable token).
  -t string
    	The name of template to use.
  -tJson string
    	The name of jsonfile to test the template with.
  -u string
    	Specify baseuUrl towards protfoleon (default "https://portfoleon.herokuapp.com/api/v1")
  -v string
    	Name of Portfoleon view to dump.
  -w string
    	Name of Portfoleon workspace to use.
```

Example
```
# Extract all data from Projects using the first organization and workspace and dump it in the file output 
protofoleon -k "PORTFOLEON_KEY" -v "Projects" -c 2 -f "output" 
# Do same data extraction, but now run out put first through the template before wirting it to output
protofoleon -k "PORTFOLEON_KEY" -v "Projects" -c 2 -f "output" -t "template.tpl" 
# Changes all workitem status to Gray within Projects if there was no status report for the last 45 days
protofoleon -k "PORTFOLEON_WRITE_KEY" -v "Projects" -days 45 -status "Gray"
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

## Building
We have added a build script `build.sh` which will create mac, linux and windows executable within the `dist` directory.

## Template engine
It is possible to run the result of the view directly through a GO html/template using the '-t'. An example of a template can be found in `example.tpl`
We have extended te template engine with the following functions
* now=current date
* slice=array creator
* strip=remove all html for string