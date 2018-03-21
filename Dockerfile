FROM nginx

LABEL maintainer="Chun Liu <https://github.com/chunliu>"

# Update nginx settings to use it as a reverse proxy
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Setup Go app
WORKDIR /gotodo
COPY server/pages/index.html ./pages/index.html
COPY server/static/gotodo.js ./static/gotodo.js
COPY server/static/style.css ./static/style.css
COPY server/gotodo .
COPY gotodo.sh .
RUN chmod +x ./gotodo \
    && chmod +x ./gotodo.sh \
    && chmod +r ./pages/index.html \
    && chmod +r ./static/gotodo.js \
    && chmod +r ./static/style.css

EXPOSE 80 443

CMD [ "/bin/bash", "/gotodo/gotodo.sh" ]
