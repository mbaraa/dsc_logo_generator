FROM alpine:latest AS build

RUN apk add cairo cairo-dev go git

WORKDIR /app
COPY . .

RUN go get
RUN go build -ldflags="-w -s"

FROM alpine:latest AS run

RUN apk add cairo

WORKDIR /app

COPY --from=build /app/b64_files.json ./
COPY --from=build /app/dsc_logo_generator ./dsc_logo_generator
COPY --from=build /app/client/static/fonts/* /usr/share/fonts/

EXPOSE 1105

CMD ["./dsc_logo_generator"]
