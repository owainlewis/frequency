FROM scratch

MAINTAINER Owain Lewis <owain.lewis@oracle.com>

COPY bin/kcd /kcd

CMD ["/kcd"]
