# builder image
FROM golang:1.17.3-alpine as builder
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN GOOS=linux go build -a -o api .

# generate clean, final image for end users
FROM scratch
COPY --from=builder /app /app

WORKDIR /app
EXPOSE 5000
ENTRYPOINT [ "./api" ]
