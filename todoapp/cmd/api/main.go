package main

import (
	"github.com/go-chi/chi/v5"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	v1 "todoapp/cmd/api/handler/v1"
	"todoapp/pkg/repository/gormadp"
	"todoapp/service/todo"
)

func main() {
	r := chi.NewRouter()
	var (
		db  *gorm.DB
		err error
	)

	dsn := "root:my-secret-pw@tcp(127.0.0.1:3306)/todo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	repo := gormadp.New(db)

	if err = repo.Migrate(); err != nil {
		log.Fatalln(err.Error())
	}

	/*if err = db.Create(&dbmodels.Todo{
		Text: "çöpü at",
		Done: true,
	}).Error; err != nil {
		log.Fatalln(err.Error())
	}*/

	s := todo.New(repo)

	h := v1.New(s)

	r.Get("/todo", h.ListTodos)
	r.Post("/todo", h.AddTodo)
	r.Post("/todo/{id}", h.UpdateTodo)
	r.Delete("/todo/{id}", h.DeleteTodo)

	log.Print("server started successfully!")

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatalln(err.Error())
	}
}
