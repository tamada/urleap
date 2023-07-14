FROM golang:1.19-bullseye as builder

ARG TOKEN

ADD . /go/urleap
WORKDIR /go/urleap
RUN export URLEAP_TOKEN=${TOKEN} && make clean && make && adduser --disabled-login --disabled-password nonroot

FROM scratch

COPY --from=builder /go/urleap/urleap /usr/bin/urleap
COPY --from=builder /etc/passwd /etc/passwd
USER nonroot

ENTRYPOINT [ "/usr/bin/urleap" ]