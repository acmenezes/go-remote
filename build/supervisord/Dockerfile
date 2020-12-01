FROM centos:7

USER root:root

RUN yum install -y epel-release; yum -y update; yum install -y python-pip; \
yum -y install iproute iproute.doc openssh-server openssh-clients vim wget; \
yum -y groupinstall 'Development Tools'; yum -y remove git; \
yum -y install https://packages.endpoint.com/rhel/7/os/x86_64/endpoint-repo-1.7-1.x86_64.rpm;  \
yum -y install git; wget https://golang.org/dl/go1.15.5.linux-amd64.tar.gz; \
tar -C /usr/local -xzf go1.15.5.linux-amd64.tar.gz; \
/usr/local/go/bin/go get github.com/go-delve/delve/cmd/dlv; pip install supervisor; \
mkdir -p /root/go/src/github.com/project; echo "root:go1234" | chpasswd; /usr/bin/ssh-keygen -A;

COPY supervisord.conf /etc/supervisor/supervisord.conf
COPY sshd_config /etc/ssh/sshd_config

EXPOSE 2222
EXPOSE 2345

CMD ["/usr/bin/supervisord"]


