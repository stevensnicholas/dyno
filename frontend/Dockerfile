FROM node:16.15.0 as builder
WORKDIR /app
COPY . .
RUN npm ci
RUN npm run build
RUN echo '{ "backend": "http://localhost:8080" }' > build/settings.json

FROM nginx:1.21.6-alpine
COPY nginx.conf /etc/nginx/nginx.conf
RUN rm -rf /usr/share/nginx/html/*
COPY --from=builder /app/build /usr/share/nginx/html
COPY tests-entrypoint.sh /tests-entrypoint.sh
EXPOSE 8080
HEALTHCHECK --interval=1s --timeout=1s --start-period=2s --retries=3 CMD [ "curl", "--fail", "http://localhost:8080" ]
