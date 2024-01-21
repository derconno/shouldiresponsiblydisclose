FROM --platform=$BUILDPLATFORM golang:latest AS builder

RUN mkdir "/opt/shouldiresponsiblydisclose"
WORKDIR /opt/shouldiresponsiblydisclose
COPY go.mod main.go /opt/shouldiresponsiblydisclose/

ARG TARGETOS TARGETARCH
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -ldflags "-s -w" -o shouldiresponsiblydisclose

FROM scratch

COPY --from=builder /opt/shouldiresponsiblydisclose/shouldiresponsiblydisclose /shouldiresponsiblydisclose
COPY static/* /static/
COPY templates/* /templates/
COPY data.json /data.json

EXPOSE 8080

ENTRYPOINT ["/shouldiresponsiblydisclose"]
