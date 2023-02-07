# startup

idéia do projeto é ser uma central de controle e organização para um pequeno vendedor que as vezes se perde em pra quem vendeu, quando vendeu, se pagou, se não pagou

## Funcionalidades

- cadastro e login de usuários
- crud de clientes
- crud de produtos
- crud de vendas
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
- desenvolver o backend com go
- desenvolver o frontend com react