cd  /opt/app/
chmod +x run.sh
docker-compose stop app swagger
docker-compose up --build -d app swagger