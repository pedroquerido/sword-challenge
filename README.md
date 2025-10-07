# sword-challenge
Sword Challenge

## Modules
- tasks-service - a RESTful API that provides actions on tasks
- notifications-agent - a basic amqp consumer 
- mock-auth-service - a mock auth server that relies on a config file to mock user authentication

## Architecture
This challenge's architecture is sort of a proof of concept on how a distributed, microservices architecture could be designed. It uses *traefik* (containo.us/traefik/) as an API gateway with the mock-auth-server validating authentication on user requests and the tasks-service relying on forwarded headers.

**Disclaimer:** As all POCs, it is not production ready and should not be used at that level.

## Tasks-Service
A basic CRUD that performs actions on the **Task** resource. Uses AES Encryption, a MySQL database and a RabbitMQ Publisher. It's endpoints are:

 - GET /tasks 
 - POST /tasks
 - GET /tasks/{taskId}
 - PATCH /tasks/{taskId}
 - DELETE /tasks/{taskId}
 - GET /users/{userId}/tasks

**Note:** With *traefik* it's base path is /task, so append that to all requests if testing integrated.

## Mock-Auth-Server
A mock authentication server that uses the contents of the *users.yaml* file and the *x-api-key* header to manage authorizations. 

## Notifications-Agent
A amqp-based consumer that listens on events published (for now, just the one task created event) on his queue and does actions based on them.

## Contents
In addition to the modules described, the repo contains two docker-compose files: one for setting up the application and the other for running the tests. It also contains a *dynamic.toml* file for  *traefik* configurations.

## Future Developments
For a more "production ready" application, some key points to build upon:
 - Add pagination to list requests
 - Improve logging and build upon error handling
 - Improve DB migrations
 - Setup proper documentation
 - Refine MySQL and RabbitMQ connections
 - Create deployment files for a possible Kubernetes integration

