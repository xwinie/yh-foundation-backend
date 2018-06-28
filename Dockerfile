FROM mkorenkov/alpine
#FROM alpine:latest
#
#RUN apk add --update curl gnupg tzdata && \
#    export GLIBC_VERSION="2.23-r1" && \
#    curl -o glibc.apk -L "https://github.com/andyshinn/alpine-pkg-glibc/releases/download/${GLIBC_VERSION}/glibc-${GLIBC_VERSION}.apk" && \
#    apk add --allow-untrusted glibc.apk && \
#    curl -o glibc-bin.apk -L "https://github.com/andyshinn/alpine-pkg-glibc/releases/download/${GLIBC_VERSION}/glibc-bin-${GLIBC_VERSION}.apk" && \
#    apk add --allow-untrusted glibc-bin.apk && \
#    /usr/glibc-compat/sbin/ldconfig /lib /usr/glibc/usr/lib && \
#    echo 'hosts: files mdns4_minimal [NOTFOUND=return] dns mdns4' >> /etc/nsswitch.conf && \
#    export GOSU_VERSION="1.7" && \
#    curl -L -o /tmp/gosu "https://github.com/tianon/gosu/releases/download/$GOSU_VERSION/gosu-amd64" && \
#    curl -L -o /tmp/gosu.asc "https://github.com/tianon/gosu/releases/download/$GOSU_VERSION/gosu-amd64.asc" && \
#    export GNUPGHOME="$(mktemp -d)" && \
#    gpg --keyserver ha.pool.sks-keyservers.net --recv-keys B42F6819007F00F88E364FD4036A9C25BF357DD4 && \
#    gpg --batch --verify /tmp/gosu.asc /tmp/gosu && \
#    rm -r "$GNUPGHOME" /tmp/gosu.asc && \
#    mv /tmp/gosu /usr/local/bin/ && \
#    chmod +x /usr/local/bin/gosu && \
#    gosu nobody true && \
#    apk del curl gnupg && \
#    rm -f glibc.apk glibc-bin.apk && \
#    rm -rf /var/cache/apk/*
#RUN apk --update add openjdk7-jre

RUN mkdir -p /root/app

WORKDIR /root/app

COPY dist .
#COPY liquibase .
RUN chmod +x -R /root

ENTRYPOINT ("/root/app/yh-foundation-backend")
EXPOSE 8080
