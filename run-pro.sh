cd  /opt/app/
docker-compose stop app swagger
docker-compose up --force-recreate --build --remove-orphans -d app swagger

