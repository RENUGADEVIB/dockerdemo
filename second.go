package main

import (

	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	//"math/rand"
	"strconv"
	"fmt"
)

type emp struct{
	Name string `json:"name"`
	Id int 	`json:"id"`
}
var abc []emp 

func getemp(w http.ResponseWriter, r* http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(abc)
}

func postemp(w http.ResponseWriter, r * http.Request){
	w.Header().Set("Content-type","application/json")
	var a emp
	_=json.NewDecoder(r.Body).Decode(&a)
	abc=append(abc,a)
}
func putemp(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-type","application/json")
	param:=mux.Vars(r)
	id:=param["id"]
	pa, err :=strconv.Atoi(id)
	if err ==nil{

	for index,item := range abc{
		if item.Id==pa{
		//abc=append(abc[:index],abc[index+1:]...)
			
		var a emp
		_= json.NewDecoder(r.Body).Decode(&a)
		fmt.Println(a.Name)
		a.Id=pa
		//abc=append(abc,a)
		abc[index].Name=a.Name
		//json.NewEncoder(w).Encode(&a)
		//return
		}
	}
	json.NewEncoder(w).Encode(abc)
}
}
func main(){
	r:= mux.NewRouter()
	abc=append(abc,emp{Name:"renu",Id:2})
	r.HandleFunc("/emp",getemp).Methods("GET")
	r.HandleFunc("/emp_post",postemp).Methods("POST")
	r.HandleFunc("/putemp/{id}",putemp).Methods("PUT")
	http.ListenAndServe(":8030", r)
}