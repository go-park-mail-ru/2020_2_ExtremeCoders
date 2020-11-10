package config

import "time"

const (
	DbUser       = "postgres"
	DbPassword   = "1538"
	DbDB         = "maila"
	Port         = ":8080"
	ReadTimeout  = 10 * time.Second
	WriteTimeout = 10 * time.Second
)

var AllowedOriginsCORS = []string{"http://localhost:3000", "http://127.0.0.1:3000", "http://95.163.209.195:3000"}
var AllowedHeadersCORS = []string{"Version", "Authorization", "Content-Type"}
var AllowedMethodsCORS = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
