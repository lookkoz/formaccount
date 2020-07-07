FROM golang:1.11-alpine

# install git in order to use Go modules
# update first to ensure the latest version of git is installed
RUN set -ex; \
    apk update; \ 
    apk add --no-cache git

WORKDIR /go/src/formaccount

CMD ./waitforaccountapi.sh accountapi 8080 go clean -testcache . \
        && ACCOUNT_API_HOST="accountapi" ACCOUNT_API_PORT="8080" CGO_ENABLED=0 \
        go test --cover ./...