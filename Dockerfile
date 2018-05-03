FROM golang:1.9.5-alpine3.7

WORKDIR /go/src/app

RUN apk --update add  --virtual build-deps \
			build-base \
		&& apk add \
			tzdata \
			make \
			git \
		&& rm -rf /var/cache/apk/* \ 
		&& cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime

COPY . $APP

RUN go get -u github.com/golang/dep/cmd/dep
RUN go get -u github.com/tockins/realize
RUN dep ensure

CMD ["go", "run", "server.go"]
