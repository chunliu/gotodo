FROM nginx

LABEL maintainer="Chun Liu <https://github.com/chunliu>"

# Update nginx settings to use it as a reverse proxy
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Setup Go app
WORKDIR /gotodo
COPY server/pages/ ./pages/
COPY server/static/ ./static/
COPY server/gotodo .
COPY gotodo.sh .
RUN chmod +x ./gotodo \
    && chmod +x ./gotodo.sh \
    && chmod +r ./pages/index.html \
    && chmod +r ./static/gotodo.js \
    && chmod +r ./static/style.css

EXPOSE 80 443

CMD [ "/bin/bash", "/gotodo/gotodo.sh" ]
