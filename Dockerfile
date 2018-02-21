#
# Build the go server.
#
FROM golang:1.9-alpine as server

RUN apk add --no-cache \
  ca-certificates \
  git

RUN go get -u github.com/golang/dep/...
COPY ./server /go/src/github.com/phase2/outrigger-dashboard
WORKDIR /go/src/github.com/phase2/outrigger-dashboard
RUN dep ensure -v && \
  GOOS=linux GOARCH=amd64 go build -o dist/outrigger-dashboard

#
# Build the frontend website.
#
FROM node:8-alpine as frontend

COPY ./frontend/package.json ./frontend/package-lock.json /code/frontend/
WORKDIR /code/frontend
RUN npm install
# Now copy the rest of the codebase. This split of the steps ensures that the
# npm install stage will only be repeated if the package.json or
# package-lock.json files have changes.
COPY ./frontend /code/frontend
RUN npm run build

#
# Build the operating container.
#
FROM alpine:3.6

COPY --from=server /go/src/github.com/phase2/outrigger-dashboard/dist/outrigger-dashboard /outrigger-dashboard
COPY --from=frontend /code/frontend/ /app

EXPOSE 80

ENTRYPOINT [ "/outrigger-dashboard" ]

# Copy operation resets permissions.
RUN chmod +x /outrigger-dashboard
