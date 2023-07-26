package controller

import (
	"encoding/json"
	"net/http"
	"scratch-db-api/service"

	"github.com/graphql-go/graphql"
)

func GetAllRecords() {
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var requestBody map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		query, ok := requestBody["query"].(string)
		if !ok {
			http.Error(w, "Missing or invalid 'query' field in the request body", http.StatusBadRequest)
			return
		}

		result := graphql.Do(graphql.Params{
			Schema:        service.Query(),
			RequestString: query,
		})
		json.NewEncoder(w).Encode(result)
	})
}
