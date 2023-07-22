FROM golang:latest

ENV GOPATH=/go
ENV PATH=$GOPATH/bin:$PATH


RUN mkdir -p $GOPATH/src/nginx

WORKDIR $GOPATH/src/nginx

COPY . .

RUN go build -o $GOPATH/bin/nginx


CMD ["nginx"]