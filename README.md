# Google Cloud Run Proxy

[![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![Docker](https://img.shields.io/badge/Docker-Enabled-2496ED?style=flat&logo=docker)](https://www.docker.com/)
[![License: GPL-3.0](https://img.shields.io/badge/License-GPLv3-yellow.svg?style=for-the-badge&logo=gplv3&logoColor=white)](https://opensource.org/license/gpl-3-0)

A lightweight, high-performance TCP proxy server written in Go. Designed for seamless deployment on Google Cloud Run, this proxy efficiently forwards traffic to a specified target server (e.g., a V2Ray server), making it an ideal solution for bypassing network restrictions or masking backend services.

## 🚀 Features

- **Lightweight & Fast**: Built with Go for minimal resource footprint and high concurrency.
- **Cloud Run Ready**: Optimized for stateless, containerized environments like Google Cloud Run.
- **Docker Support**: Includes a multi-stage Dockerfile for easy building and deployment.
- **Simple Configuration**: Configurable via environment variables.

## 📋 Prerequisites

- **Go**: Version 1.21 or higher (for local development).
- **Docker**: For containerized deployment.
- **Google Cloud SDK**: If deploying to Google Cloud Run.

## ⚙️ Configuration

The application is configured using environment variables:

| Variable | Description | Default |
| :--- | :--- | :--- |
| `PORT` | The port the server listens on. | `8080` (or provided by Cloud Run) |
| `V2RAY_SERVER_IP` | The IP address or hostname of the target server. | **Required** |

> [!NOTE]
> The target port is currently hardcoded to `:80` in the code. Ensure your target server is listening on port 80 or modify `main.go` if needed.

## 🛠️ Installation & Usage

### Local Development

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/praveenkarunarathne/Google-Cloud-Run-Proxy.git
    cd Google-Cloud-Run-Proxy
    ```

2.  **Set environment variables:**
    ```bash
    export PORT=8080
    export V2RAY_SERVER_IP=1.2.3.4  # Replace with your target IP
    ```

3.  **Run the application:**
    ```bash
    go run main.go
    ```

### 🐳 Running with Docker

You can use the pre-built public image:

```bash
docker run -p 8080:8080 \
  -e PORT=8080 \
  -e V2RAY_SERVER_IP=1.2.3.4 \
  docker.io/praveenkarunarathne/google-cloud-run-proxy
```

Alternatively, to build from source:

1.  **Build the Docker image:**
    ```bash
    docker build -t cloud-run-proxy .
    ```

2.  **Run the container:**
    ```bash
    docker run -p 8080:8080 -e PORT=8080 -e V2RAY_SERVER_IP=1.2.3.4 cloud-run-proxy
    ```

## ☁️ Deploy to Google Cloud Run

Deploy directly using the public image:

```bash
gcloud run deploy cloud-run-proxy \
  --image docker.io/praveenkarunarathne/google-cloud-run-proxy \
  --platform managed \
  --region us-central1 \
  --allow-unauthenticated \
  --set-env-vars V2RAY_SERVER_IP=1.2.3.4
```

*Replace `1.2.3.4` with your target IP.*

## 📄 License

This project is licensed under the GPL-3.0 License - see the [LICENSE](LICENSE) file for details.
