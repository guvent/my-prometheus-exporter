### Operating System

* Debian 10
* Go 1.15

Golang 1.15.8 Install;

```
wget https://golang.org/dl/go1.15.8.linux-amd64.tar.gz

rm -rf /usr/local/go && tar -C /usr/local -xzf go1.15.8.linux-amd64.tar.gz

mkdir $HOME/work

- open .bashrc and .profile files put these line on end of line;

export GO111MODULE="on"
export GOROOT=/usr/local/go
export GOPATH=$HOME/work
export GOMODCACHE=$GOPATH/pkg/mod
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin

- save and run this command;

source .bashrc

go version

```

go build -o main




Good luck...