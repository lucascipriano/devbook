package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repository"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func CraeteUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}
	var user models.User
	if erro = json.Unmarshal(bodyRequest, &user); erro != nil {
		log.Fatal(erro)
		return
	}

	db, erro := database.ConnectDB()
	if erro != nil {
		log.Fatal(erro)
		return
	}
	repositorio := repository.NewRepositoryUsers(db)
	userID, erro := repositorio.Create(user)
	if erro != nil {
		log.Fatal(erro)
		return
	}
	w.Write([]byte(fmt.Sprintf("id inserido: %d", userID)))

}
func GrabAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os usuário"))
}
func GrabUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um usuário"))
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuário"))
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando usuário"))
}
