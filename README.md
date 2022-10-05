# telegram-random-quotes

## build
```bash
GOOS=linux GOARCH=amd64 go build /Volumes/Data/coding/golang/cronJobProject/cronjob.go
```
## upload to host:
```bash
scp cronjob root@163.44.206.132:/usr/local/src/telegram-bot/
```
## create service:
```bash
vi /etc/systemd/system/telegram_bot.service
```

then add code:
```bash
[Unit]
Description=Telegram-bot  service
After=multi-user.target

[Service]
User=root
Group=root
Type=simple
Restart=always
RestartSec=5s
ExecStart=/usr/local/src/telegram-bot/cronjob

[Install]
WantedBy=multi-user.target
```

## run
```bash
sudo systemctl start telegram_bot.service
sudo systemctl enable telegram_bot.service
sudo systemctl status telegram_bot.service
```

## create service
```bash
vi /etc/nginx/sites-available/telegram_bot
```
then add this code
```azure
server {
        listen 80;

        location /dev {
                proxy_pass http://127.0.0.1:8081/api/v1;
                proxy_set_header Host $host;
                proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        }
         location /production {
                proxy_pass http://127.0.0.1:8082/api/v1;
                proxy_set_header Host $host;
                proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        }
}
```