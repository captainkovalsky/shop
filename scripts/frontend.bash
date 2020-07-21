#!/usr/bin/env bash

set -e

cd $1

npm ci
npm run build:aot
cd -