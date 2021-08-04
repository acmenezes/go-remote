FROM centos:8

USER root:root

RUN yum install -y epel-release; yum -y update; yum install -y supervisor; \
yum -y install openssh-server openssh-clients vim wget zsh; \
yum -y groupinstall 'Development Tools'; wget https://golang.org/dl/go1.16.5.linux-amd64.tar.gz; \
tar -C /usr/local -xzf go1.16.5.linux-amd64.tar.gz; \
/usr/local/go/bin/go get github.com/go-delve/delve/cmd/dlv; \
/usr/local/go/bin/go get github.com/ramya-rao-a/go-outline; \
/usr/local/go/bin/go get golang.org/x/tools/cmd/goimports; \
/usr/local/go/bin/go get github.com/rogpeppe/godef; \
wget https://github.com/operator-framework/operator-sdk/releases/download/v1.8.0/operator-sdk_linux_amd64; \
chmod +x operator-sdk_linux_amd64; \
mv operator-sdk_linux_amd64 /usr/local/bin/operator-sdk; \
wget https://mirror.openshift.com/pub/openshift-v4/clients/ocp/4.7.12/openshift-client-linux-4.7.12.tar.gz; \
tar -C /usr/local/bin/ -xzf openshift-client-linux-4.7.12.tar.gz; \
mkdir -p /root/go/src/github.com/project; echo "root:go1234" | chpasswd; /usr/bin/ssh-keygen -A; \
yum install -y util-linux-user; chsh -s /bin/zsh; wget https://github.com/robbyrussell/oh-my-zsh/raw/master/tools/install.sh -O - | zsh;

COPY supervisord.conf /etc/supervisord.conf
COPY sshd_config /etc/ssh/sshd_config
COPY .zshrc /root/.zshrc

EXPOSE 2222

CMD ["/usr/bin/supervisord"]


