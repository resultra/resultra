FROM ubuntu:latest
MAINTAINER Resultra Contributors
EXPOSE 43400


RUN apt-get update


# By default, the email doesn't include ping support. This is nice
# to have for development and test. For example, the postgres
# server on the same Docker network can be pinged.
RUN apt-get install -y iputils-ping

# Without the ca-certificates package, the docker image 
# will be unable to send transactional emails, and generate
# errors like the following: smtp error: x509: certificate signed by unknown authority
RUN apt-get install -y ca-certificates

ADD dest.tar.gz /usr/local/resultra/

WORKDIR /usr/local/resultra
ENTRYPOINT ./bin/resultraServer --config /var/resultra/config/resultraConfig.json