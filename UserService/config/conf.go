package config

import "time"

const (
<<<<<<< HEAD
	DbUser = "postgres"
	//DbPassword   = "123456yhn"
	DbPassword   = "1538"
=======
	DbUser     = "postgres"
	DbPassword = "123456yhn"
	//DbPassword   = "1538"
>>>>>>> eac4bad70660374b439c87a65683a0fb275c4b7d
	DbDB         = "maila"
	Port         = ":8080"
	ReadTimeout  = 10 * time.Second
	WriteTimeout = 10 * time.Second
)

var AllowedOriginsCORS = []string{"http://localhost:3000", "http://127.0.0.1:3000", "http://95.163.209.195:3000"}
var AllowedHeadersCORS = []string{"Version", "Authorization", "Content-Type", "csrf_token"}
var AllowedMethodsCORS = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}

const (
	SizeSID  = 32
	CsrfSize = 32
)

var SidRunes = "1234567890_qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"
var CsrfRunes = "1234567890_qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"
