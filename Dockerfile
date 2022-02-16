FROM golang:1.18beta2-alpine3.15
WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app
COPY ./resources /usr/local/bin/resources

CMD ["app", "-active-profiles=default,dev", "-resource-root=/usr/local/bin/resources"]
# Run with this profile to enable user endpoints
#CMD ["app", "-active-profiles=default,dev,with_users", "-resource-root=/usr/local/bin/resources"]