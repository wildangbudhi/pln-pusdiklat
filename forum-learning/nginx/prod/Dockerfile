FROM ubuntu:18.04

RUN apt update -y \
    && apt install nginx curl vim -y \
    && apt-get install software-properties-common -y \
    && add-apt-repository ppa:certbot/certbot -y \
    && apt-get update -y \
    && apt-get install python-certbot-nginx -y \
    && apt-get clean

RUN apt-get update -y \
    && apt-get install -y --no-install-recommends gettext \
    && rm -rf /var/lib/apt/lists/*

COPY ./templates /etc/nginx/templates
COPY ./init.sh ./root/
RUN ["chmod", "+x", "/root/init.sh"]

STOPSIGNAL SIGTERM

ENTRYPOINT ["nginx", "-g", "daemon off;"]