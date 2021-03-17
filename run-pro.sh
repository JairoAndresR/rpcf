cd  /opt/app/
chmod +x run.sh
sh run.sh
docker-compose stop app swagger
docker-compose up --force-recreate --build --remove-orphans -d app swagger
