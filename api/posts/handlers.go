package posts

import (
	// "net/http"

	// apiDT "github.com/JesusJMM/monomio/api/apiDataTypes"
	"github.com/JesusJMM/monomio/postgres"
	// "github.com/gin-gonic/gin"
)

type PostsHandler struct {
  db *postgres.Queries
}

func New(db *postgres.Queries) PostsHandler{
  return PostsHandler{
    db: db,
  }
}
