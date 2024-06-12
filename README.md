## Api Rest Feita em GO

Este Repo é um exemplo de uma API RESTful escrita em Go usando o framework Gin. 
A API se conecta a um banco de dados PostgreSQL e implementa uma estrutura básica de repositório, usecase e controller.

## Pré-requisitos
 - Go 1.16 ou superior
 - Docker (opcional, para rodar o PostgreSQL via Docker)
 - PostgreSQL instalado e configurado


## Configuração do Banco de Dados
Se você tiver o Docker instalado, pode rodar o PostgreSQL com o seguinte comando:
`docker run --name go_db -e POSTGRES_PASSWORD=1234 -e POSTGRES_DB=postgres -p 5432:5432 -d postgres:12`

A API estará disponível em http://localhost:8000.
