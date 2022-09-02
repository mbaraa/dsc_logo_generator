FROM alpine:latest as build

RUN apk add cairo cairo-dev go git

WORKDIR /app
COPY . .

RUN go get
RUN go build -ldflags="-w -s"

FROM alpine:latest as run

RUN apk add cairo

WORKDIR /app

COPY --from=build /app/b64_files.json ./
COPY --from=build /app/dsc_logo_generator ./run
COPY --from=build /app/client/static/fonts/* /usr/share/fonts/

EXPOSE 1105

CMD ["./run"]
