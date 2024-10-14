# Hosttech DDNS
![example workflow](https://github.com/MasterEvarior/hosttech-ddns/actions/workflows/test.yaml/badge.svg) ![example workflow](https://github.com/MasterEvarior/hosttech-ddns/actions/workflows/release.yaml/badge.svg)

This small docker container acts as a [dynamic DNS](https://en.wikipedia.org/wiki/Dynamic_DNS) implementation for [Hosttech](https://hosttech.ch).

## Build
To build the container yourself, simply clone the repository and then build the container with the provided docker file. You can the run it as described in the section below.
```shell
docker build . --tag hosttech-ddns
```
## Run
To run the docker container, you have to give it a couple of environment variables.
```shell
docker run -d \
  -e ZONE=example.com \
  -e DOMAINS=sub1,sub2 \
  -e API_KEY=eyJO93mopdns0JA \
  --name hosttech-ddns \
  ghcr.io/masterevarior/hosttech-ddns:latest
```
This will update the IP address for the domains `sub1.example.com` and `sub2.example.com`. 

## Development, improvements and more
Pull requests, improvements and issues are always welcome.