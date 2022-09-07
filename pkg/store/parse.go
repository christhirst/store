package store

import (
	"bytes"
	"encoding/gob"
)

func StructToDB(s interface{}) bytes.Buffer {
	var b bytes.Buffer
	e := gob.NewEncoder(&b)
	if err := e.Encode(s); err != nil {
		panic(err)
	}
	return b
}
