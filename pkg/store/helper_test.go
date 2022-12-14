package store

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"testing"
)

func TestByteToStruct(t *testing.T) {
	clientConf := []string{} //oauth.Registration{Client_id: "test", Redirect_uris: []string{"https://test.de"}}

	var b bytes.Buffer
	enc := gob.NewEncoder(&b)

	err := enc.Encode(clientConf)
	if err != nil {
		log.Fatal("encode error:", err)
	}

	st, err := ByteToStruct(b.Bytes(), []string{})
	if err == nil {
		fmt.Println(st)
		t.Error()
	}

}
