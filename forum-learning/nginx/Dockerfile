FROM nginx

WORKDIR /sslsetup

COPY ./config/ /etc/nginx/conf.d/
# ADD ./config/self-signed.conf /etc/nginx/snippets/
# ADD ./config/ssl-params.conf /etc/nginx/snippets/
# COPY ./certificate .

# RUN bash generate_cert.sh
# RUN bash generate_dhparam.sh