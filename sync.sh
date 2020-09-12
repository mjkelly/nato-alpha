#!/bin/bash
set -u
set -e
cfg=$HOME/.aws/config
profile=admin
bucket=s3://nato-alpha.michaelkelly.org
dir=www
# cf_id=EJKWL55KLXE9D
scripts=$(dirname $0)

which aws || (echo "'aws' command not found. Aborting."; exit 2)
[ -f $cfg ] || (echo "Config file $cfg does not exist. Aborting."; exit 2)

echo "Deploying..."
echo "Synchronizing directory $PWD/$dir"
aws --profile="$profile" \
  s3 sync "$dir" "$bucket" \
  --acl=public-read --cache-control=max-age=3600
# echo "Invalidating CloudFront..."
# aws --profile=admin \
#   cloudfront create-invalidation \
#   --distribution-id "${cf_id}" \
#   --paths "/*"
