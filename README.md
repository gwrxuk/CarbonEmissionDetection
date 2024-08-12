# Carbon Emission Detection

This project is a Go-based web application that interacts with NASA's satellite images to detect and query carbon emission activities. The application uses the Go GenAI SDK to communicate with AI models, analyzing the satellite imagery and providing detailed reports on detected carbon emissions.

## Features

- **RESTful API**: The application provides a RESTful API for managing items, including creation, retrieval, updating, and deletion.
- **Image Processing**: Downloads satellite images from provided URLs, processes them, and interacts with AI models to detect carbon emission activities.
- **Web Interface**: A simple web interface is included for uploading and displaying images, powered by Go templates.

## Project Structure

- `main.go`: The entry point of the application, initializing routes and starting the server.
- `handlers/item.go`: Contains HTTP handlers for managing items in the application, including creating, updating, and deleting items.
- `models/item.go`: Defines the data model for items, which includes the properties and methods related to items.
- `routers/router.go`: Configures the routes for the application, mapping HTTP endpoints to their corresponding handlers.
- `utils/helper.go`: Contains utility functions for image processing and interaction with the AI models.

## Setup and Installation

### Prerequisites

- Go 1.16+
- Git

### Steps

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/CarbonEmissionDetection.git
   cd CarbonEmissionDetection
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Run the application:

   ```bash
   go run main.go
   ```

   The server will start on `http://localhost:8001`.

### API Endpoints

- `GET /items`: Retrieve all items.
- `GET /items/{id}`: Retrieve a specific item by its ID.
- `POST /items`: Create a new item.
- `PUT /items/{id}`: Update an existing item.
- `DELETE /items/{id}`: Delete an item by its ID.

### Web Interface

Access the web interface at `http://localhost:8001` to upload images and view the results of carbon emission detection.

## Future Enhancements

- Improve the AI model's accuracy in detecting carbon emissions.
- Add more detailed logging and error handling.
- Implement authentication and authorization for the API endpoints.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
