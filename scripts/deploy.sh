#!/usr/bin/env sh

set -e

echo "deploy application"

#rsync -avz -e "ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null" --progress workspace/* $SSH_USER@$SSH_HOST:~/ci
rsync -avz -e "ssh" --progress workspace/* $SSH_USER@$SSH_HOST:~/ci