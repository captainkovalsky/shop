#!/usr/bin/env bash

set -e

sudo service shoper stop
sudo service shoper-admin stop

rm -rf  /home/circleci/shop
mkdir -p /home/circleci/shop/static
tar -xvf ./ci/back.tar -C /home/circleci/shop
cp -v ./ci/cv.json /home/circleci/shop

tar -xvf ./ci/front.tar -C /home/circleci/shop/static

touch /home/circleci/shop/.env
ls -la /home/circleci/shop
ls -la /home/circleci/shop/static

sudo service shoper start
sudo service shoper-admin start
