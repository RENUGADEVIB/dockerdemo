package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"github.com/globalsign/mgo"
	// "log"
	"fmt"
	"gopkg.in/mgo.v2/bson"

)

type persons struct{
	Name string
	Age int
}

var session,err= mgo.Dial("localhost:27017")
// if err != nil{
// 	log.Fatal(err)
// 	fmt.Println("error")
// }	
var c = session.DB("demo").C("democollection")

func getall(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	view := [] persons{}
	c.Find(bson.M{}).All(&view)
	json.NewEncoder(w).Encode(view)

}

func insertdata(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var data persons
	_=json.NewDecoder(r.Body).Decode(&data)
	c.Insert(bson.M{"name":data.Name,"age":data.Age})

}

func updatedata(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	param:=mux.Vars(r)
	Name_a:=param["name"]
	var a persons
	_=json.NewDecoder(r.Body).Decode(&a)
	fmt.Println(a)
	c.Update(bson.M{"name":Name_a},bson.M{"$set":bson.M{"name":a.Name,"age":a.Age}})
}

func delete(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	param:=mux.Vars(r)
	name_b:=param["name"]
	// var a persons 
	// _= json.NewDecoder(r.Body).Decode(&a)
	c.Remove(bson.M{"name":name_b})
}
func main(){
router:=mux.NewRouter()
router.HandleFunc("/get",getall).Methods("GET")
router.HandleFunc("/insert",insertdata).Methods("POST")
router.HandleFunc("/up/{name}",updatedata).Methods("PUT")
router.HandleFunc("/del/{name}",delete).Methods("DELETE")
http.ListenAndServe(":8050", router	)
}