package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"students-api/internal/types"
	"students-api/internal/utils/response"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.WriteJSON(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body")))
			return
		}

		if err != nil {
			response.WriteJSON(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		// validate request or request validation

		response.WriteJSON(w, http.StatusOK, map[string]string{"success": "OK"})
	}
}
