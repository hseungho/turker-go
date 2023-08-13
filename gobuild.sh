#!/bin/bash

ROOT=$(pwd)

cd $ROOT/$1
go mod init turker_go/$1

cd $ROOT
go work use $1

cd $ROOT/$1
go build
