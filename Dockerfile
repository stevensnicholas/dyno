FROM golang:1.17.7 as builder
WORKDIR /app
COPY cmd cmd
COPY internal internal
COPY go.mod go.mod
COPY go.sum go.sum
COPY Makefile Makefile
RUN make build

FROM scratch
WORKDIR /
COPY --from=builder /app/bin/main /main
EXPOSE 8080
HEALTHCHECK --interval=1s --timeout=1s --start-period=2s --retries=3 CMD [ "/main", "-healthcheck" ]
CMD [ "/main", "-http" ]
