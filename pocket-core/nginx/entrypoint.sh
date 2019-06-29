#!bin/sh
# Set utils

set -o errexit
set -o pipefail
set -o nounset

aws s3 sync $NGINX_CONFIG_BUCKET_URL /etc/nginx/

cmd="$@"
