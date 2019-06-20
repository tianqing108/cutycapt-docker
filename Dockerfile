FROM golang:1.12-alpine AS builder
ADD . /go/src/github.com/yale8848/cutycapt-docker/
RUN cd /go/src/github.com/yale8848/cutycapt-docker \
  && CGO_ENABLED=0 GOOS=linux go build ./main/app.go

FROM  ubuntu14.04

COPY --from=builder /go/src/github.com/yale8848/cutycapt-docker/app /bin/app

RUN  apt-get update \
     && apt-get install Xvfb  \
     && apt-get install -y CutyCapt

COPY fonts /usr/share/fonts/win

RUN  chmod 644 /usr/share/fonts/win/* && mkfontscale && mkfontdir && fc-cache -fv

EXPOSE 9600

ENTRYPOINT ["/bin/app"]