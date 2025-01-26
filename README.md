# account-service

## Overview

`account-service` is a Go-based microservice for managing bank accounts and transactions. It provides functionalities to create accounts, deposit and withdraw funds, and check account balances.

## Features

- Create a new account
- Deposit funds into an account
- Withdraw funds from an account
- Check the balance of an account

## Technologies Used

- Go
- Echo framework
- PostgreSQL
- Logrus for logging

## Installation

### Using Docker Compose

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/account-service.git
    cd account-service
    ```

2. Create a `docker-compose.yml` file with the following content:
    ```yaml
    version: '3.8'

    services:
      postgres:
        image: postgres:latest
        container_name: postgres
        environment:
          POSTGRES_USER: user
          POSTGRES_PASSWORD: password
          POSTGRES_DB: accountdb
        ports:
          - "5432:5432"
        volumes:
          - postgres_data:/var/lib/postgresql/data

      account-service:
        build: .
        container_name: account-service
        environment:
          DB_HOST: postgres
          DB_PORT: 5432
          DB_USER: user
          DB_PASSWORD: password
          DB_NAME: accountdb
        ports:
          - "8080:8080"
        depends_on:
          - postgres

    volumes:
      postgres_data:
    ```

3. Build and run the containers:
    ```sh
    docker-compose up --build
    ```

### Manual Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/account-service.git
    cd account-service
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

3. Set up the PostgreSQL database and update the connection string in [main.go](http://_vscodecontentref_/0).

4. Run the service:
    ```sh
    go run main.go
    ```

## Usage

1. The service will be available at `http://localhost:8080`.

## API Endpoints

### Create Account

- **URL:** `/daftar`
- **Method:** `POST`
- **Request Body:**
    ```json
    {
        "nama": "John Doe",
        "nik": "123456789",
        "no_hp": "08123456789"
    }
    ```
- **Response:**
    ```json
    {
        "no_rekening": "generated-account-number"
    }
    ```

### Deposit Funds

- **URL:** `/tabung`
- **Method:** `POST`
- **Request Body:**
    ```json
    {
        "no_rekening": "generated-account-number",
        "nominal": 500.0
    }
    ```
- **Response:**
    ```json
    {
        "saldo": 500.0
    }
    ```

### Withdraw Funds

- **URL:** `/tarik`
- **Method:** `POST`
- **Request Body:**
    ```json
    {
        "no_rekening": "generated-account-number",
        "nominal": 200.0
    }
    ```
- **Response:**
    ```json
    {
        "saldo": 300.0
    }
    ```

### Check Balance

- **URL:** `/saldo/:no_rekening`
- **Method:** `GET`
- **Response:**
    ```json
    {
        "saldo": 1300.0
    }
    ```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## License

This project is licensed under the MIT License.