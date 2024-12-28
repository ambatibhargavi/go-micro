terraform {
  required_providers {
    docker = {
      source  = "kreuzwerker/docker"
      version = "~> 2.22.0"
    }
  }
}

provider "docker" {
  host = "unix:///var/run/docker.sock"
}

# Build the Docker image
resource "docker_image" "go_app_image" {
  name = "my-go-app:latest"
  build {
    path = "${path.module}" # The directory containing the Dockerfile
  }
}

# Run the Docker container
resource "docker_container" "go_app_container" {
  name  = "my-go-app-container"
  image = docker_image.go_app_image.name

  ports {
    internal = 5000 # Port inside the container
    external = 5000 # Port exposed on the host
  }
}
