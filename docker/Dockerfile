FROM ubuntu:latest
MAINTAINER Resultra Contributors
EXPOSE 43400
RUN apt-get update && apt-get install -y iputils-ping
ADD dest.tar.gz /usr/local/resultra/
WORKDIR /usr/local/resultra
ENTRYPOINT ./bin/resultraServer --config /var/resultra/config/resultraConfig.json