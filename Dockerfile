FROM alpine:3.1
ENV GOPATH /go
RUN apk -U add go git mercurial
COPY . /go/src/github.com/progrium/docker-plugins-stub
WORKDIR /go/src/github.com/progrium/docker-plugins-stub
RUN go get
CMD go get \
	&& go build -o /bin/stub \
	&& exec /bin/stub
