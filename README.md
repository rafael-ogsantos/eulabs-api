## Isso é um teste pratico contendo um CRUD de produtos para a empresa EULABS

## Como rodar o projeto
1. Clone o projeto `git clone git@github.com:rafael-ogsantos/eulabs-api.git`
2. Entre na pasta do projeto `cd eulabs-api`
3. Execute o comando `docker-compose up -d` para subir os containers e faça requisições para a API em `http://localhost:8080/api/products`

## Como rodar os testes
Execute o comando `docker-compose exec app go test ./... -v`

## Para ler a documentação da API
Execute o comando `godoc -http=:6060 ` e acesse `http://localhost:6060/pkg/github.com/rafael-ogsantos/eulabs-api/`


