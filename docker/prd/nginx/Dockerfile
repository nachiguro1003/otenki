FROM nginx:latest
COPY ./docker/prd/nginx/conf/nginx.conf /etc/nginx/nginx.conf
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
