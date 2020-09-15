FROM buildpack-deps:bullseye-scm

ENV DEBIAN_FRONTEND noninteractive

RUN apt-get update && apt-get install -y --no-install-recommends \
  apache2-utils \
  arping \
  arp-scan \
  bash \
  zsh \
  bind9-host \
  bridge-utils \
  ctop \
  curl \
  dhcping \
  dnsutils \
  ethtool \
  file \
  fping \
  iftop \
  iperf \
  iproute2 \
  ipset \
  iptables \
  iptraf-ng \
  iputils-ping \
  ipvsadm \
  mtr \
  netcat-openbsd \
  netgen \
  net-tools \
  nftables \
  ngrep \
  nmap \
  openssl \
  scapy \
  socat \
  strace \
  tcpdump \
  tcptraceroute \
  telnet \
  tftp-hpa \
  traceroute \
  util-linux \
  vim \
  wget \
  && apt-get clean -y \
  && rm -rf /var/lib/apt/lists/*

CMD ["zsh"]
