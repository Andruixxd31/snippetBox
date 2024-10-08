package main

import (
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	files := []string{
		"../../ui/html/pages/home.tmpl.html",
		"../../ui/html/partials/nav.tmpl.html",
		"../../ui/html/base.tmpl.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		app.errorLog.Println(err.Error())
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.Write([]byte(`{"status:"Success", "message":"snippet obtained"}`))
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", http.MethodPost)
		// w.WriteHeader(405)
		// w.Write([]byte("Method Not Allowed"))
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.Write([]byte(`{"status:"Success", "message":"Create a new snippet"}`))
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(201) // WriteHeader can only be called once in a method
	w.Header().Set("Content-type", "application/json")
	w.Write([]byte(`{"status:"Success", "message":"snippet created"}`))
}
