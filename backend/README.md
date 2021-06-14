Start with 

`go run main.go`

Available Endpoints

`localhost:8080/choices` Get
`localhost:8080/play` Post

With docker 

`docker build -t backend .`

This will create container with name backend. 

List all running containers 

`docker images`

To run `backend` container 

`docker run -p 8080:8080 backend` 

This will run API endpoint in port 8080 of host machine.