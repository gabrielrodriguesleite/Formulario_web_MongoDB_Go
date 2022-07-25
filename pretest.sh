if [ `docker ps | grep -i formulario_web_mongodb_go_mongo --count` -lt 1 ]

# inicia o container se este não existir
then 
  docker-compose up -d;
  clear;
  echo "\nClick to open: http://localhost:8080\n\r";
  echo "Ctrl+C to kill server\n\r"
  echo 'run "go generate" again to remove container\n\r'
  go run main.go

# caso exista remove o container
else

  docker-compose down --remove-orphans

fi
