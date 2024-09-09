FROM node:18.15.0-alpine as builder_front

WORKDIR /build
ADD front/latest .
RUN npm install &&\
    npm run build

FROM golang:1.21.12 as builder_back

WORKDIR /build
ADD . .
RUN make build-docker &&\
    cd bin &&\
    mv *_docker webhook

FROM alpine:3.15

WORKDIR /webhook
COPY --from=builder_back /build/bin .
COPY --from=builder_front /build/dist ./front
COPY examples/config.yml /etc/webhook/config.yml

CMD ["./webhook"]
           