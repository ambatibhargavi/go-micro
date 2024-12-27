To-Do List
A simple and efficient web application to manage your tasks and keep track of your daily to-dos. The application allows users to add, update, and remove tasks easily. You can also mark tasks as complete or pending based on their progress.

Features
Add Tasks: Add new tasks to your to-do list with a title and description.
Edit Tasks: Update the details of existing tasks.
Mark as Completed: Track the progress by marking tasks as completed.
Delete Tasks: Remove tasks when they are no longer needed.
Filter Tasks: Filter tasks by status (completed/pending).
Responsive UI: Works well on all screen sizes, including mobile devices.
Technologies Used
Frontend: HTML, CSS, JavaScript
Backend: Go (Golang)
Deployment: Docker (for containerization) & Kubernetes (for deployment)
Installation
Prerequisites
Make sure you have the following installed:

Go (>=1.18)
Docker (for containerization)
Kubernetes (for deployment)
Steps to Run the Application
Clone the repository:


git clone https://github.com/ambatibhargavi/todo-list.git
cd todo-list
Install dependencies:

Go doesn’t have a formal package manager, but you can use go mod to manage dependencies. Ensure the Go modules are set up:

bash
Copy code
go mod tidy
Run the application locally:
go run main.go
Access the application:

Open your browser and go to http://localhost:5000/
Docker Setup
To run the application in a container:

Create a Dockerfile :
Build the Docker image:

docker build -t todo-list .
Run the Docker container:


docker run -p 8080:8080 todo-list
Access the app:

Open your browser and go to http://localhost:8080/.

Deployment on Kubernetes
Build and push the Docker image to a container registry (e.g., Docker Hub).

Create Kubernetes deployment and service files (deployment.yaml, service.yaml).

Apply the Kubernetes configurations:

kubectl apply -f deployment.yaml
kubectl apply -f service.yaml
Access the application:

Use the appropriate method (e.g., kubectl port-forward) to expose the app.

Usage
After launching the app, users can:
Add new tasks by entering a task name and description.
Mark tasks as completed by clicking on a checkbox.
Delete tasks by clicking the delete icon.
Filter tasks based on their completion status (Completed/Pending).
