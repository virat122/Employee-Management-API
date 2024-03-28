package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateEmployee(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var emp Employee

	json.NewDecoder(r.Body).Decode(&emp)
	Database.Create(&emp)
	json.NewEncoder(w).Encode(emp)

}

func GetEmplyees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var emps []Employee
	Database.Find(&emps)
	json.NewEncoder(w).Encode(emps)

}

// niche kam baki  h       abhi

func GetEmplyeeByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applicatio/json")

	var employee Employee
	Database.First(&employee, mux.Vars(r)["eid"])
	json.NewEncoder(w).Encode(&employee)

}
func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applicatio/json")

	var employee Employee
	Database.First(&employee, mux.Vars(r)["eid"])
	json.NewDecoder(r.Body).Decode(&employee)
	Database.Save(&employee)

	json.NewEncoder(w).Encode(employee)
}
func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var emp Employee

	Database.Delete(&emp, mux.Vars(r)["eid"])
	json.NewEncoder(w).Encode("employee is deleted !! ")

}
