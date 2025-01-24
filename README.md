# labs-auction-fullcycle

## Descrição desafio
Objetivo: Adicionar uma nova funcionalidade ao projeto já existente para o leilão fechar automaticamente a partir de um tempo definido.

- Uma função que irá calcular o tempo do leilão, baseado em parâmetros previamente definidos em variáveis de ambiente;
- Uma nova go routine que validará a existência de um leilão (auction) vencido (que o tempo já se esgotou) e que deverá realizar o update, fechando o leilão (auction);
- Um teste para validar se o fechamento está acontecendo de forma automatizada;

## Como executar
```bash
make build
```

## Como testar
1 - Setar a variável de ambiente AUCTION_TIME no arquivo [.env](cmd/auction/.env) com o tempo que deseja que o leilão fique aberto.

2- Criar o leilão de um produto.
```curl
curl --location 'localhost:8080/auction' \
--header 'Content-Type: application/json' \
--data '{
    "product_name": "Fusca",
    "category": "Carro",
    "description": "Fusca 1980",
    "condition": 0
}'
```
3- Verificar se o produto foi criado.
```curl
curl --location 'http://localhost:8080/auction?status=0' \
--header 'Content-Type: application/json'
```
4- Após esperar o periodo que foi definido na variável de ambiente AUCTION_TIME, verificar se o produto foi fechado.
```curl
curl --location 'http://localhost:8080/auction?status=1' \
--header 'Content-Type: application/json'
```