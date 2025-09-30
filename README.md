# POC Test Pyramid

Proof of Concept para demonstrar a pirâmide de testes em Go.

## Como executar

```bash
go run main.go
```

## Testes

### Testes Unitários
```bash
go test ./service/...
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
