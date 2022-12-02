#!/bin/sh

CGO_ENABLED=0 go build .
docker run --name renv --rm -it -v $(pwd):/data $@