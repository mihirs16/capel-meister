#!/usr/bin/env bash

docker compose build --no-cache
docker compose up --detach
docker compose nginx certbot --non-interactive --nginx -d projects.mihirsingh.dev -m mihirs16@gmail.com --agree-tos
docker compose nginx certbot renew --dry-run

