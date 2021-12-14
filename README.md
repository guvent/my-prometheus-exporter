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

go mod download

go build -o my_exporter

mkdir /var/lib/my_exporter

mv my_exporter /var/lib/my_exporter

chmod +x /var/lib/my_exporter/my_exporter

touch /var/lib/my_exporter/.env

vi /etc/systemd/system/myexporter.service

```
[Unit]
Description=My channel exporter
After=network.target
StartLimitIntervalSec=0

[Service]
Type=simple
Restart=always
RestartSec=1
WorkingDirectory=/var/lib/my_exporter
EnvironmentFile=/var/lib/my_exporter/.env
ExecStart=/var/lib/my_exporter/my_exporter

[Install]
WantedBy=multi-user.target
```

systemctl daemon-reload
systemctl enable --now myexporter.service
systemctl status myexporter.service

open web browser and go to http://<SERVER_ADDRESS>:9888/

add end of line for scrape_configs: on prometheus.yml 

```
    - job_name: "my_exporter"
      scrape_interval: 10s
      scrape_timeout: 10s
      metrics_path: "/metrics"
    
      static_configs:
        - targets: ["localhost:9888"]
```

systemctl restart prometheus.service

Good luck...