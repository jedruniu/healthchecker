### Healthcheck

This project allows you to monitor health of serivces based on:
- availabilty of endpoint
- existence of key in redis
- file beeing touched in last minute
- any bash script
  
One can access healthcheck via HTTP server.

One can use this instead of using Docker-based healthcheck or to monitor any service on the host.