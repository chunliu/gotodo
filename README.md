# gotodo

Todo list web api implemented with Go

## Enable SSL

1. Create SSL key:

    ```openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout ~/tmp/nginx.key -out ~/tmp/nginx.crt -subj "/CN=gotodo/O=gotodo"```

2. To test the SSL locally with docker image:

    ```docker run --name gotodo -p 80:80 -p 443:443 -d -v ~/tmp:/etc/nginx/ssl chunliu/gotodoimg```
    
3. 