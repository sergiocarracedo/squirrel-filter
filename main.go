package main

import (
	"fmt"
	"github.com/Masterminds/squirrel"
	"test/squirrel_filter"
	"time"
)

type SubGroup struct {
	Country string `qg:"required"`
}
type Filter struct {
	Name        string    `sqfilter:"=, required" json:"name" db:"name"`
	Surname     string    `sqfilter:"like" json:"surname" db:"surname"`
	City        string    `sqfilter:"contains, db = city" json:"city"`
	BirthdayGte time.Time `sqfilter:">,db=birthday" json:"birthday_gte"`
	BirthdayLte time.Time `sqfilter:"<,db=birthday" json:"birthday_lte"`
}

func main() {

	filter := Filter{
		Name:        "Sergio",
		Surname:     "Carracedo",
		City:        "Vigo",
		BirthdayGte: time.Now(),
		BirthdayLte: time.Now(),
	}

	builder := squirrel.Select("*").From("table")
	filterConditions, _ := squirrel_filter.GetConditions(filter)
	builder = builder.Where(filterConditions)

	query, args, err := builder.ToSql()
	fmt.Println(query, args, err)

}
