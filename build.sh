#!/bin/sh


bash fetch-deps.sh
go get
bash bundle-me.sh
make
