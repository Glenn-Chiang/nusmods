package handler

import (
	"encoding/json"
	"net/http"

	"github.com/nusmodifications/nusmods/website/api/optimiser/_models"
	"github.com/nusmodifications/nusmods/website/api/optimiser/_solver"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	// Allow CORS from all origins
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// only allow POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// get selected modules from request
	var optimiserRequest models.OptimiserRequest
	err := json.NewDecoder(r.Body).Decode(&optimiserRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	solver.Solve(w, optimiserRequest)

}
