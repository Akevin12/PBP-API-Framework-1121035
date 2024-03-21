package main

import (
	"github.com/go-martini/martini"
	_ "github.com/go-sql-driver/mysql"
	"github.com/modul2/controllers"
)

func main() {
	// Membuat instance Martini
	m := martini.Classic()

	// Routes
	r := martini.NewRouter()
	r.Get("/users/:ID", controllers.GetUser)
	r.Get("/users", controllers.GetAllUser)
	r.Post("/users", controllers.CreateUser)
	r.Put("/users/:ID", controllers.UpdateUser)
	r.Delete("/users/:ID", controllers.DeleteUser)

	// Hubungkan router ke aplikasi Martini
	m.MapTo(r, (*martini.Routes)(nil))
	m.Action(r.Handle)

	// Menjalankan aplikasi pada port 8000
	m.RunOnAddr(":8888")
}
