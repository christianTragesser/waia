Who Am I App
=========

An application which shows the hostname, IP address and version of the application you're interacting with.

### Dependencies
- Docker

### Run

```
docker run -p 8080:8080 registry-gitlab.FQDN_HERE/root/waia/waia
```

### Customize 
Change the background color to your liking by setting the environment variable `BG_COLOR`.

```
docker run -e BG_COLOR="blue" -p 8080:8080 registry-INSTANCE_NAME_HERE.FQDN_HERE/root/waia/waia
```

### License

Apache 2.0
