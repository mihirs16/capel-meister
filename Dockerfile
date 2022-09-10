# load base image for nginx with alpine
FROM nginx:stable-alpine

# install certbot and nginx plugin
RUN apk add --update python3 py3-pip; \ 
    apk add certbot; \ 
    pip install certbot-nginx;

# copy default.conf
COPY ./nginx/conf.d/ /etc/nginx/conf.d/

