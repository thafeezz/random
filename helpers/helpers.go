package helpers

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
)

/*
Handler struct
route handlers are "member" functions of the Handler struct
*/
type Handler struct {
	db *sql.DB
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{db: db} // dependency injection
}

func (h *Handler) GetEmails(c *gin.Context) {
	rows, err := h.db.Query("SELECT username FROM test_schema.test_table")
	if err != nil {

	}
}

func (h *Handler) AddEmail(c *gin.Context) {
}

func ConnectDb() (*sql.DB, error) {
	db, err := sql.Open("postgres", "jdbc:postgresql://localhost:5432/postgres")
	fmt.Println("Connected to the database")

	return db, err
}
