FROM centos
RUN yum install -y epel-release \
  && yum install -y Xvfb \
  && yum install -y mesa-dri-drivers \
  && yum install -y qtwebkit-devel \
  && yum install -y qt-devel \
  && yum install fonts-chinese \
  && yum install -y CutyCapt \
  && rm -rf /var/cache/yum \
  && dbus-uuidgen > /var/lib/dbus/machine-id