#!/usr/bin/env sh

set -e

echo "copy artifacts to the server"

rsync -avz -e "ssh" --progress workspace/* $SSH_USER@$SSH_HOST:~/ci

echo "Deploy Services"

sshpass -e ssh -t $SSH_USER@$SSH_HOST './ci/scripts/deploy-services.bash'
