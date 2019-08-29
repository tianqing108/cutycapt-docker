FROM centos

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

RUN yum -y install java-1.8.0-openjdk-devel.x86_64