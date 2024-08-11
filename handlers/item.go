package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"carbon/models"
	"github.com/gorilla/mux"
	"carbon/utils"
	"strings"
)

var items []models.Item

func GetItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func GetItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for _, item := range items {
		if item.ID != nil && *item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.Error(w, "Item not found", http.StatusNotFound)
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var item models.Item
	json.NewDecoder(r.Body).Decode(&item)
	id := len(items) + 1
	item.ID = &id;
	items = append(items, item)
	json.NewEncoder(w).Encode(item)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for index, item := range items {
		if item.ID != nil && *item.ID == id {
			items = append(items[:index], items[index+1:]...)
			var updatedItem models.Item
			json.NewDecoder(r.Body).Decode(&updatedItem)
			updatedItem.ID = &id
			items = append(items, updatedItem)
			json.NewEncoder(w).Encode(updatedItem)
			return
		}
	}
	http.Error(w, "Item not found", http.StatusNotFound)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for index, item := range items {
		if item.ID != nil && *item.ID == id {
			items = append(items[:index], items[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(items)
}

func Prediction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var item models.Item
	json.NewDecoder(r.Body).Decode(&item)
	item.Description = strings.ReplaceAll(utils.Predict(item.Url),"\n", "<br>");

	json.NewEncoder(w).Encode(item)
}


