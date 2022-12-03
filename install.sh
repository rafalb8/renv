#!/bin/sh
URL=$(curl -s https://api.github.com/repos/rafalb8/renv/releases/latest \
    | grep "browser_download_url.*amd64.tar.gz" \
    | cut -d : -f 2,3 \
    | tr -d \")

mkdir -p /tmp/renv
curl -sfL ${URL} | tar xvz -C /tmp/renv
sudo cp /tmp/renv/renv /bin/renv