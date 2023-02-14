# charge-friends

projetinho dos cria pra registrar quem ta te devendo e mandar email e whatsapp pra pessoa lembrar de te pagar de forma automatizada

## Requisitos
- golang 1.19
- docker + docker compose
- python + pip
- aws-cli 
- aws-local 

## Comandos
- cp app.env.example app.env : setup nas variaveis de ambiente, devem estar de acordo com o docker-compose
- docker-compose up -d : subir os projetos
- aws configure : configurar a aws, se atentar para região ser a mesma do app.env, outros dados não importam
- awslocal sqs create-queue --queue-name charge-friends-queue: criar fila

## Funcionalidades

- cadastro e login de usuários
- crud de cobranças 
- crud de clientes
- notificações, disparo manual ou agendado
    - via whatsapp 
    - via e-mail 

## Entidades

usuario
- id
- nome 
- e-mail
- password
- chave pix
- data de cadastro

cobrança
- id
- valor
- observação
- data para cobrança
- status (paid, unpaid)
- id devedor
- id usuario

cliente
- id
- nome
- email
- whatsapp
