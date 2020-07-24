#!/usr/bin/env bash

set -e

sudo service shoper stop
sudo service shoper-admin stop

rm -rf  /home/circleci/shop
mkdir -p /home/circleci/shop/static
tar -xvf back.tar -C /home/circleci/shop
tar -xvf front.tar -C /home/circleci/shop/static

ls -la /home/circleci/shop
ls -la /home/circleci/shop/static

sudo service shoper start
sudo service shoper-admin start
