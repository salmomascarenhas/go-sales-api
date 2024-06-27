# go-sales-api 🎫

## Descrição 📄
A `go-sales-api` é uma API para vendas de ingressos, onde parceiros fornecem informações e vendemos seus ingressos. A aplicação gerencia eventos, assentos e ingressos.

## Requisitos 📋
- Go 1.16+
- Docker
- Docker Compose
- MySQL

## Instalação 🛠️
1. Clone o repositório:
   ```bash
   git clone https://github.com/salmomascarenhas/go-sales-api.git
   cd go-sales-api
   ```

2. Configure o ambiente:
   - Edite `docker-compose.yaml` conforme necessário.
   - Verifique as configurações do banco de dados em `mysql-init/init.sql`.

3. Suba os containers Docker:
   ```bash
   docker-compose up -d
   ```

## Executando a Aplicação ▶️
Você pode utilizar o Makefile para facilitar a execução da aplicação:

1. Inicialize o banco de dados:
   ```bash
   make init
   ```

2. Gere os arquivos do Swagger:
   ```bash
   make swag
   ```

3. Compile o código:
   ```bash
   make build
   ```

4. Execute o servidor:
   ```bash
   make run
   ```

## Estrutura do Projeto 🗂️
- **cmd/**: Contém os arquivos principais para iniciar a aplicação.
  - **events/main.go**: Ponto de entrada da aplicação.

- **docs/**: Documentação e arquivos Swagger.

- **internal/events/**: Código da aplicação, incluindo lógica de negócios e manipulação de dados.
  - **domain/**: Contém as definições de domínios como `Event`, `Spot` e `Ticket`.
    - **event.go**: Definição da estrutura e métodos de eventos.
    - **spot.go**: Definição da estrutura e métodos de assentos.
    - **ticket.go**: Definição da estrutura e métodos de ingressos.
  - **infra/**: Infraestrutura, incluindo `http` handlers e repositórios.
    - **http/**: Handlers HTTP para as rotas da aplicação.
    - **repository/**: Implementação dos repositórios.
    - **service/**: Serviços auxiliares e integração com parceiros.
  - **usecase/**: Casos de uso da aplicação.
    - **buy_tickets.go**: Caso de uso para compra de ingressos.
    - **create_event.go**: Caso de uso para criação de eventos.
    - **create_spots.go**: Caso de uso para criação de assentos.
    - **get_event.go**: Caso de uso para obtenção de detalhes de eventos.
    - **list_events.go**: Caso de uso para listar eventos.
    - **list_spots.go**: Caso de uso para listar assentos.

- **mysql-init/**: Scripts de inicialização do banco de dados MySQL.
  - **init.sql**: Script SQL para inicialização do banco de dados.

- **Makefile**: Arquivo para automatização de tarefas comuns no desenvolvimento.

## Rotas 🌐
### Eventos
- **GET /events**: Lista todos os eventos.
- **POST /events**: Cria um novo evento.
- **GET /events/{eventID}**: Detalhes de um evento específico.
- **GET /events/{eventID}/spots**: Lista todos os assentos de um evento.
- **POST /events/{eventID}/spots**: Cria novos assentos para um evento.

### Compra de Ingressos
- **POST /checkout**: Compra ingressos para um evento específico.

## Notas 📌
- Certifique-se de que todas as dependências estão instaladas antes de executar a aplicação.
- Consulte a documentação oficial do Go para dúvidas sobre o ambiente de desenvolvimento.

Para mais detalhes, você pode visitar o repositório no [GitHub](https://github.com/salmomascarenhas/go-sales-api).
