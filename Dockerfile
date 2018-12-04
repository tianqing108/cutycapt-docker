FROM golang AS builder
ADD . /go/src/github.com/yale8848/cutycapt-docker/
RUN cd /go/src/github.com/yale8848/cutycapt-docker \
  && CGO_ENABLED=0 GOOS=linux go build ./main/app.go

FROM centos
COPY --from=builder /go/src/github.com/yale8848/cutycapt-docker/app /bin/app
RUN yum install -y epel-release \
  && yum install -y Xvfb \
  && yum install -y xorg-x11-fonts* \
  && yum install -y google-noto-sans-simplified-chinese-fonts.noarch \
  && yum install -y mesa-dri-drivers \
  && yum install -y CutyCapt \
  && rm -rf /var/cache/yum \
  && dbus-uuidgen > /var/lib/dbus/machine-id

EXPOSE 9600
ENTRYPOINT ["/bin/app"]