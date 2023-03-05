#!/bin/sh

CGO_ENABLED=0 go build .
# docker run --name renv --rm -it -v $(pwd):/data $@

docker build -t renv -f - . << EOF
FROM $1
ADD renv /bin/renv

RUN useradd -ms /bin/sh -G wheel user || adduser -Ds /bin/sh -G wheel user
RUN renv install sudo
RUN echo '%wheel ALL=(ALL:ALL) NOPASSWD: ALL' > /etc/sudoers.d/wheel
USER user

WORKDIR /home/user
ADD test .renv
WORKDIR /home/user/.renv
EOF

docker run --name renv --rm -it renv