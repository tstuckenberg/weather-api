# Weather API Challenge

## Prerequisites

### Install Go 1.10.1 from the website
[Golang](https://golang.org/dl/ "Golang")

### Create Go environment

mkdir bin # for go binaries

mkdir pkg # for go external libraries

mkdir src # for your own go code

export GOPATH=$PATH:[your_path]/

export GOBIN=$GOPATH/bin


### Install dependencies

```go get```

### Start api

```go install```

```cd [your_path]/bin/```

```./api```

## Get weather

```<domain>/v1/weather?city=sydney```

```json
{
  "weather": {
    "wind_speed": 17.7,
    "temperature_degrees": 20
  },
  "success": true
}
```
