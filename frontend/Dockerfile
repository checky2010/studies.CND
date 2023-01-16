# Use debian as a base for the builder
FROM debian:latest AS builder

# Install dependencies for flutter build
RUN apt-get update
RUN apt-get install -y curl git wget unzip libgconf-2-4 gdb libstdc++6 libglu1-mesa fonts-droid-fallback lib32stdc++6 python3 sed
RUN apt-get clean

# Downloat flutter and add it to the PATH
RUN git clone https://github.com/flutter/flutter.git /usr/local/flutter
ENV PATH="${PATH}:/usr/local/flutter/bin:/usr/local/flutter/bin/cache/dart-sdk/bin"

# Check that flutter is working and upgrade the master channel
RUN flutter doctor -v
RUN flutter channel master
RUN flutter upgrade

# Create a directory, enter it, and copy the source code
WORKDIR /app
COPY ./ ./

# Set the API_URL as ARG and build the service with it
ARG API_URL
RUN flutter build web --dart-define=API_URL=$API_URL

# Use a small nginx image as base for the deployment
FROM nginx:1.23.3-alpine AS app

# Expose port 80 and copy the files to nginxs' html directory
EXPOSE 80
COPY --from=builder /app/build/web /usr/share/nginx/html