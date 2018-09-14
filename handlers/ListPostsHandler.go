package handlers

import (
	"database/sql"
	"net/http"
	"fmt"
)

type ListPostsHandler struct {
	Database *sql.DB
}


func (h *ListPostsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Some stuff!")
}
