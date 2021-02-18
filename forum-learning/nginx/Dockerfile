FROM nginx

WORKDIR /sslsetup

COPY ./config/ /etc/nginx/conf.d/
# ADD ./snippets/self-signed.conf /etc/nginx/snippets/
# ADD ./snippets/ssl-params.conf /etc/nginx/snippets/
# COPY ./certificate .

# RUN bash generate_cert.sh
# RUN bash generate_dhparam.sh