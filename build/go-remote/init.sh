export GOBIN=/usr/local/go/bin/
export PATH=$PATH:$GOBIN
export PATH=$PATH:$HOME/go/bin
echo 'export GOBIN=/usr/local/go/bin/' >> /root/.bashrc
echo 'export PATH=$PATH:$GOBIN' >> .bashrc
echo 'export PATH=$PATH:$HOME/go/bin' >> /root/.bashrc
go get github.com/go-delve/delve/cmd/dlv
mkdir -p go/src/github.com/acmenezes/
git clone https://github.com/acmenezes/podconfig-operator.git  go/src/github.com/acmenezes/podconfig-operator
cd go/src/github.com/acmenezes/podconfig-operator
dlv debug --listen=:2345 --headless=true --log --api-version=2
