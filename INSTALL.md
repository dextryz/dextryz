# Install

Setup guide on AlmaLinux in a DigitalOcean droplet.

- Install snap

```shell
dnf install epel-release
dnf update -y
dnf upgrade -y
dnf install snapd -y
systemctl enable --now snapd.socket
ln -s /var/lib/snapd/snap /snap
reboot
```

- Setup Nginx

```shell
dnf install neovim nginx -y

systemctl start nginx
systemctl enable nginx

touch /etc/nginx/conf.d/dextryx.com.conf
nvim /etc/nginx/conf.d/dextryx.com.conf

nginx -s reload
```

- Install certbot

```shell
snap install --classic certbot
ln -s /snap/bin/certbot /usr/bin/certbot
certbot --nginx
```

- Install Go

```shell
dnf install wget -y
cd /usr/local/src/
wget https://go.dev/dl/go1.21.6.linux-amd64.tar.gz
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.21.6.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
go version
```

- Disable SELinux

```shell
sestatus 
setenforce 0
```

- Spinup Server

```shell
dnf install git -y
git clone https://github.com/dextryz/dextryz.git
cd dextryz
go run .
```

## nvim /etc/nginx/conf.d/dextryx.com.conf

```
server {

   server_name dextryz.com; 

   location / {
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header Host $http_host;
        proxy_pass http://127.0.0.1:8080;
    }
}
```
