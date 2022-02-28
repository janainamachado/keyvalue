package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

type Vault struct {
	Id         int    `json:"id"`
	VaultKey   string `json:"key"`
	VaultValue string `json:"value"`
}

var vaults []Vault

func getVaultItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var vault Vault

	params := mux.Vars(r)
	key := params["key"]

	Db.First(&vault, "vault_key = ?", key)

	json.NewEncoder(w).Encode(vault)
}

func getAllItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var vaults Vault
	result := Db.Find(&vaults)

	json.NewEncoder(w).Encode(result)
}

func deleteVaultItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	key := params["key"]

	Db.Delete(&Vault{}, "vault_key = ?", key)

}

func upsertKeyValue(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var requestVault Vault
	json.NewDecoder(r.Body).Decode(&requestVault)

	var existingVault Vault

	Db.First(&existingVault, "vault_key = ?", requestVault.VaultKey)

	if (Vault{}) == existingVault {
		Db.Select("VaultKey", "VaultValue").Create(&requestVault)

		json.NewEncoder(w).Encode(requestVault)
	} else {
		existingVault.VaultKey = requestVault.VaultKey
		existingVault.VaultValue = requestVault.VaultValue

		Db.Save(&existingVault)

		json.NewEncoder(w).Encode(existingVault)
	}
}

func setHeaders(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "content-type")
		w.Header().Set("Access-Control-Request-Method", "*")
		w.Header().Set("Content-Type", "application/json")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}

func main() {
	var err error
	Db, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/keyvalue")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DB Connected!")

	r := mux.NewRouter()

	r.HandleFunc("/vault", getAllItems).Methods("GET")
	r.HandleFunc("/vault/{key}", getVaultItem).Methods("GET")
	r.HandleFunc("/vault", upsertKeyValue).Methods("POST")
	r.HandleFunc("/vault/{key}", deleteVaultItem).Methods("DELETE")

	fmt.Printf("Starting server at port 5000")
	log.Fatal(http.ListenAndServe(":5000", setHeaders(r)))
}
