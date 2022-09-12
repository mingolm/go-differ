package go_differ

import (
	"database/sql/driver"
	"reflect"
)

type EntryValues map[string]driver.Value

func (d *differ) getSnapshot(value interface{}) EntryValues {
	var (
		err error
		m   = map[string]driver.Value{}
		v   = reflect.Indirect(reflect.ValueOf(value))
		nf  = v.NumField()
	)

	for i := 0; i < nf; i++ {
		f := v.Type().Field(i)
		column := f.Tag.Get("differ")
		if column == "-" { // 跳过
			continue
		} else if column == "" {
			column = f.Name
		}
		face := v.Field(i).Interface()
		if v, ok := face.(driver.Valuer); ok {
			m[column], err = v.Value()
			if err != nil {
				panic(err.Error())
			}
		} else {
			m[column] = face
		}
	}

	for _, f := range d.ignoreFields {
		delete(m, f)
	}

	return m
}