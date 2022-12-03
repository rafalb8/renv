#!/bin/sh

CGO_ENABLED=0 go build .
# docker run --name renv --rm -it -v $(pwd):/data $@

docker build -t renv -f - . << EOF
FROM $1
ADD renv /bin/renv

WORKDIR /root
ADD test .renv
WORKDIR /root/.renv
EOF

docker run --name renv --rm -it renv