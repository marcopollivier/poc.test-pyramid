# Bruno API Tests

Collection de testes para a API de desconto da POC Test Pyramid.

## Como usar

1. Instale o Bruno: https://www.usebruno.com/
2. Abra a pasta `_bruno` no Bruno
3. Execute os testes

## Endpoints disponíveis

- **POST /discount** - Calcula desconto via JSON
- **GET /discount** - Calcula desconto via query params

## Cenários de teste

- ✅ Desconto válido (10%)
- ✅ Desconto válido (50%)
- ❌ Preço negativo
- ❌ Desconto inválido (>100%)
- ❌ Parâmetros inválidos no GET

## Ambiente

- **Local**: http://localhost:8080
