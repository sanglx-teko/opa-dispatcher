FROM golang:1.12-alpine
WORKDIR /x-app
COPY opa-dispatcher /usr/local/bin/opa-dispatcher
COPY configs.json .
CMD [ "opa-dispatcher" ]
EXPOSE 1323

