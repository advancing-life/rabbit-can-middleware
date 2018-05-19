FROM golang:1.9.5-alpine3.7

WORKDIR /go/src/app

COPY . $APP

RUN apk --update add  --virtual build-deps \
			build-base \
		&& apk add \
			tzdata \
			make \
			git \
		&& apk add --no-cache docker \
		&& rm -rf /var/cache/apk/* \ 
		&& cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime


RUN go get -u -v github.com/golang/dep/cmd/dep 
RUN go get -u -v github.com/tockins/realize

CMD ["go", "run", "server.go"]
