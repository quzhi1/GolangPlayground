package main

// Run locally:
// gin --appPort 3004 --port 80 -i run main.go
// Build docker image:
// docker build -t docker-sync .
// Run docker image:
// docker run -d -p 80:80 -v $(pwd):/go/src/docker-sync docker-sync
func main() {
	route := initiateRoutes()

	route.Run(":3004")

}
