FROM debian:testing-slim

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update && \
    apt-get install -y \
	init \
	procps \
	openssh-client \
	mtools \
	zstd \
	sudo \
	iproute2 \
	iputils-ping \
	tar \
	curl \
	vim \
	gawk \
	jq \
	conntrack \
	iptables 

RUN systemctl enable systemd-networkd
RUN echo "root:root" | chpasswd
RUN apt-get clean && rm -rf /var/cache/* && \
    journalctl --vacuum-size=1K && rm /etc/machine-id && rm /var/lib/dbus/machine-id
