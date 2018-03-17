FROM nginx

LABEL maintainer="Chun Liu <https://github.com/chunliu>"

# Update nginx settings to use it as a reverse proxy
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Setup Go app
WORKDIR /gotodo
COPY pages .
COPY gotodo .
COPY gotodo.sh .
RUN chmod +x ./gotodo \
    && chmod +x ./gotodo.sh 

EXPOSE 80 443

CMD [ "/bin/bash", "/gotodo/gotodo.sh" ]
