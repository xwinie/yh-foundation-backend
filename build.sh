#!/bin/bash

CGO_ENABLED=0 gox -osarch="linux/amd64" -ldflags "-s -w " -gcflags -m -output dist/yh-foundation-backend
# docker run --rm -it -v $GOPATH:/go golang:latest bash -c 'cd $GOPATH/src/yh-foundation-backend && CGO_ENABLED=0 go build -ldflags "-s -w " -gcflags -m -o dist/yh-foundation-backend'
# cp -rf conf dist
upx -f -9 dist/yh-foundation-backend && docker build -t api .

# delivery ---

# docker run  -v $GOPATH/src/yh-foundation-backend/static:/root/app/static -p 8080:8080  api
# docker exec -it d460ba92ea0a  /bin/bash
# cd liquibase && sh migration.sh