```

_"required"_ forces to set this filter the value can't be empty
_operator_ Condition operator =, !=, <, <=, >, >=, like, in, isnull, etc
_group_operator_ And or Or

`sqfilter:"[required],[/operator/]`
if the field is an struct
`sqfilter:"[required],[/group_operator/]`

type Filter struct {
	Name     string `sqfilter:"=, required" json:"name" db:"name"`
	Surname  string `sqfilter:"like" json:"surname" db:"surname`
	City     string `sqfilter:"contains" json:"city"`
	BirthdayGte string `sqfilter:">" qgdb:"birthday" json:"birthday"`
	BirthdayLte string `sqfilter:"<" qgdb:"birthday" json:"birthday"`
}

```