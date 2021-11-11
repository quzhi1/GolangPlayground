package main

import (
	"fmt"
	"github.com/go-martini/martini"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var logger = log.New(os.Stdout, "logger: ", log.Lshortfile)

func main() {
	m := martini.Classic()

	// Simple response
	m.Get("/default", func() string {
		return "Hello world!"
	})

	// Simple response with status code
	m.Get("/status-code", func() (int, string) {
		return 418, "i'm a teapot"
	})

	// You can inject http.Request and ResponseWriter (or either)
	m.Get("/service-injection", func(res http.ResponseWriter, req *http.Request) { // res and req are injected by Martini
		res.WriteHeader(200)
		_, _ = fmt.Fprint(res, "This makes Martini completely compatible with golang's http.HandlerFunc interface.")
	})

	// Server logging
	m.Get("/secret-logging", func() string {
		logger.Print("This is a secret logging")
		return "See server log"
	})

	// Three ways of define path parameter
	m.Get("/hello-named/:name", func(params martini.Params) string {
		return "Hello " + params["name"]
	})
	m.Get("/hello-wild/**", func(params martini.Params) string {
		return "Hello " + params["_1"]
	})
	m.Get("/hello-regex/(?P<name>[a-zA-Z]+)", func(params martini.Params) string {
		return fmt.Sprintf("Hello %s", params["name"])
	})

	// Add middleware
	m.Get("/authorize", authorize, func() string {
		return "Authorized"
	})

	// Sub routing (you can also insert middleware)
	m.Group("/v1", func(r martini.Router) {
		r.Get("/get", func() string {
			return "Response for /v1/get"
		})
		r.Post("/post", func(req *http.Request) string {
			buf := new(strings.Builder)
			_, _ = io.Copy(buf, req.Body)
			return "Response for /v1/get with request body: " + buf.String()
		})
	}, authorize)

	// You can map a service
	m.Map(GetDatabase())
	m.Get("/data/:name", func(params martini.Params, database *MyDatabase) string {
		role, _ := database.Data[params["name"]]
		return role
	})

	// Inject global middleware
	m.Use(GlobalMiddleware)

	// Inject sandwich middleware
	m.Use(SandwichMiddleware)

	m.RunOnAddr(":8080")
	//m.Run() // Run on 3000
}

// Middlewares

func authorize() {
	logger.Print("Authorization done")
}

func GlobalMiddleware(_res http.ResponseWriter, req *http.Request) {
	logger.Print("Processing " + req.URL.Path)
}

func SandwichMiddleware(c martini.Context) {
	log.Println("before a request")
	c.Next()
	log.Println("after a request")
}

// Database helper

type MyDatabase struct {
	Data map[string]string
}

func GetDatabase() *MyDatabase {
	data := map[string]string{
		"zhi":     "Husband",
		"fangzhu": "Wife",
	}
	return &MyDatabase{data}
}

// If you want to live-load this, run
// gin --appPort 8080 --port 80 run main.go
// appPort is the configured port for original app
// port is the gin proxy
// Please use gin proxy port to debug
