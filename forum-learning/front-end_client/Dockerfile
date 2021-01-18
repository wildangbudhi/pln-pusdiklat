FROM tiangolo/uwsgi-nginx-flask:python3.6-alpine3.7
RUN apk --update add bash nano
WORKDIR /app
COPY . .

ENV STATIC_PATH /app/static

RUN pip install --upgrade pip
RUN pip install -r requirements.txt