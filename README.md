# gotodo

**gotodo** is a simple SPA to demo how to use several different technologies together in a single project. 

The backend Web API was developed with Go which implemented the CRUD in REST. It used the in memory object and didn't have a persistent data store. The frontend was developed with React and Antd. It fetches data from the backend. 

The app can be deployed to a web server directly, containerized in a docker image, or deployed to a Kubernetes cluster. 

## Build

To build the server component, go to `server` folder and run `go build -o gotodo`. 

To build the front end component, go to `frontend` folder and run `npm run prepublish`.

## Enable SSL

1. Create SSL key:

    ```openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout ~/tmp/nginx.key -out ~/tmp/nginx.crt -subj "/CN=gotodo/O=gotodo"```

2. To test the SSL locally with docker image:

    ```docker run --name gotodo -p 80:80 -p 443:443 -d -v ~/tmp:/etc/nginx/ssl chunliu/gotodoimg```
    