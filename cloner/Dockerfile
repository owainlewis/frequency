FROM alpine:3.7

RUN apk update && apk add --no-cache \
    ca-certificates \
    git \
    openssh-client \
    && update-ca-certificates

COPY scripts /

CMD /clone.sh
