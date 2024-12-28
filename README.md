# Retail Pulse Service

## Backend Intern Assignment

---

### 1. Description
The **Retail Pulse Service** is a Go-based backend application designed to process retail data, manage job assignments, and facilitate resource downloading (e.g., images). It is highly scalable and deployable using containerized solutions like Docker.  

This service processes images collected from retail stores and calculates their perimeters. The results are stored for reporting and analysis purposes. Key objectives include:  

- **Job Submission**: Accept and process jobs with store visit details and associated image URLs.
- **Image Processing**:
  - Download images from provided URLs.
  - Compute image perimeters as `2 * (Height + Width)`.
  - Simulate GPU processing with random delays (0.1 to 0.4 seconds).
- **Store Master Integration**: Validate `store_id` using metadata like `store_name` and `area_code`.
- **Job Status Monitoring**: Track job progress (`Completed`, `Ongoing`, `Failed`) and identify errors.

The system supports high concurrency, allowing simultaneous processing of multiple jobs with thousands of images while ensuring reliability and error reporting.

---

### 2. Assumptions

- **Data**: Input data is assumed to be valid and provided in a structured format (e.g., `StoreMasterAssignment.csv`).
- **Input Validation**: Jobs include correctly formatted JSON payloads, valid `store_id` values, and accessible image URLs.
- **Concurrency**: The service efficiently handles multiple concurrent jobs.
- **Error Scenarios**: Partial or complete job failures occur if images fail to download or `store_id` values are invalid.
- **Sleep Times**: Random delays simulate realistic GPU processing.

---

### 3. Installation and Testing Instructions

#### **Option 1: Using Docker**

1. **Prerequisites**:
   - Install [Docker](https://www.docker.com/) and Docker Compose.
   - Ensure ports specified in `docker-compose.yml` are available.

2. **Setup**:
   - Navigate to the project directory.
   - Build and start the containerized application:
     ```bash
     docker-compose build
     docker-compose up
     ```
   - Access the service via the endpoint specified in `docker-compose.yml`.

3. **Logs**:
   - Application logs are stored in `retail-pulse.log`. Monitor them using:
     ```bash
     tail -f retail-pulse.log
     ```

4. **Testing**:
   - Use sample data (`StoreMasterAssignment.csv`) for testing.
   - Test APIs using tools like [Postman](https://www.postman.com/) or `curl`.

---

#### **Option 2: Local Setup Without Docker**

1. **Install Dependencies**:
   - Ensure Go (version specified in `go.mod`) is installed. If not initialized, use:
     ```bash
     go mod init retail-pulse-service
     ```
   - Install dependencies:
     ```bash
     go mod tidy
     ```

2. **Run the Application**:
   - Start the application:
     ```bash
     go run main.go
     ```

3. **Testing**:
   - Test APIs locally at `http://localhost:8080`.
   - Use tools like Postman or `curl`.

---

### 4. Work Environment

- **Programming Language**: Go (Golang)
- **Dependency Management**: Go Modules (`go.mod`, `go.sum`)
- **Containerization**: Docker and Docker Compose
- **Editor/IDE**: Visual Studio Code with Go extensions

**Key APIs**:
- `/api/submit/`: Submit a new job.
- `/api/status`: Monitor job status.

---

### 5. Future Improvements

#### **Scalability**:
- Integrate GPU processing for large datasets.
- Add Kubernetes support for dynamic scaling.

#### **Error Handling**:
- Provide detailed error messages for image download failures.
- Retry failed downloads with exponential backoff.

#### **Logging and Monitoring**:
- Use centralized logging (e.g., ELK stack, Fluentd).
- Integrate metrics tools like Prometheus and Grafana.

#### **API Enhancements**:
- Add pagination to `/api/status` for jobs with numerous images.
- Provide a summary of processed and failed images.

#### **Frontend Dashboard**:
- Create a user-friendly dashboard to upload jobs and view results graphically.

#### **Resource Management**:
- Automate cleanup of old jobs and images.
- Shift image storage to cloud platforms like AWS S3 or Google Cloud Storage.

#### **Security**:
- Add JWT-based authentication for APIs.
- Enforce HTTPS for secure data transmission.

---

### 6. Project Structure

#### **Configuration and Setup**:
- `docker-compose.yml`: Multi-service deployment configuration.
- `Dockerfile`: Service containerization.
- `go.mod`, `go.sum`: Dependency management for Go.

#### **Main Application**:
- `main.go`: Application entry point.

#### **Logs and Data**:
- `retail-pulse.log`: Log file.
- `StoreMasterAssignment.csv`: Test/sample data file.

#### **Code Organization**:
- `api/job_handler.go`: API handlers.
- `models/`: Contains `job.go` and `store.go` models.
- `services/job_processor.go`: Core job processing logic.
- `utils/downloader.go`: Download utilities.
- `utils/logger.go`: Custom logging utilities.

#### **Resources**:
- `downloaded_images/`: Directory for downloaded images.

---
