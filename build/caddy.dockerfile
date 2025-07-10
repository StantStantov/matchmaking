FROM caddy:alpine AS base

COPY ./build/conf/Caddyfile /etc/caddy/Caddyfile
