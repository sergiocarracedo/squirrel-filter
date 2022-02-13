package squirrel_filter

import (
	"errors"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/fatih/structtag"
	"reflect"
	"strings"
)

const (
	OPERATOR_EQUAL          = "="
	OPERATOR_NOT_EQUAL      = "!="
	OPERATOR_LESS           = "<"
	OPERATOR_LESS_OR_EQUAL  = "<="
	OPERATOR_GREAT          = ">"
	OPERATOR_GREAT_OR_EQUAL = ">="
	OPERATOR_STARTS         = "like"
	OPERATOR_CONTAINS       = "contains"
)

var (
	operators = []string{OPERATOR_EQUAL, OPERATOR_NOT_EQUAL, OPERATOR_LESS, OPERATOR_LESS_OR_EQUAL,
		OPERATOR_GREAT, OPERATOR_GREAT_OR_EQUAL, OPERATOR_STARTS, OPERATOR_CONTAINS}
	defaultOperator = OPERATOR_EQUAL
)

const TAG = "sqfilter"

var (
	errNoDbTarget    = func(f string) error { return errors.New(fmt.Sprintf("%s: No db target field was set", f)) }
	errEmptyDbTarget = func(f string) error { return errors.New(fmt.Sprintf("%s: Empty db target", f)) }
)

type fieldOptions struct {
	Operator    string
	DbFieldName string
	Required    bool
}

func getOptions(field reflect.StructField) (options fieldOptions, err error) {
	var tags *structtag.Tags
	tags, err = structtag.Parse(string(field.Tag))
	if err != nil {
		panic(err)
	}
	sqFilterTag, _ := tags.Get(TAG)
	if sqFilterTag == nil {
		return
	}

	dbTag, _ := tags.Get("db")

	rawOptions := append(sqFilterTag.Options, sqFilterTag.Name)

	for _, option := range rawOptions {
		rawOption := strings.Trim(option, " ")
		op := strings.Split(rawOption, "=")
		optionName := strings.Trim(strings.ToLower(op[0]), " ")
		var optionValue string
		if len(op) == 2 {
			optionValue = strings.Trim(op[1], " ")
		}

		// Required
		if optionName == "required" {
			options.Required = true
			continue
		}

		// DbFieldName
		if optionName == "db" {
			if optionValue == "" {
				err = errEmptyDbTarget(field.Name)
				return
			}
			options.DbFieldName = optionValue
			continue
		}

		// Operator
		for _, operator := range operators {
			if rawOption == operator {
				options.Operator = operator
				break
			}
		}
	}

	if options.Operator == "" {
		options.Operator = defaultOperator
	}

	if options.DbFieldName == "" {
		if dbTag != nil {
			options.DbFieldName = dbTag.Name
		} else {
			err = errNoDbTarget(field.Name)
		}
	}
	return
}

func GetConditions(filter interface{}) (sqConditions sq.And, err error) {
	t := reflect.TypeOf(filter)
	v := reflect.ValueOf(filter)

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		value := v.Field(i)

		options, err := getOptions(f) //sqFilterTag, dbTag)
		if err != nil {
			panic(err)
		}

		var cond sq.Sqlizer

		switch options.Operator {
		case OPERATOR_EQUAL:
			cond = sq.Eq{options.DbFieldName: value}
		case OPERATOR_NOT_EQUAL:
			cond = sq.NotEq{options.DbFieldName: value}
		case OPERATOR_LESS:
			cond = sq.Lt{options.DbFieldName: value}
		case OPERATOR_LESS_OR_EQUAL:
			cond = sq.LtOrEq{options.DbFieldName: value}
		case OPERATOR_GREAT:
			cond = sq.Gt{options.DbFieldName: value}
		case OPERATOR_GREAT_OR_EQUAL:
			cond = sq.LtOrEq{options.DbFieldName: value}
		case OPERATOR_STARTS:
			cond = sq.Like{options.DbFieldName: fmt.Sprintf("%s%", value)}
		case OPERATOR_CONTAINS:
			cond = sq.Like{options.DbFieldName: fmt.Sprintf("%%s%", value)}
		}

		if cond != nil {
			sqConditions = append(sqConditions, cond.(sq.Sqlizer))
		}
	}

	return
}
