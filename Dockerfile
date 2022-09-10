# load base image for nginx with alpine
FROM nginx:stable-alpine

# install certbot
RUN apk add --update python3 py3-pip; \ 
    apk add certbot; \ 
    pip install certbot-nginx;

# copy default.conf
COPY ./nginx/conf.d/ /etc/nginx/conf.d/

# generate ssl certificate
RUN certbot --non-interactive --nginx -d projects.mihirsingh.dev -m mihirs16@gmail.com --agree-tos; \
    certbot renew --dry-run;

