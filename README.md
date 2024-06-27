# go-sales-api 🎫

## Description 📄
`go-sales-api` is an API for ticket sales, where partners provide information and we sell their tickets. The application manages events, seats, and tickets.

## Requirements 📋
- Go 1.16+
- Docker
- Docker Compose
- MySQL

## Installation 🛠️
1. Clone the repository:
   ```bash
   git clone https://github.com/salmomascarenhas/go-sales-api.git
   cd go-sales-api
   ```

2. Configure the environment:
   - Edit `docker-compose.yaml` as needed.
   - Check the database settings in `mysql-init/init.sql`.

3. Start the Docker containers:
   ```bash
   docker-compose up -d
   ```

## Running the Application ▶️
You can use the Makefile to simplify running the application:

1. Initialize the database:
   ```bash
   make init
   ```

2. Generate Swagger files:
   ```bash
   make swag
   ```

3. Compile the code:
   ```bash
   make build
   ```

4. Run the server:
   ```bash
   make run
   ```

## Project Structure 🗂️
- **cmd/**: Contains the main files to start the application.
  - **events/main.go**: Application entry point.

- **docs/**: Documentation and Swagger files.

- **internal/events/**: Application code, including business logic and data handling.
  - **domain/**: Contains domain definitions like `Event`, `Spot`, and `Ticket`.
    - **event.go**: Structure and methods for events.
    - **spot.go**: Structure and methods for seats.
    - **ticket.go**: Structure and methods for tickets.
  - **infra/**: Infrastructure, including `http` handlers and repositories.
    - **http/**: HTTP handlers for application routes.
    - **repository/**: Repository implementations.
    - **service/**: Auxiliary services and partner integration.
  - **usecase/**: Application use cases.
    - **buy_tickets.go**: Use case for ticket purchase.
    - **create_event.go**: Use case for event creation.
    - **create_spots.go**: Use case for seat creation.
    - **get_event.go**: Use case for retrieving event details.
    - **list_events.go**: Use case for listing events.
    - **list_spots.go**: Use case for listing seats.

- **mysql-init/**: MySQL database initialization scripts.
  - **init.sql**: SQL script for database initialization.

- **Makefile**: File for automating common development tasks.

## Routes 🌐
### Events
- **GET /events**: Lists all events.
- **POST /events**: Creates a new event.
- **GET /events/{eventID}**: Details of a specific event.
- **GET /events/{eventID}/spots**: Lists all seats for an event.
- **POST /events/{eventID}/spots**: Creates new seats for an event.

### Ticket Purchase
- **POST /checkout**: Purchases tickets for a specific event.

## Notes 📌
- Ensure all dependencies are installed before running the application.
- Refer to the official Go documentation for development environment queries.

For more details, you can visit the repository on [GitHub](https://github.com/salmomascarenhas/go-sales-api).
