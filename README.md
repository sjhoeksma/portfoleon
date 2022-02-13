# Portfoleon
This is a smal tool to extract view data from portfoleon.com in json format so it can be used in other tools

## Steps
* Login using the APIKEY and convert it into Bearer token
* Select the organization by name, if no nama is specified the first organization available is used
* Select the workspace by name, if no name is specified the first workspace available is used
* Select the view by name or alldata if not set and convert data in to a single JSON file. In view data you can specify the number (default 0) of last status reports should be added

