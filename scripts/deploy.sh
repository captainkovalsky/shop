#!/usr/bin/env sh

set -e

echo "deploy application"
ssh $SSH_USER@$SSH_HOST rm -rf ci
scp workspace/back.tar $SSH_USER@$SSH_HOST:/ci
scp workspace/front.tar $SSH_USER@$SSH_HOST:/ci

