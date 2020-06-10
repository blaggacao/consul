#!/bin/bash
set -xe

# Install netlify-cli
sudo npm install netlify-cli -g

# Deploy site to netlify
# Assumes NETLIFY_SITE_ID and NETLIFY_AUTH_TOKEN env variables are set
output=$(netlify deploy --dir=./website/build --json)

# Grab deploy URL
url=$(echo "$output" | jq --raw-output '.deploy_url')

# Checks broken links
wget \
  --delete-after \
  --level inf \
  --no-verbose \
  --recursive \
  --no-directories \
  --no-host-directories \
  --page-requisites \
  --spider \
  $url
