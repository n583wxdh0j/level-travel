package controllers

import (
	"html/template"
	"level-travel/database"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	libs, err := database.GetLibraries()
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	t, err := template.New("webpage").ParseFiles("index.html")
	if err != nil {
		http.Error(w, "Parse files error", http.StatusInternalServerError)
		return
	}
	if err := t.Execute(w, libs); err != nil {
		http.Error(w, "Parse files error", http.StatusInternalServerError)
		return
	}
}
