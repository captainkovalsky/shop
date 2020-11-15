#!/usr/bin/env sh

set -e

echo "copy artifacts to the server"

rsync -avz -e "ssh" --progress workspace/* $SSH_USER@$SSH_HOST:~/ci
