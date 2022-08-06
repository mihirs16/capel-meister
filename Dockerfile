# load base image for nginx with alpine
FROM nginx:stable-alpine

# copy default.conf
COPY ./nginx/conf.d/ /etc/nginx/conf.d/

# copy certbot setup script
COPY . .
RUN chmod +x ./nginx/cert-setup.sh

