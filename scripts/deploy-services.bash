#!/usr/bin/env sh

#circleci@ubuntu-s-1vcpu-1gb-ams3-01:~/ci/scripts$ cat /lib/systemd/system/shoper.service
#[Unit]
#Description=shoper
#
#[Service]
#Type=simple
#Restart=always
#RestartSec=5s
#ExecStart=/home/circleci/shop/shop-linux-amd64 web
#WorkingDirectory=/home/circleci/shop
#
#[Install]
#WantedBy=multi-user.target

#[Unit]
#Description=shoper-admin
#
#[Service]
#Type=simple
#Restart=always
#RestartSec=5s
#ExecStart=/home/circleci/shop/shop-linux-amd64 admin --username admin --password admin --port 9991
#WorkingDirectory=/home/circleci/shop/
#Environment="GIN_MODE=release"
#[Install]
#WantedBy=multi-user.target

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
