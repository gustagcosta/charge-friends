# charge-friends

projetinho dos cria pra registrar quem ta te devendo e mandar email e whatsapp pra pessoa lembrar de te pagar de forma automatizada

## Funcionalidades

- cadastro e login de usuários (2 rotas)
- crud de cobranças (4 rotas)
- notificações, será disparada manualmente via rota ou em um serviço em background
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

devedor
- id
- nome
- email
- whatsapp
