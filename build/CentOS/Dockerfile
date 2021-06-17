FROM centos:7

USER root:root

RUN yum install -y epel-release; yum -y update; yum install -y python-pip; \
yum -y install iproute iproute.doc openssh-server openssh-clients vim wget zsh; \
yum -y groupinstall 'Development Tools'; yum -y remove git; \
yum -y install https://packages.endpoint.com/rhel/7/os/x86_64/endpoint-repo-1.7-1.x86_64.rpm;  \
yum -y install git; wget https://golang.org/dl/go1.16.2.linux-amd64.tar.gz; \
tar -C /usr/local -xzf go1.16.2.linux-amd64.tar.gz; \
/usr/local/go/bin/go get github.com/go-delve/delve/cmd/dlv; pip install supervisor; \
/usr/local/go/bin/go get goimports; /usr/local/go/bin/go get godef; /usr/local/go/bin/go get go-outline; \
wget https://github.com/operator-framework/operator-sdk/releases/download/v1.5.0/operator-sdk_linux_amd64; \
chmod +x operator-sdk_linux_amd64; \
mv operator-sdk_linux_amd64 /usr/local/bin/operator-sdk; \
wget https://mirror.openshift.com/pub/openshift-v4/clients/ocp/4.7.0/openshift-client-linux-4.7.0.tar.gz; \
tar -C /usr/local/bin/ -xzf openshift-client-linux-4.7.0.tar.gz; \
mkdir -p /root/go/src/github.com/project; echo "root:go1234" | chpasswd; /usr/bin/ssh-keygen -A; \
chsh -s /bin/zsh; wget https://github.com/robbyrussell/oh-my-zsh/raw/master/tools/install.sh -O - | zsh;

COPY supervisord.conf /etc/supervisor/supervisord.conf
COPY sshd_config /etc/ssh/sshd_config
COPY kubeconfig /root/.kube/config
COPY .zshrc /root/.zshrc

EXPOSE 2222
EXPOSE 2345

CMD ["/usr/bin/supervisord"]


