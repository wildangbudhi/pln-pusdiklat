FROM tiangolo/uwsgi-nginx-flask:python3.6-alpine3.7
RUN apk --update add bash nano
WORKDIR /app
COPY . .
RUN pip install -r requirements.txt