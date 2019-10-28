FROM golang AS builder
ADD . /go/src/github.com/yale8848/cutycapt-docker/
RUN cd /go/src/github.com/yale8848/cutycapt-docker \
  && CGO_ENABLED=0 GOOS=linux go build ./main/app.go

FROM yale8848/cutycapt-docker:base_ubuntu_v2
COPY --from=builder /go/src/github.com/yale8848/cutycapt-docker/app /bin/app

COPY ENTRYPOINT.sh /ENTRYPOINT.sh

EXPOSE 9600
ENTRYPOINT ["sh","/ENTRYPOINT.sh"]