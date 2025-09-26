# Item Comparison API

A backend API for retrieving, saving, and managing products to enable item comparison. Products are stored as JSON files for simplicity and easy access.

This project provides a lightweight backend API for managing product data, designed to support item comparison features. The API exposes endpoints to retrieve, create, update, and delete products.

For simplicity, all products are stored as JSON files, which keeps the data human-readable, version-friendly, and easily accessible without requiring a database.

---

## Table of Contents

- [Project Structure](#project-structure)
- [Usage](#usage)
  - [List Products](#list-products)
  - [Save Product](#save-product)
  - [Retrieve Product](#retrieve-product)
  - [Update Product](#update-product)
  - [Delete Product](#delete-product)
- [Data Storage](#data-storage)

---

## Project Structure

```
item-comparison-api/
├── api/         # HTTP handlers and routing
├── services/    # Business logic for product management
├── repository/  # Data access layer (read/write JSON files)
├── dto/         # Data Transfer Objects (request/response payloads)
├── models/      # Core domain models (e.g., Product)
├── data/        # JSON files storing product data
├── main.go      # Application entry point
```

- **api/**: Defines HTTP endpoints and request handling logic.
- **services/**: Contains business logic for comparing and managing products.
- **repository/**: Handles reading and writing product data to JSON files.
- **dto/**: Structures for API request and response payloads.
- **models/**: Core data models used throughout the application.
- **data/**: Directory where product data is persisted as JSON files.

---

## Usage

### List Products

**Request**
```http
GET /products
```

**Response**
```json
[
  {
    "id": "123",
    "name": "Product A",
    "price": 100.0,
    "brand": "Product brand",
    "image_url": "http://example.com/producta.jpg",
    "rating": 4.2,
    "specifications": {
      "color": "red",
      "size": "M"
    }
  },
  {
    "id": "612",
    "name": "Product B",
    "price": 20.0,
    "brand": "B brand",
    "image_url": "http://example.com/productb.jpg",
    "rating": 3.7,
    "specifications": {
      "color": "blue",
      "type": "wireless"
    }
  }
]
```

### Save Product

**Request**
```http
POST /products
Content-Type: application/json
X-Seller-ID: 123

[
  {
    "id": 8,
    "name": "Product C",
    "description": "Good cleaning product",
    "price": 8,
    "brand": "C brand",
    "image_url": "http://example.com/productc.jpg",
    "rating": 4,
    "specifications": {
        "color": "Green"
    }
  }
]
```

**Response**
```raw
    Product saved successfully
    201 Created
```

### Retrieve Product

**Request**
```http
GET /products/8
```

**Response**
```json
{
    "id": 8,
    "name": "Product C",
    "description": "Good cleaning product",
    "price": 8,
    "brand": "C brand",
    "image_url": "http://example.com/productc.jpg",
    "rating": 4,
    "specifications": {
        "color": "Green"
    },
    "seller_id": "123"
}
```
### Update Product

**Request**
```http
PUT /products
```

**Request**
```http
PUT /products
Content-Type: application/json
X-Seller-ID: 123

[
  {
    "id": 8,
    "name": "Product C",
    "description": "Good cleaning product",
    "price": 800,
    "brand": "C brand",
    "image_url": "http://example.com/productc.jpg",
    "rating": 5,
    "specifications": {
        "color": "Purple"
    }
  }
]
```

**Response**
```raw
    Products updated successfully
    200 OK
```

### Delete Product

**Request**
```http
DELETE /products/8
```

**Response**
```raw
    Product deleted successfully
    200 OK
```

---

## 📂 Data Storage

Product information is persisted as individual JSON files within the `./data` directory. Each product is stored in its own file, which makes the data:
- Simple to manage – no database setup or migration is required.
- Human-readable – files can be easily inspected and edited.
- Portable – product data can be versioned, backed up, or shared without extra tooling.

---

## 🚀 Future Improvements

Due to limited time, some features and refinements were not implemented in the current version but are important to keep in mind for future iterations:
- Centralized error handling – create a dedicated error handler function to standardize responses and improve maintainability.
- Structured logging – integrate a specialized logging library for better log formatting, levels, and traceability.
- Extended comparison features – enhance the API to support richer product comparison logic (e.g., multi-criteria filters, side-by-side output).
- Validation improvements – add stricter input validation rules and reusable validators.
- Authentication & Authorization – implement proper access controls (e.g., JWT) to ensure only authorized sellers can manage their products.
- Automated testing – add unit tests and integration tests to ensure API reliability and prevent regressions.
- Database integration – optionally migrate from JSON file storage to a database for scalability and performance.