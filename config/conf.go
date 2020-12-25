package config

import "time"

const (
	DbUser = "postgres"
	//DbUser = "mark"
	DbPassword = "123456yhn"
	//DbPassword = "987654321"
	//DbPassword = "1538"
	//DbPassword   = "mark"
	DbDB = "maila"
	//DbDB         = "mail_db"
	Port         = ":8080"
	ReadTimeout  = 10 * time.Second
	WriteTimeout = 10 * time.Second
	AccessKey    = "vUEv3F69WEeN1D85oiiFgt"
	SecretKey    = "c5yvQ6ANBnxvU2txz6dQwY7rJjDvMmVxVEakNjgJfH4X"
	BucketName   = "maila"
	BucketID     = "mcs6132821991"
	Password     = "CherDan985fy1aasdf681553"
	Token        = ""
	MailDomain   = "mailer.ru.com"
)

var AllowedOriginsCORS = []string{
	"https://mailer.ru.com", "http://localhost:3000", "http://127.0.0.1:3000",
	"http://localhost", "http://127.0.0.1", "http://95.163.209.195:3000",
	"http://localhost:80", "http://127.0.0.1:80",
	"http://localhost", "http://127.0.0.1",
	"http://localhost:3000", "http://127.0.0.1:3000",
	"http://95.163.209.195:3000", "http://95.163.209.195",
	"http://95.163.209.195:3000", "http://95.163.209.195:80",
	"http://95.163.209.195", "https://95.163.209.195", "https://mailer.ru.com",
	"https://localhost:8080",
}
var AllowedHeadersCORS = []string{"Version", "Authorization", "Content-Type", "csrf_token"}
var AllowedMethodsCORS = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}

const (
	SizeSID  = 32
	CsrfSize = 32
)

var SidRunes = "1234567890_qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"
var CsrfRunes = "1234567890_qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"
