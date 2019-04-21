FROM golang:1.12
MAINTAINER "Hamid Emamian <emami.he@gmail.com>"

COPY . /go/src/kubernetes-pod-waiter/
RUN cd /go/src/kubernetes-pod-waiter/ ; \
    go build -o /usr/local/waiter/sbin/waiter

WORKDIR /app

ENV PATH=$PATH:/usr/local/waiter/sbin/

CMD ["waiter"]
