# squirrel-filter: sql conditions from a struct
This package provides a layer over [squirrel](https://github.com/Masterminds/squirrel) (a sql query builder) to 
simplify filtering.

## Background
Working with items in databases is very common the need of filter items from a list using http query params to set 
filter's values.

Even using a query builder like Squirrel, create the query need a lot of "if"s to check if the value is present and 
add the filter if its necessary

## Objective 
This package allows you to generate squirrel's conditions for the query builder from an struct using tags, 
you could populate the struct values using the package `json` or populate and validate using [validator](https://github.com/go-playground/validator)


# Instalation
Use go get
```bach
go get github.com/sergiocarracedo/squirrel-filter
```
Then import the package in your code
```golang
import "github.com/sergiocarracedo/squirrel-filter
```


# Usage
Add the tag `sqFilter` to your struct tags 

Example:
```golang
type Filter struct {
	Name     string `sqFilter:"=" json:"name"`
	Surname  string `sqFilter:"like, required" json:"surname" db:"surname`
	City     string `sqFilter:"contains" json:"city"`
	BirthdayGte string `sqFilter:">, db=birthday" json:"birthday"`
	BirthdayLte string `sqFilter:"<,db=birthday" json:"birthday"`
}
```

Getting the conditions for squirrel using
```golang
import sqFilter "github.com/sergiocarracedo/squirrelfilter"
import sq "github.com/Masterminds/squirrel"
    conditions, err := sqFilter.GetConditions(filter) // Filter is a variable of type Filter with values
	... .Where(conditions)
```
This returns something like: 
```golang
sq.And{sq.Eq{"name": value}, sq.ILike{"surname": value} ...}
```

### Options

* Operator: `=,` `!=`, `<`, `<=`, `>`, `>=`, `like`, `contains`
  * `like` and  `contains` are case insensitive 
  * If no operator is set, `=` will be used
* Required: `required` makes mandatory provide a non-zero value for the field
* Target Field: `db=[field]` set the field name in database to filter. but default is the struct's field name in lowercase
    * The package can also use the tags `db:"field_name"`
