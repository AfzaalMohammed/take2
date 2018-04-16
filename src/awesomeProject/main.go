package main

import (
	"io/ioutil"
	"net/http"
	"fmt"
)
type page struct {
Title string
Body[]byte
}

func (p *page) save () error{
f := p.Title + ".txt"
return ioutil.WriteFile(f, p.Body, 0600)
}
func load(title string) (*page, error) {
f := title + ".txt"
body, err := ioutil.ReadFile(f)
if err != nil {
return nil, err
}
return &page{Title: title, Body: body}, nil
}
func view(w http.ResponseWriter, r *http.Request){
title := r.URL.Path[len("/Test/"):]
p, _ := load(title)
fmt.Fprintf(w,"<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}
func main(){
p := &page{Title: "Test", Body:[]byte("BRING IT ON.... ")}
p.save()
http.HandleFunc("/Test/", view)
http.ListenAndServe(":8000", nil)

}
