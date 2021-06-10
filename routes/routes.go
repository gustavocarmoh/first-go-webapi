package routes

import (
	"net/http"

	produtos "github.com/gustavocarmoh/controller"
)

func Routes () {
	http.HandleFunc("/", produtos.Index)
	http.HandleFunc("/new", produtos.New)
	http.HandleFunc("/insert", produtos.Insert)
	http.HandleFunc("/delete", produtos.Delete)
	http.HandleFunc("/edit", produtos.Edit)
	http.HandleFunc("/update", produtos.Update)
}
