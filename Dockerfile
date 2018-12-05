FROM centos
RUN yum install -y epel-release \
  && yum install -y Xvfb \
  && yum install -y xorg-x11-fonts* \
  && yum install -y mesa-dri-drivers \
  && yum install -y qtwebkit-devel \
  && yum install -y qt-devel \
  && yum install -y CutyCapt