package models

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Coin struct {
	Id uint
	Name string 
	CreatedAt time.Time
	UpdatedAt time.Time
}


