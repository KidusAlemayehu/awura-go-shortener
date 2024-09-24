
# Awura URL Shortener

## Setup Instructions

### Docker Compose Commands

1. **Build and Start the Application:**

   ```
   docker compose build

   docker-compose up 
   ```

2. **Stop the Application:**

   ```
   docker-compose down
   ```

### Environment Variables

Create a `.env` file in the root directory with the following content for the database connection string:

```
DB_CONNECTION_STRING="host={db_name} user={User} password={Password} dbname={Database Name} sslmode=disable"
```

### Sample Request for Shortening a URL

- **URL:** `http://localhost:18080/shorten`
- **Method:** `POST`
- **Body:**

  ```
  {
    "original_url": "https://www.example.com"
  }
  ```

**Response:**

```
{
  "short_url": "http://localhost:18080/r/abc12345"
}
```
