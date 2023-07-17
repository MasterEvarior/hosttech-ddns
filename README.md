# Hosttech DDNS
This small docker container acts as a [dynamic DNS](https://en.wikipedia.org/wiki/Dynamic_DNS) implementation for [Hosttech](https://hosttech.ch).

## Build
To build the container simply clone the repository and then build the container with the docker file
```shell
docker build . --tag hosttech-ddns
```
## Run
To run the docker container, you have to give it a couple of environment variables.
```shell
docker run -d \
  -e ZONE=example.com \
  -e DOMAIN=sub \
  -e API_KEY=eyJO93mopdns0JA \
  --name hosttech-ddns \
  hosttech-ddns:latest
```
This will update the IP address for the domain `sub.example.com`. 

## Development, improvements and more
Pull requests, improvements and issues are always welcome.