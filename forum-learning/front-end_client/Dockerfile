FROM ubuntu:18.04

RUN apt-get -y update
RUN DEBIAN_FRONTEND=noninteractive apt install -y --no-install-recommends software-properties-common
RUN add-apt-repository -y ppa:deadsnakes
RUN DEBIAN_FRONTEND=noninteractive apt install -y --no-install-recommends python3.8
RUN DEBIAN_FRONTEND=noninteractive apt install -y --no-install-recommends python3-pip
RUN apt-get -y install python3-setuptools
RUN apt-get -y install python3-dev

WORKDIR /app
COPY ./app .

RUN apt-get -y update
RUN apt-get -y install gcc

RUN pip3 install --upgrade pip
RUN pip3 install wheel
RUN pip3 install -r requirements.txt
RUN pip3 install uwsgi

CMD [ "uwsgi", "--py-autoreload", "1", "--ini", "uwsgi.ini" ]