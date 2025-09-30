# Dev Containers

## O que é

Dev Containers permite desenvolver dentro de um container Docker diretamente no VS Code, garantindo ambiente consistente e isolado.

## Como funciona

1. VS Code detecta arquivo `.devcontainer/devcontainer.json`
2. Constrói/usa imagem Docker especificada
3. Monta seu código dentro do container
4. Conecta VS Code ao ambiente containerizado
5. Extensões e ferramentas rodam dentro do container

## Vantagens para este projeto

- **Ambiente consistente**: PostgreSQL + Kafka sempre disponíveis
- **Isolamento**: Não polui sua máquina local
- **Testes mais confiáveis**: Infraestrutura sempre no mesmo estado
- **Onboarding fácil**: Qualquer dev só precisa abrir no VS Code
- **CI/CD**: Mesmo ambiente local e produção

## Como usar no VS Code

1. Instala extensão "Dev Containers"
2. Abre projeto no VS Code
3. Popup aparece: "Reopen in Container" 
4. Ou: `Cmd+Shift+P` → "Dev Containers: Reopen in Container"
5. VS Code reconstrói ambiente e conecta

## Para testes de integração

- PostgreSQL e Kafka já estarão rodando
- Testes podem conectar diretamente sem Docker Compose
- Ambiente limpo a cada rebuild
- Infraestrutura sempre no mesmo estado inicial
