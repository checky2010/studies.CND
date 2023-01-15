FROM debian:latest AS builder
RUN apt-get update
RUN apt-get install -y curl git wget unzip libgconf-2-4 gdb libstdc++6 libglu1-mesa fonts-droid-fallback lib32stdc++6 python3 sed
RUN apt-get clean

RUN git clone https://github.com/flutter/flutter.git /usr/local/flutter

ENV PATH="${PATH}:/usr/local/flutter/bin:/usr/local/flutter/bin/cache/dart-sdk/bin"

RUN flutter doctor -v
RUN flutter channel master
RUN flutter upgrade

RUN mkdir /app
WORKDIR /app
COPY ./ /app

ARG API_URL
RUN flutter build web --dart-define=API_URL=$API_URL

FROM nginx:1.23.3-alpine AS app
EXPOSE 80
COPY --from=builder /app/build/web /usr/share/nginx/html