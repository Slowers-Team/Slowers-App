FROM docker.io/library/node:18-alpine AS build-stage

WORKDIR /build
COPY . .
RUN npm clean-install && npm run build

FROM docker.io/library/nginx:alpine-slim

RUN chmod g+rwx /var/cache/nginx /var/run /var/log/nginx /etc/nginx/conf.d
RUN chmod g+rw /etc/nginx/conf.d/default.conf
RUN mkdir -m 775 /etc/nginx/templates

COPY --from=build-stage /build/nginx.conf.template /etc/nginx/templates/default.conf.template
COPY --from=build-stage /build/dist/ /usr/share/nginx/html/

EXPOSE 8081
