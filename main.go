package main

import (
	
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)
func HelloWorld(w http.ResponseWriter, r *http.Request){
Time:=time.Now()
 fmt.Fprint(w,"helloworld kk",Time)
 
 
}
func MyRouter(){
	
	mux:=mux.NewRouter().StrictSlash(true)
	mux.HandleFunc("/",HelloWorld).Methods("GET")
	mux.HandleFunc("/users",Users).Methods("GET")
	mux.HandleFunc("/user/{name}/{email}",NewUser).Methods("POST")
	mux.HandleFunc("/user/{name}/{email}",ReplaceUser).Methods("PUT")
	mux.HandleFunc("/delete",DeleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8083",mux))

}
func main(){
	fmt.Println("Running")
	InitialMigration()
	MyRouter()
}