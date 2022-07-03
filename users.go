package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)
var db *gorm.DB 
var Error error

type User struct{
	gorm.Model
	Name string
	Email string
}

func InitialMigration(){
	db,err:=gorm.Open("sqlite3","test_db")
	if err!= nil{
		fmt.Println("failed to connect to the database")
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&User{})
}

func Users(w http.ResponseWriter, r *http.Request) {
	db,err:=gorm.Open("sqlite3","test_db")
	if err!=nil{
		fmt.Println("failed to connect to the db")
	} 
	fmt.Println(db.HasTable(&User{}))
	defer db.Close()
	var user []User
	db.Find(&user)
	json.NewEncoder(w).Encode(user)
	fmt.Println(user)
	fmt.Fprintf(w,"succesfully getting all users")
}

func NewUser(w http.ResponseWriter, r *http.Request){
	db,err:=gorm.Open("sqlite3","test_db")
		if err!=nil{
			panic("failed to connect to the database")
		}
		defer db.Close()
		vars:=mux.Vars(r)
		name:=vars["name"]
		email:=vars["email"]
		db.AutoMigrate(&User{})
		db.Create(&User{Name: name,Email: email})
		fmt.Fprintf(w,"creating new user success!")
	}
	
func ReplaceUser(w http.ResponseWriter, r *http.Request){
	db,err:=gorm.Open("sqlite3","test_db")
	if err!=nil{
		panic("failed to connect to the database")
	}
	defer db.Close()
	vars:=mux.Vars(r)
	name:=vars["name"]
	email:=vars["email"]

	var user User
	db.Where("name=?",name).Find(&user)
	user.Email=email

	db.Save(&user)
	db.AutoMigrate(&user)
    
	fmt.Fprintf(w,"Updating users")
}
func DeleteUser(w http.ResponseWriter, r *http.Request){
	db,err:=gorm.Open("sqlite3","test_db")
		if err!=nil{
			panic("failed to connect to the database")
		}
		defer db.Close()

	vars:=mux.Vars(r)
	name:=vars["name"]
	var user User

	db.Where("user=?",name).Find(&user)
	db.Delete(&user)
	db.AutoMigrate(&user)
	fmt.Fprintf(w,"%v is deleted from the database ",name)
}
