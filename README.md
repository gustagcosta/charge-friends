# startup

idéia do projeto é ser uma central de controle e organização para um pequeno vendedor que as vezes se perde em pra quem vendeu, quando vendeu, se pagou, se não pagou

## Funcionalidades

- cadastro e login de usuários (2 rotas)
- crud de clientes (4 rotas)
- crud de produtos (4 rotas)
- crud de vendas (4 rotas)
- relatórios
    - inadimplência
    - lucro por data
- aviso de notificações de inadimplência
    - via whatsapp 
    - via e-mail 

## Entidades

usuario
- id
- nome do dono
- nome do negócio
- e-mail
- password
- data de cadastro

cliente
- id
- nome
- whatsapp*
- e-mail*
- observação*

produto
- título
- descrição
- valor em centavos
- porcentagem de lucro

venda
- data
- data de pagamento
- cliente
- produtos

## Etapas

- sql do banco de dados
    - usar docker-compose e rodar o sql
- backend 
    - instalar tools
        - tasks go `go install github.com/go-task/task/v3/cmd/task@latest`
        - air `go install github.com/cosmtrek/air@latest`
    - base login register
    - cruds 
    - relatórios
    - notificações
- desenvolver o frontend com react
