FROM golang:latest

RUN apt-get update -yqq
RUN DEBIAN_FRONTEND=noninteractive apt-get -yqq install tzdata
RUN ln -fs /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN apt-get install -yqq git wget curl python3 python3-pip tar unzip zip
RUN apt-get clean
RUN pip3 install coscmd