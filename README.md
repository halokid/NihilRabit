NihilRabit
============================
a project for monitoring apps log file and recover the app process when you need.


[ ] http apps list
[ ] monitoring app log file
[ ] parser log file 
[ ] recover app process
[ ] support more nodes with http & one config file 
 


### project structure
- master   the main process, send order to nodes, data show, http
- nodes    monitoring log files, get config setting from master with http, do things like recover app process
- config   monitoring apps config file


### http restful api for master host
```markdown

# get node log path
GET    http://masterIp:port/getNodeLog/appName

#response:
{
  "path" :  "/path/to/log" 
}

---------------------------------------------



```
