#!/usr/bin/env sh

set -e

echo "deploy application"

rsync -avz -e "ssh" --progress workspace/* $SSH_USER@$SSH_HOST:~/ci
ssh $SSH_USER@$SSH_HOST .~/ci/scripts/deploy-services.bash
