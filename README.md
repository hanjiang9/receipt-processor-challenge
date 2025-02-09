# receipt-processor-challenge

This is a simple web service for processing receipts and calculating points based on specific rules. The service provides endpoints to process receipts and retrieve points for a given receipt ID.

## Features

- **Process Receipts**: Submit a receipt to calculate and store points.
- **Retrieve Points**: Get the points for a specific receipt using its ID.

## Setup and Installation

### Prerequisites

- Docker
- Go 1.21 (if running locally without Docker)

### Running with Docker

1. **Build the Docker Image**:

   ```sh
   docker build -t receipt-processor .
   ```

2. **Run the Docker Container**:

   ```sh
   docker run -p 8080:8080 receipt-processor
   ```

## API Documentation

### Endpoints

#### POST `http://localhost:8080/receipts/process`

- **Description**: Processes a receipt and returns a unique ID for the receipt.
- **Request Body**: JSON object representing the receipt details.
- **Response**: JSON object containing the unique receipt ID.

**Request Example**:
```json
{
"retailer": "Target",
"purchaseDate": "2022-01-01",
"purchaseTime": "13:01",
"items": [
{
"shortDescription": "Mountain Dew 12PK",
"price": "6.49"
},
{
"shortDescription": "Emils Cheese Pizza",
"price": "12.25"
},
{
"shortDescription": "Knorr Creamy Chicken",
"price": "1.26"
},
{
"shortDescription": "Doritos Nacho Cheese",
"price": "3.35"
},
{
"shortDescription": "Klarbrunn 12-PK 12 FL OZ",
"price": "12.00"
}
],
"total": "35.35"
}

**Response Example**:
```json
{
"id": "7cdc04a3-d917-4131-bd80-64b56bdaf742"
}

#### GET `http://localhost:8080/receipts/{id}/points`

- **Description**: Retrieves the points for a specific receipt ID.
- **Response**: JSON object containing the points for the receipt.

**Response Example**:
```json
{
"points": 28
}
