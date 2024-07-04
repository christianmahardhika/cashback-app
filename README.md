# cashback-app

## Description

This project is a simple cashback application that allows users to register, login, and view their cashback balance. The application also allows users to view their cashback balance in different currencies.

## Features
- Tenant calculate cashback based on the user's total purchase amount and promocode.
- The application will have business rules to calculate the cashback based on the user's total purchase amount and promocode.

```mermaid
flowchart TD
    start((Start)) --> A(Tenant APP / Service)
    A -->|send request to calculate items on cart and applied promo code| B(Cashback App)
    B -->|calculate cashback| B
    B -->|return cashback final amount| C(Tenant APP / Service)
    C -->|display cashback| D(User)
    D -->|End| END((End))
```

## Architecture Highlights

- The application is built using (GoFr)[https://gofr.dev/docs/quick-start/introduction] which is a framework that allows developers to build microservices in Go with ease.

- the application is contain 2 services:
    - `admin service` which is responsible for CMS operations.
    - `core service` which is responsible for the core business logic of the application.
- Both services are connected to a database.

```mermaid
graph TD
    A[admin service] -->|connects to| B[database]
    C[core service] -->|connects to| B[database]
```