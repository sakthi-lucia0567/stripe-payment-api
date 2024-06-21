# Stripe API Integration

This repository provides an API service to interact with Stripe for creating payment intents, capturing payment intents, creating refunds, and fetching payment intents. The service is built using Go and the `chi` router for handling HTTP requests.

## Features

- Health check endpoint
- Error testing endpoint
- Create payment intent
- Capture payment intent
- Create refund
- Fetch all payment intents

## Prerequisites

- Go 1.16+
- Stripe API Key
- `.env` file for environment variables
- `PORT` your desired port
- `STRIPE_KEY` your stripe api key.

## Getting Started

### Clone the Repository

```sh
git clone https://github.com/sakthi-lucia0567/stripe-payment-api.git
cd stripe-api-integration
```

### Install Dependencies

Ensure you have `go-chi`, `godotenv`, and `go-chi/cors` installed. You can install them using:

```sh
go get github.com/go-chi/chi/v5
go get github.com/go-chi/cors
go get github.com/joho/godotenv
```

### Environment Variables

Create a `.env` file in the root of your project and add the following environment variables:

```plaintext
PORT=your_port_number
STRIPE_API_KEY=your_stripe_api_key
```

### Run the Server

```sh
go run main.go
```

## Endpoints

### Health Check

- **URL:** `/api/v1/healthcheck`
- **Method:** `GET`
- **Description:** Checks if the server is running.

### Error Testing

- **URL:** `/api/v1/err`
- **Method:** `GET`
- **Description:** Endpoint to test error handling.

### Create Payment Intent

- **URL:** `/api/v1/create_intent`
- **Method:** `POST`
- **Description:** Creates a new payment intent in Stripe.

### Capture Payment Intent

- **URL:** `/api/v1/capture_intent/{id}/capture`
- **Method:** `POST`
- **Description:** Captures a payment intent using its ID.

### Create Refund

- **URL:** `/api/v1/create_refund/{id}`
- **Method:** `POST`
- **Description:** Creates a refund for a payment intent using its ID.

### Fetch Payment Intents

- **URL:** `/api/v1/get_intents`
- **Method:** `GET`
- **Description:** Fetches all payment intents from Stripe.

## Middleware

The server uses the following middleware:

- `Logger`: Logs incoming HTTP requests.
- `Recoverer`: Recovers from panics and writes a 500 HTTP response.
- `CORS`: Configures CORS headers.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contributions

Contributions are welcome! Please fork the repository and create a pull request.

## Contact

For any inquiries or issues, please contact [sakthivela0567@gmail.com](mailto:sakthivela0567@gmail.com).
