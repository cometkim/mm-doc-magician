FROM golang:alpine

RUN set -ex; \
    apk add --no-cache curl git && \ 
    curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh && \
    dep version

WORKDIR /go/src/app

COPY Gopkg.lock .
COPY Gopkg.toml .
RUN dep ensure -v -vendor-only

COPY . .
RUN go install -v ./... 

CMD ["app"]
