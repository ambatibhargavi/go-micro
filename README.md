To-Do List Application with Docker and Terraform
This repository contains a simple To-Do List application implemented in Go and containerized using Docker. The application is deployed and managed using Terraform for automation and infrastructure provisioning.

🚀 Features
A lightweight Go-based To-Do List web application.
Fully containerized using Docker.
Automated deployment using Terraform.
Accessible via http://localhost:5000.
🛠️ Technologies Used
Go (Golang): Backend logic for the To-Do List app.
Docker: Containerization of the application.
Terraform: Infrastructure as Code (IaC) to manage Docker resources.
📂 Project Structure
<img width="735" alt="Screenshot 2024-12-28 at 13 04 48" src="https://github.com/user-attachments/assets/7bd59b9e-4759-45c1-ab56-2b1df5e61e47" />


├── main.go          # Go application source code
├── go.mod           # Go module dependencies
├── Dockerfile       # Dockerfile for containerizing the app
├── main.tf          # Terraform configuration
├── README.md        # Project documentation
🖥️ Application Preview
A simple To-Do List app where users can add, view, and manage tasks.

## Build and Run with Terraform
Initialize Terraform:
Access the application at http://localhost:5000.

🐳 Docker Details
Dockerfile
The Dockerfile is used to build the Go application into a Docker image
terraform destroy

<img width="1296" alt="Screenshot 2024-12-28 at 12 04 32" src="https://github.com/user-attachments/assets/fbf00dc4-27b6-4a71-b2c9-d679616818e1"/>
<img width="348" alt="Screenshot 2024-12-28 at 12 04 17" src="https://github.com/user-attachments/assets/e5d72a32-d0ea-493c-bd67-e1863642ce4d" />
<img width="630" alt="Screenshot 2024-12-28 at 12 03 43" src="https://github.com/user-attachments/assets/429d43e3-a464-486b-a2bc-77c7b0c7dc31" />
<img width="611" alt="Screenshot 2024-12-28 at 09 30 44" src="https://github.com/user-attachments/assets/82541b48-04df-460f-9070-cd3b7cafcf65" />
<img width="917" alt="Screenshot 2024-12-28 at 09 30 28" src="https://github.com/user-attachments/assets/bc67ba59-26d9-4716-99ce-6a30a516d5df" />
