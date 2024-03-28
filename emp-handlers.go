package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateEmployee(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return

	}

	w.Header().Set("Content-Type", "application/json")
	var emp Employee

	json.NewDecoder(r.Body).Decode(&emp)
	Database.Create(&emp)
	json.NewEncoder(w).Encode(emp)

}

func GetEmplyees(w http.ResponseWriter, r *http.Request) {
	// check if request type
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	var emps []Employee
	Database.Find(&emps)
	json.NewEncoder(w).Encode(emps)

}

func GetEmplyeeByID(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "applicatio/json")

	var employee Employee
	Database.First(&employee, mux.Vars(r)["eid"])
	json.NewEncoder(w).Encode(&employee)

}
func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "applicatio/json")

	var employee Employee
	Database.First(&employee, mux.Vars(r)["eid"])
	json.NewDecoder(r.Body).Decode(&employee)
	Database.Save(&employee)

	json.NewEncoder(w).Encode(employee)
}
func DeleteEmployee(w http.ResponseWriter, r *http.Request) {

	if r.Method != "DELETE" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var emp Employee

	Database.Delete(&emp, mux.Vars(r)["eid"])
	json.NewEncoder(w).Encode("employee is deleted !! ")

}
