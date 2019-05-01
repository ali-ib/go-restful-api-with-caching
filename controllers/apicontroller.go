package controllers

import (
	"encoding/json"
	"fmt"
	"gomongo/configs"
	"gomongo/dao"
	"gomongo/models"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

var peopleArr []models.Person

// Fetching data from database into cache
func fetchData() {
	var err error
	peopleArr, err = dao.Find(bson.M{})
	if err != nil {
		log.Fatal(err)
	}
}

// Scheduling the fetching process every CACHETIME seconds
func SyncCache() {
	ticker := time.NewTicker(configs.Configs.CACHETIME * time.Second)
	for ; true; <-ticker.C {
		fetchData()
		fmt.Println(peopleArr)
	}
}

// API handler for getting a person by his ID
func GetPersonById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	for _, person := range peopleArr {
		if person.ID.Hex() == id {
			json.NewEncoder(w).Encode(bson.M{"result": "success", "data": person})
			return
		}
	}
	json.NewEncoder(w).Encode(bson.M{"result": "failure", "data": nil})

}

// API handler for getting all people
func GetAllPeople(w http.ResponseWriter, r *http.Request) {
	people, err := dao.Find(bson.M{})
	if err != nil {
		json.NewEncoder(w).Encode(bson.M{"result": "failure", "error": err.Error()})
		return
	}
	json.NewEncoder(w).Encode(bson.M{"result": "success", "data": people})
}

// API handler for creating a new person
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var person models.Person
	person.Name = r.FormValue("name")
	age, _ := strconv.Atoi(r.FormValue("age"))
	person.Age = age
	id, err := dao.InsertOne(person)
	if err != nil {
		json.NewEncoder(w).Encode(bson.M{"result": "failure", "error": err.Error()})
		return
	}
	json.NewEncoder(w).Encode(bson.M{"result": "success", "createdID": id})
	fetchData()
}

// API handler for deleting a person by his ID
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	deletedCount, err := dao.DeleteOne(r.FormValue("id"))
	if err != nil {
		json.NewEncoder(w).Encode(bson.M{"result": "failure", "error": err.Error()})
		return
	}
	json.NewEncoder(w).Encode(bson.M{"result": "success", "deletedCount": deletedCount})
	fetchData()
}

// API handler for updating a person by his ID
func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	personID := r.FormValue("id")
	var person models.Person
	person.Name = r.FormValue("name")
	age, _ := strconv.Atoi(r.FormValue("age"))
	person.Age = age
	modifiedCount, err := dao.UpdateOne(person, personID)
	if err != nil {
		json.NewEncoder(w).Encode(bson.M{"result": "failure", "error": err.Error()})
		return
	}
	json.NewEncoder(w).Encode(bson.M{"result": "success", "modifiedCount": modifiedCount})
	fetchData()
}
