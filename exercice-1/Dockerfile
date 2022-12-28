FROM golang:1.19-alpine as backend
RUN mkdir /app
COPY go.mod /app
COPY go.sum /app
WORKDIR /app
RUN go mod download
COPY . /app
RUN CGO_ENABLED=0 go build -o activity2 .


FROM alpine as runner
RUN mkdir -p /app
WORKDIR /app
COPY --from=backend /app/activity2 ./activity2
RUN chmod a+x activity2
ENTRYPOINT /app/activity2