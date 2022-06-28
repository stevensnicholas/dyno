FROM node:16.15.0 as builder
WORKDIR /app
RUN apt update -y && \
    apt install -y libgtk2.0-0 libgtk-3-0 libgbm-dev libnotify-dev \
    libgconf-2-4 libnss3 libxss1 libasound2 libxtst6 xauth xvfb
COPY . .
RUN npm ci
