FROM golang

MAINTAINER Daniel T. Morgan <daniel@morgan.cymru>

RUN apt-get update && apt-get install -y --no-install-recommends \
		git \
		mercurial \
	&& rm -rf /var/lib/apt/lists/*

RUN go get github.com/jteeuwen/go-bindata/...
RUN go get github.com/elazarl/go-bindata-assetfs/...
RUN go get github.com/gorilla/mux
RUN go get github.com/braintree/manners
RUN go get -u -v github.com/spf13/hugo


ADD . /go/src/plod.tv

RUN /go/bin/hugo gen autocomplete

RUN cd /go/src/plod.tv; /go/bin/hugo

RUN cd /go/src/plod.tv; /go/bin/go-bindata -o /plod.tv.public.go/plod.tv.go -ignore=\\.gitkeep public/...

RUN go install plod.tv/plod.tv.public.go

ENTRYPOINT /go/bin/plod.tv.public.go

EXPOSE 8080