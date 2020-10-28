package Delivery

import (
	"CleanArch/app/Models"
	"fmt"
	"net/http"
)

func (yaFood *Delivery)SendLetter(w http.ResponseWriter, r *http.Request){
	fmt.Print("Send Letter: ")
	fmt.Print("\n\n")
	if r.Method != http.MethodPost {
		return
	}
	var user Models.User
	user.Name = r.FormValue("name")
	user.Surname = r.FormValue("surname")
	user.Email = r.FormValue("email")
	user.Password = r.FormValue("password1")
}