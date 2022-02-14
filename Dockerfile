FROM golang:1.16-alpine as builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY controllers/ ./controllers
COPY core/ ./core
COPY dto/ ./dto
COPY flow/ ./flow
COPY middleware ./middleware
COPY service/ ./service
COPY main.go ./
#COPY .env ./


RUN export CGO_ENABLED=0 && go build -o /mazu-admin-api


##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=builder /mazu-admin-api /mazu-admin-api
#git COPY --from=builder /app/.env/ .
COPY --from=builder /app/flow/ .

EXPOSE 80

#USER nonroot:nonroot

#ENTRYPOINT ["/mazu-admin-api"]

CMD ["./main"]