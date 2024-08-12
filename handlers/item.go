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

func Prediction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var item models.Item
	json.NewDecoder(r.Body).Decode(&item)
	item.Description = strings.ReplaceAll(utils.Predict(item.Url),"\n", "<br>");

	json.NewEncoder(w).Encode(item)
}


