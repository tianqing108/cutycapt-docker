FROM golang AS builder
ADD . /go/src/github.com/yale8848/cutycapt-docker/
RUN cd /go/src/github.com/yale8848/cutycapt-docker \
  && CGO_ENABLED=0 GOOS=linux go build ./main/app.go

FROM centos
COPY --from=builder /go/src/github.com/yale8848/cutycapt-docker/app /bin/app

RUN yum install -y epel-release \
  && yum install -y Xvfb \
  && yum install -y xorg-x11-fonts* \
  && yum install -y mesa-dri-drivers \
  && yum install -y qtwebkit-devel \
  && yum install -y qt-devel \
  && yum install -y CutyCapt \
  && dbus-uuidgen > /var/lib/dbus/machine-id

COPY fonts /usr/share/fonts/win
RUN  chmod 644 /usr/share/fonts/win/* && mkfontscale && mkfontdir && fc-cache -fv

EXPOSE 9600
ENTRYPOINT ["/bin/app"]