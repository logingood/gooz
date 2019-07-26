# Build binary
FROM alpine:3.9 AS builder

ENV GOPATH /go
ENV PATH /go/bin/:$PATH

RUN apk add --no-cache git make musl-dev go &&\
    mkdir -p /go/src /go/bin /go/src/github.com/logingood/gooz &&\
    go get -u github.com/golang/dep/cmd/dep/...

COPY . /go/src/github.com/logingood/gooz
WORKDIR /go/src/github.com/logingood/gooz

RUN dep ensure -v &&\
    go build -o /bin/gooz

# Build image
FROM alpine:3.9

COPY --from=builder /bin/gooz /bin/gooz

ARG REVISION
LABEL REVISION=$REVISION

RUN echo "$REVISION" > /REVISION &&\
    addgroup -g 1000 -S app &&\
    adduser -u 1000 -S app -G app

USER 1000

CMD ["/bin/gooz"]
