ARG builder_img="node:12.13.1-alpine"
ARG runtime_img="nginx:1.16.0-alpine"

FROM $builder_img as builder

WORKDIR /app

ENV PATH /app/node_modules/.bin:$PATH

COPY service-status/*.json /app/
RUN rm *-lock.json \
    && npm install \
    && npm install react-scripts@3.3.0 -g

COPY service-status/public public
COPY service-status/src src
RUN npm run build

FROM $runtime_img
COPY --from=builder /app/build /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh
EXPOSE 80
ENTRYPOINT ["/entrypoint.sh"]
CMD ["nginx", "-g", "daemon off;"]