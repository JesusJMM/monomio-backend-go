package main

import (
	"database/sql"
	"fmt"

	"github.com/JesusJMM/monomio/api"
	"github.com/JesusJMM/monomio/postgres"
	_ "github.com/lib/pq"
)

func main(){
  dbConn, err := sql.Open("postgres", "user=monomioapp dbname=monomioapp sslmode=disable") 
  if err != nil {
    fmt.Println(err)
    panic(err)
  }
  db := postgres.New(dbConn)

  r := api.NewHandler(*db)
  r.Run()
}
