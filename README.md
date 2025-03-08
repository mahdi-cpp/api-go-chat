# api-go-chat

# docker run -d --name web-portainer --restart always -p 9000:9000 -v /var/run/docker.sock:/var/run/docker.sock -v portainer_data:/data portainer/portainer-ce

# docker run -d --name postgres --restart always -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=admin@123456 -e POSTGRES_DB=mqtt -p 5432:5432 postgres

#  docker run -d --name redis_service --restart always -p 6379:6379 redis_service
#  docker run -d --name nginx --restart=always -p 8081:80 -v /var/cloud:/usr/share/nginx/html nginx

# docker run  -d  --name nginx  --restart=always --network host  -p 8081:80  -v /home/mahdi/files:/usr/share/nginx/html:ro  nginx
# docker run -d --network host --name emqx emqx/emqx:latest