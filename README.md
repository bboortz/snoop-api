# Snoop API

A simple API/webserver that is responding with the request as a body. It is useful for debugging.
The snoop API can be packed as a normal binary or using docker.



# Why?

1. For debugging, e.g. 
** debugging your internet connection. Answering which IP are you using?
** debugging your proxy. Answering which information is your browser sending?
2. Having a simple API that can be used from other applications



# Status

This is a very early version of snoop API. In future it is planned to extend it with
* a simple javascript UI, e.g. using [Vue.js](https://vuejs.org/)
* base64 decoding for base64 encoded header fields or parameters


# How to ... ?


## How to build

The normal build is building the docker container.

```
./build.sh
```

You can also set the variable `DEV` in case you want to just build the software and not package it using docker. 

```
DEV=1 ./build.sh
```


## How to run

The normal run is running the softwware in a docker container

```
./run.sh
```

You can also set the variable `DEV` in case you want to just run the software and run it using docker.

```
DEV=1 ./run.sh
```


## How to access

```
curl -k https://localhost:8443
```


## How to configure

Put a file `snoop.yaml` to the currect directory. The content can look like this:
For a configuration example, please have a look into the file `examples/snoop.yaml`



## How to test the performance 

I have tested using [siege](https://www.joedog.org/siege-home/)
You can run performances tests easily using 

```
siege -r 1 -c 100 --no-follow https://localhost:8443
```


## Hot to debug

You can inspect the running docker container using

```
docker log -f snoop
```


## How to contribute

fork -> create branch -> create pull request

