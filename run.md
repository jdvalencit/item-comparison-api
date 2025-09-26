## Run Instructions

1. **Prerequisites**:  
   - Go version **1.21.5** or higher installed.

2. **Clone the repository**:
   ```sh
   git clone https://github.com/jdvalencit/item-comparison-api.git
   cd item-comparison-api
   ```

3. **Install dependencies**:
   ```sh
   go mod tidy
   ```

4. **Configure environment**:
The project uses a config.env file to define runtime configuration.
Currently, it includes the directory where products are stored:

   ```env
    DATA_DIR=data
   ```

    `DATA_DIR` → path to the folder where product JSON files are saved and retrieved.
    You can change this value to point to a custom directory if needed.

5. **Run the server**:
   ```sh
   go run ./cmd/server/main.go
   ```

6. **Access the API**:
By default, the server starts on port 8080:
   ```raw
   http://localhost:8080
   ```

---