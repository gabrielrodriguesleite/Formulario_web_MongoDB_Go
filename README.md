# Formulario_web_MongoDB_Go
Um formulário com dados mantidos em mongoDB escrito em Go

## Para rodar a aplicação é necessário ter o MongoDB rodando

Uma opção é rodar a partir do docker com o arquivo docker-compose provido no projeto

##### PARA SUBIR O CONTAINER:

`docker-compose up`

Depois disso o banco deve estar disponível em http://localhost:27017

A seguinte mensagem deve ser exibida:

`It looks like you are trying to access MongoDB over HTTP on the native driver port.`

###### PARA FINALIZAR O CONTAINER:

`docker-compose down --remove-orphans`

## Rodar a aplicação

Na pasta raiz execute

`go run main.go`

Abra on navegador em http://localhost:8080

Preencha os dados se tudo der certo haverá no banco de dados "webform" um novo documento em"newsletter"

## Para conectar ao banco utilize a seguinte URI mongodb://root:12345678@localhost:27017
