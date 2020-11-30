FROM quay.io/acmenezes/c7-systemd

USER root:root

ENV GOBIN=/usr/local/go/bin/
ENV PATH=$PATH:$GOBIN
ENV PATH=$PATH:$HOME/go/bin

RUN yum -y install iproute iproute.doc openssh-server openssh-clients vim wget
RUN yum -y groupinstall 'Development Tools'
RUN yum -y remove git
RUN yum -y install https://packages.endpoint.com/rhel/7/os/x86_64/endpoint-repo-1.7-1.x86_64.rpm 
RUN yum -y install git
RUN wget https://golang.org/dl/go1.15.5.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.15.5.linux-amd64.tar.gz
RUN go get github.com/go-delve/delve/cmd/dlv
RUN mkdir -p go/src/github.com/acmenezes/
RUN echo "root:go1234" | chpasswd
RUN yum clean all; systemctl enable sshd.service

COPY ./sshd_config /etc/ssh/sshd_config
COPY ./kubeconfig /root/.kube/config

CMD ["/usr/sbin/init"]