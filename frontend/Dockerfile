FROM nginx

# Update nginx settings to serve static files
COPY nginx.conf /etc/nginx/conf.d/default.conf
# Copy static files
WORKDIR /gotodofe
COPY dist/index.html .
COPY dist/gotodo.js .
COPY dist/style.css .

RUN chmod +r ./index.html \
  && chmod +r ./gotodo.js \
  && chmod +r ./style.css

CMD ["nginx", "-g", "daemon off;"]