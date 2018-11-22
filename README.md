### Healthcheck

This project allows you to monitor health of serivces based on:
- availabilty of endpoint
- existence of key in redis
- file beeing touched in last minute
- any bash script
  
One can access healthcheck via HTTP server.
One can use this instead of using Docker-based healthcheck or to monitor any service on the host.

## Config description:

Config is a json with a list of entries. Every entry looks like this:
```
{
  "type": "redis_based" | "api_call_based" | "file_based" | "shell_based",
  "name": "whatever you want, some unique ID",
  "failedThreshold": 1, // after how many failed checks service is considered to be unhealthy
  "passedThreshold": 1, // after how many passed checks service is considered to be healthy
  "Interval": 1,        // in second - how frequently should we ask service for health
  "target": "some_key"  // this means something different for each type, for redis_based - key that has to exist, for api_call_based - endpoint that has to return 200, for file_based - file path, for shell_based - shell command that has to return 0.
}
```