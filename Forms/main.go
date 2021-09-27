package main

import (
	"net/http"
	"html/template"
)

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main(){
	http.HandleFunc("/",index)
	http.HandleFunc("/process",processor)
	http.ListenAndServe(":8080",nil)
}

func index(w http.ResponseWriter, r *http.Request){
	tpl.ExecuteTemplate(w,"index.html",nil)
}

func processor(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET" {
		http.Redirect(w,r,"/",http.StatusSeeOther)
		return
	}
	fname := r.FormValue("fname")
	lname := r.FormValue("lname")

	d := struct{
		First string
		Last string
	}{
		First: fname,
		Last: lname,
	}

	tpl.ExecuteTemplate(w,"processor.html",d)
}