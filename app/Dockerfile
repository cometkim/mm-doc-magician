FROM golang:alpine

RUN set -ex; \
    apk add --no-cache curl git && \ 
    curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh && \
    dep version

WORKDIR /go/src/app
COPY . .

RUN set -ex; \
    dep ensure -v && \
    go install -v ./... 

CMD ["app"]
