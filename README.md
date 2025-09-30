# POC Test Pyramid

Proof of Concept para demonstrar a pirâmide de testes em Go com PostgreSQL e Kafka.

## Pré-requisitos

- Docker e Docker Compose
- Go 1.21+

## Como executar

### 1. Subir infraestrutura
```bash
make up
```

### 2. Executar aplicação
```bash
make run
```

### 3. Parar tudo
```bash
make down
```

## Testes

### Testes Unitários
```bash
go test ./service/... -v
```

## API Endpoints

### POST /discount
```json
{
  "price": 100.0,
  "discount": 10.0
}
```

### GET /discount?price=100&discount=10

## Estrutura do projeto

- `main.go` - Ponto de entrada da aplicação
- `service/` - Camada de lógica de negócio
- `handler/` - Handlers da API REST
- `repository/` - Camada de persistência
- `messaging/` - Publisher do Kafka
- `model/` - Modelos de dados
- `config/` - Configurações
- `docker-compose.yml` - Infraestrutura local

## Infraestrutura

- **PostgreSQL**: Banco de dados na porta 5432
- **Kafka**: Message broker na porta 9092 (KRaft mode)
