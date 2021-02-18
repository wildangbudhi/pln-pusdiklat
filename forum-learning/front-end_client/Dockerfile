FROM python:3.8.7-slim

WORKDIR /app
COPY ./app .

RUN apt-get -y update
RUN apt-get -y install gcc
RUN pip install --upgrade pip
RUN pip install wheel
RUN pip install -r requirements.txt
RUN pip install uwsgi

CMD [ "uwsgi", "--py-autoreload", "1", "--ini", "uwsgi.ini" ]