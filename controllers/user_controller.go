package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-martini/martini"

	m "github.com/modul2/model"
)

// GetUser berdasarkan ID
func GetUser(params martini.Params, w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	user := m.Users{}
	err := db.QueryRow("SELECT * FROM users WHERE ID=?", params["ID"]).Scan(&user.ID, &user.Name, &user.Age, &user.Address, &user.UserType)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// Get All user
func GetAllUser(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []m.Users
	for rows.Next() {
		var user m.Users
		err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Address, &user.UserType)
		if err != nil {
			http.Error(w, "Failed to scan user data", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}
	json.NewEncoder(w).Encode(users)
}

// CreateUser
func CreateUser(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	name := r.FormValue("Name")
	age, _ := strconv.Atoi(r.Form.Get("Age"))
	address := r.FormValue("Address")
	userType, _ := strconv.Atoi(r.Form.Get("UserType"))

	result, err := db.Exec("INSERT INTO users (Name, Age, Address, UserType) VALUES (?, ?, ?, ?)", name, age, address, userType)
	if err != nil {
		http.Error(w, "Create Failed", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	id, _ := result.LastInsertId()
	user := m.Users{
		ID:       int(id),
		Name:     name,
		Age:      age,
		Address:  address,
		UserType: userType,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// UpdateUser
func UpdateUser(params martini.Params, w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	id, _ := strconv.Atoi(params["ID"])

	var user m.Users
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	result, err := db.Exec("UPDATE users SET Name=?, Age=?, Address=?, UserType=? WHERE ID=?", user.Name, user.Age, user.Address, user.UserType, id)
	if err != nil {
		http.Error(w, "Update Failed", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	user.ID = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// DeleteUser
func DeleteUser(params martini.Params, w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	id, _ := strconv.Atoi(params["ID"])

	result, err := db.Exec("DELETE FROM users WHERE ID=?", id)
	if err != nil {
		http.Error(w, "Delete Failed", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Delete Success"})
}
