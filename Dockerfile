FROM scratch

MAINTAINER Owain Lewis <owain.lewis@oracle.com>

COPY bin/frequency /frequency

CMD ["/frequency"]
