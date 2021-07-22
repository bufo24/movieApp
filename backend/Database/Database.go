package Database

import (
	"fmt"
	"github.com/go-pg/pg/v10"
)

func init() {
	opt, err := pg.ParseURL("postgres://postgres:JessePostgres2003?!@192.168.0.11:5432/MovieApp?sslmode=verify-full")
	if err != nil {
	panic(err)
	}

	db := pg.Connect(opt)
	fmt.Println(db)
}
