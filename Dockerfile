FROM golang as build

WORKDIR /app
ADD . ./
ENV CGO_ENABLED=0
ENV GOPROXY=https://proxy.golang.org
RUN go build -o credit_scoreing_v2 main.go

FROM alpine

RUN apk add tzdata ca-certificates

WORKDIR /app
COPY --from=build /app/credit_scoreing_v2 ./

ENTRYPOINT ["/app/credit_scoreing_v2"]


