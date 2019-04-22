FROM golang:1.12
MAINTAINER "Hamid Emamian <emami.he@gmail.com>"

COPY . /go/src/kubernetes-pod-dependency-handler
RUN cd /go/src/kubernetes-pod-dependency-handler/ ; \
    go build -o /go/bin/waiter

WORKDIR /go/bin

CMD ["waiter", "-s"]
