FROM golang:onbuild
MAINTAINER    Youngwan <tiny657@gmail.com>
RUN go get github.com/tiny657/url-shortener
RUN go get -u github.com/go-redis/redis
RUN go install github.com/tiny657/url-shortener

# install redis
RUN apt-get update
RUN apt-get -y install redis-server

ADD script/initializeDocker.sh ./
CMD ./initializeDocker.sh

EXPOSE 8080
