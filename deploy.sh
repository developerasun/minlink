docker build -t minlink-backend:latest .
docker service rm minlink_backend
docker stack up -c ./service.yaml minlink
docker service ls