package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lame/pkg/app"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

type Serv struct {
	s *gin.Engine
}

type StatusStruct struct {
	Status string `json:"status"`
	Data   string `json:"data"`
}

func ApiStatus(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	response := StatusStruct{
		"success",
		"API is running success",
	}

	c.JSON(http.StatusOK, response)
	return
}

func ReverseProxy(c *gin.Context) {
	remote, _ := url.Parse(fmt.Sprintf("%s", os.Getenv("FRONTENDLINK")))
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.Director = func(req *http.Request) {
		req.Header = c.Request.Header
		req.Host = remote.Host
		req.URL = c.Request.URL
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
	}

	proxy.ServeHTTP(c.Writer, c.Request)
}

func Run(s *gin.Engine) error {
	if err := Routes(s); err != nil {
		log.Printf("Server - there was an error calling Routes: %v", err)
		return err
	}

	fmt.Println(fmt.Sprintf(":%s", os.Getenv("PORT")))
	err := s.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
	if err != nil {
		log.Printf("Server - there was an error calling Run on router: %v", err)
		return err
	}

	return nil
}

func Routes(s *gin.Engine) error {
	s.NoRoute(ReverseProxy)
	s.GET("/api/v1/status/", ApiStatus)
	s.POST("/api/v1/request/", app.ContactUs)
	return nil
}

func main() {
	fmt.Println("Start the server")
	if err := run(); err != nil {
		log.Println(os.Stderr, fmt.Sprintf("This is the startup error: %s\n", err))
		os.Exit(1)
	}
}

func run() error {
	router := gin.Default()
	router.Use(app.CorsConfig())

	err := Run(router)
	if err != nil {
		return err
	}
	return nil
}
