package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
type UserController struct{}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Aquí se insertaría el usuario en la base de datos
	w.WriteHeader(http.StatusCreated)
}

func (uc *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	// Obtener el ID del usuario de los parámetros de la URL
	userID := chi.URLParam(r, "userID")
	// Aquí se realizaría la consulta a la base de datos para obtener el usuario correspondiente al ID
	user := User{ID: userID, Name: "John Doe", Email: "john.doe@example.com"}
	if user.ID == "" {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (uc *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Obtener el ID del usuario de los parámetros de la URL
	userID := chi.URLParam(r, "userID")
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Aquí se actualizaría el usuario en la base de datos
	w.WriteHeader(http.StatusOK)
}

func (uc *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Obtener el ID del usuario de los parámetros de la URL
	userID := chi.URLParam(r, "userID")
	// Aquí se eliminaría el usuario de la base de datos
	w.WriteHeader(http.StatusOK)
}

func main() {
	router := chi.NewRouter()
	userController := &UserController{}
	router.Post("/users", userController.CreateUser)
	router.Get("/users/{userID}", userController.GetUser)
	router.Put("/users/{userID}", userController.UpdateUser)
	router.Delete("/users/{userID}", userController.DeleteUser)
	http.ListenAndServe(":8080", router)
}
