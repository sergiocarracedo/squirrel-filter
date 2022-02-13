package squirrelfilter

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDbTarget(t *testing.T) {
	t.Run("no db target get default value", func(t *testing.T) {
		testFilter := struct {
			Name string `sqFilter:"="`
		}{
			Name: "sergio",
		}

		conditions, err := GetConditions(testFilter)
		assert.Nil(t, err)
		assert.Equal(t, sq.And{sq.Eq{"name": "sergio"}}, conditions)
	})

	t.Run("db target", func(t *testing.T) {
		testFilter := struct {
			Name string `sqFilter:"=,db=testOp"`
		}{
			Name: "sergio",
		}

		conditions, err := GetConditions(testFilter)
		assert.Nil(t, err)

		assert.Equal(t, sq.And{sq.Eq{"testOp": "sergio"}}, conditions)
	})

	t.Run("db target using db tag", func(t *testing.T) {
		testFilter := struct {
			Name string `sqFilter:"=" db:"test"`
		}{
			Name: "sergio",
		}

		conditions, err := GetConditions(testFilter)
		assert.Nil(t, err)
		assert.Equal(t, sq.And{sq.Eq{"test": "sergio"}}, conditions)
	})

	t.Run("db target using db tag and option", func(t *testing.T) {
		testFilter := struct {
			Name string `sqFilter:"=, db=testOp" db:"testDb"`
		}{
			Name: "sergio",
		}

		conditions, err := GetConditions(testFilter)
		assert.Nil(t, err)
		assert.Equal(t, sq.And{sq.Eq{"testOp": "sergio"}}, conditions)
	})
}

func TestRequired(t *testing.T) {
	t.Run("required fields fails", func(t *testing.T) {
		type TestFilter struct {
			Name string `sqFilter:"=, required,db=test"`
		}
		testFilter := TestFilter{
			Name: "",
		}

		_, err := GetConditions(testFilter)
		if assert.Error(t, err) {
			assert.IsType(t, ErrRequiredFilter{}, err)
		}
	})

	t.Run("required fields", func(t *testing.T) {
		type TestFilter struct {
			Name string `sqFilter:"=, required, db=test"`
		}
		testFilter := TestFilter{
			Name: "",
		}

		_, err := GetConditions(testFilter)
		if assert.Error(t, err) {
			assert.IsType(t, ErrRequiredFilter{}, err)
		}
	})

	t.Run("test simple filter: only operators", func(t *testing.T) {
		testFilter := struct {
			Name string `sqFilter:"="`
			Age  int    `sqFilter:">"`
		}{
			Name: "sergio",
			Age:  44,
		}

		conditions, err := GetConditions(testFilter)
		assert.Nil(t, err)
		assert.Equal(t, sq.And{sq.Eq{"name": "sergio"}, sq.Gt{"age": 44}}, conditions)
	})
}

func TestOperators(t *testing.T) {

	t.Run("default value", func(t *testing.T) {
		type TestFilter struct {
			Name string `sqFilter:""`
		}
		testFilter := TestFilter{
			Name: "test",
		}

		conditions, err := GetConditions(testFilter)
		assert.Nil(t, err)
		assert.Equal(t, sq.And{sq.Eq{"name": "test"}}, conditions)
	})

	t.Run("equal", func(t *testing.T) {
		type TestFilter struct {
			Name string `sqFilter:"="`
		}
		testFilter := TestFilter{
			Name: "test",
		}

		conditions, err := GetConditions(testFilter)
		assert.Nil(t, err)
		assert.Equal(t, sq.And{sq.Eq{"name": "test"}}, conditions)
	})

	t.Run("great", func(t *testing.T) {
		type TestFilter struct {
			Name string `sqFilter:">"`
		}
		testFilter := TestFilter{
			Name: "test",
		}

		conditions, err := GetConditions(testFilter)
		assert.Nil(t, err)
		assert.Equal(t, sq.And{sq.Gt{"name": "test"}}, conditions)
	})

	t.Run("great or equal", func(t *testing.T) {
		type TestFilter struct {
			Name string `sqFilter:">="`
		}
		testFilter := TestFilter{
			Name: "test",
		}

		conditions, err := GetConditions(testFilter)
		assert.Nil(t, err)
		assert.Equal(t, sq.And{sq.GtOrEq{"name": "test"}}, conditions)
	})

	t.Run("less", func(t *testing.T) {
		type TestFilter struct {
			Name string `sqFilter:"<"`
		}
		testFilter := TestFilter{
			Name: "test",
		}

		conditions, err := GetConditions(testFilter)
		assert.Nil(t, err)
		assert.Equal(t, sq.And{sq.Lt{"name": "test"}}, conditions)
	})

	t.Run("less or equal", func(t *testing.T) {
		type TestFilter struct {
			Name string `sqFilter:"<="`
		}
		testFilter := TestFilter{
			Name: "test",
		}

		conditions, err := GetConditions(testFilter)
		assert.Nil(t, err)
		assert.Equal(t, sq.And{sq.LtOrEq{"name": "test"}}, conditions)
	})

	t.Run("like", func(t *testing.T) {
		type TestFilter struct {
			Name string `sqFilter:"like"`
		}
		testFilter := TestFilter{
			Name: "test",
		}

		conditions, err := GetConditions(testFilter)
		assert.Nil(t, err)
		assert.Equal(t, sq.And{sq.ILike{"name": "test%"}}, conditions)
	})

	t.Run("contains", func(t *testing.T) {
		type TestFilter struct {
			Name string `sqFilter:"contains"`
		}
		testFilter := TestFilter{
			Name: "test",
		}

		conditions, err := GetConditions(testFilter)
		assert.Nil(t, err)
		assert.Equal(t, sq.And{sq.ILike{"name": "%test%"}}, conditions)
	})
}
