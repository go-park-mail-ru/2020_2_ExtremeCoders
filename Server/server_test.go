package Server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	go Start()
	client := http.Client{}
	message := map[string]interface{}{
		"Name":     "name",
		"Surname":  "surname",
		"Email":    "123@mail.ru",
		"Img":      "",
		"Password": "123",
	}
	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := client.Post("http://127.0.0.1:8080/signup", "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatalf(err.Error())
	}

	dec := json.NewDecoder(resp.Body)
	dec.DisallowUnknownFields()
	var ans Answer
	var exp Answer
	exp.Code = 200
	exp.Description = "ok"
	err = dec.Decode(&ans)
	if err != nil {
		log.Fatalf(err.Error())
		return
	}
	fmt.Println("response", ans)
	if ans.Code != exp.Code || ans.Description != exp.Description {
		log.Fatalf("Bad Request")
	}

	message = map[string]interface{}{
		"email":    "123@mail.ru",
		"password": "123",
	}

	bytesRepresentation, err = json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("RESPONSE")
	resp, err = client.Post("http://127.0.0.1:8080/signin", "application/json", bytes.NewBuffer(bytesRepresentation))

	dec = json.NewDecoder(resp.Body)

	dec.DisallowUnknownFields()
	exp.Code = 200
	exp.Description = "ok"
	var ans2 Answer
	err = dec.Decode(&ans2)
	if err != nil {
		log.Fatalf(err.Error())
		return
	}
	fmt.Println("RESPONSE", ans2)
	if ans2.Code != exp.Code || ans2.Description != exp.Description {
		log.Fatalf("Bad Request")
	}
	resp, err = client.Get("http://127.0.0.1:8080/signin")

	dec = json.NewDecoder(resp.Body)

	dec.DisallowUnknownFields()
	exp.Code = 400
	exp.Description = "Do not require request's method, expected POST"
	var ans3 Answer
	err = dec.Decode(&ans3)
	if err != nil {
		log.Fatalf(err.Error())
		return
	}
	fmt.Println("RESPONSE", ans3)
	if ans3.Code != exp.Code || ans3.Description != exp.Description {
		log.Fatalf("Bad Request")
	}
	resp, err = client.Get("http://127.0.0.1:8080/profile")

	dec = json.NewDecoder(resp.Body)

	dec.DisallowUnknownFields()
	exp.Code = 401
	exp.Description = "not authorized user"
	var ans4 Answer
	err = dec.Decode(&ans4)
	if err != nil {
		log.Fatalf(err.Error())
		return
	}
	fmt.Println("RESPONSE", ans4)
	if ans4.Code != exp.Code || ans4.Description != exp.Description {
		log.Fatalf("Bad Request")
	}
}

func TestServer2(t *testing.T) {
	go Start()
	client := http.Client{}
	message := map[string]interface{}{
		"Email":    "198233@mail.ru",
		"Password": "12345",
	}

	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := client.Post("http://127.0.0.1:8080/signin", "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatalf(err.Error())
	}

	dec := json.NewDecoder(resp.Body)
	dec.DisallowUnknownFields()
	var ans Answer
	var exp Answer
	exp.Code = 404
	exp.Description = "Do not find this user in db"
	err = dec.Decode(&ans)
	if err != nil {
		log.Fatalf(err.Error())
		return
	}
	fmt.Println("response", ans)
	if ans.Code != exp.Code || ans.Description != exp.Description {
		log.Fatalf("Bad Request")
	}
}

func TestServer3(t *testing.T) {
	go Start()
	client := http.Client{}
	message := map[string]interface{}{
		"Name":     "name",
		"Surname":  "surname",
		"Email":    "123@mail.ru",
		"Img":      "",
		"Password": "123",
	}

	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := client.Post("http://127.0.0.1:8080/signup", "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatalf(err.Error())
	}

	dec := json.NewDecoder(resp.Body)
	dec.DisallowUnknownFields()
	var ans Answer
	var exp Answer
	exp.Code = 401
	exp.Description = "This Email has already exists"
	err = dec.Decode(&ans)
	if err != nil {
		log.Fatalf(err.Error())
		return
	}
	fmt.Println("response", ans)
	if ans.Code != exp.Code || ans.Description != exp.Description {
		log.Fatalf("Bad Request")
	}

	message = map[string]interface{}{
		"Name":    "name",
		"Surname": "surname",
	}
	bytesRepresentation, err = json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}
	resp, err = client.Post("http://127.0.0.1:8080/profile", "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatalf(err.Error())
	}

	dec = json.NewDecoder(resp.Body)
	dec.DisallowUnknownFields()

	exp.Code = 401
	exp.Description = "not authorized user"
	err = dec.Decode(&ans)
	if err != nil {
		log.Fatalf(err.Error())
		return
	}
	fmt.Println("response", ans, exp)
	if ans.Code != exp.Code || ans.Description != exp.Description {
		log.Fatalf("Bad Request")
	}
}
