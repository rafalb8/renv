#!/bin/sh

set -ex

CGO_ENABLED=0 go build -ldflags="-s -w -X main.Version=docker" ./cmd/renv

docker build -t renv -f - . << EOF
FROM $1
ADD renv /bin/renv
RUN if ! getent group wheel >/dev/null; then groupadd wheel; fi
RUN useradd -ms /bin/sh -G wheel user || adduser -Ds /bin/sh -G wheel user
RUN renv install sudo
RUN echo '%wheel ALL=(ALL:ALL) NOPASSWD: ALL' > /etc/sudoers.d/wheel
USER user

WORKDIR /home/user
ADD test .renv
WORKDIR /home/user/.renv
EOF

docker run --name renv --rm -it renv