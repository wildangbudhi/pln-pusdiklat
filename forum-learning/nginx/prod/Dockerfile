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
RUN envsubst < /etc/nginx/templates/prod.conf.template > /etc/nginx/sites-available/default.conf
RUN rm etc/nginx/sites-enabled/default.conf
RUN ln -s /etc/nginx/sites-available/default.conf /etc/nginx/sites-enabled/

STOPSIGNAL SIGTERM

CMD ["nginx", "-g", "daemon off;"]