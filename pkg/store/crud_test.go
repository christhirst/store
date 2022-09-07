package store

import (
	"fmt"
	"testing"
)

func TestAddEntryToDb(t *testing.T) {
	bdb, err := InitDB()
	if err != nil {
		t.Error(err)
	}
	defer bdb.Close()
	key := "cli44ent 4"
	key2 := "client2"
	key3 := "client3"
	clientConf := []string{} //oauth.Registration{Client_id: "test2", Redirect_uris: []string{"https://test.de"}}
	b := StructToDB(clientConf)
	n := map[string][]byte{key: b.Bytes(), key2: b.Bytes(), key3: b.Bytes()}
	for i, v := range n {
		err := AddEntryToDb(bdb, i, v)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestDeleteEntryToDb(t *testing.T) {
	bdb, err := InitDB()
	if err != nil {
		t.Error(err)
	}
	defer bdb.Close()
	key := "cli44ent 4"
	clientConf := []string{} //oauth.Registration{Client_id: "test2", Redirect_uris: []string{"https://test.de"}}
	b := StructToDB(clientConf)
	n := map[string][]byte{key: b.Bytes()}
	for i, v := range n {
		err := AddEntryToDb(bdb, i, v)
		if err != nil {
			t.Error(err)
		}
	}
	sslice := []string{key}

	errs := DeleteEntryToDb(bdb, sslice)
	for _, v := range errs {
		if v != nil {
			t.Error()
		}
	}

}

func TestViewEntryDb(t *testing.T) {
	bdb, err := InitDB()
	if err != nil {
		t.Error(err)
	}
	defer bdb.Close()
	key := "client testclient"
	clientConf := []string{} // oauth.Registration{Client_id: "test2", Redirect_uris: []string{"https://test.de"}}
	b := StructToDB(clientConf)
	n := map[string][]byte{key: b.Bytes()}
	for i, v := range n {
		err := AddEntryToDb(bdb, i, v)
		if err != nil {
			t.Error(err)
		}
	}

	bentry, err := ViewEntryDb(bdb, key)
	if err != nil {
		t.Error(err)
	}
	_, err = ByteToStruct(bentry, []string{})
	if err != nil {

		t.Error()
	}
}

func TestPrefixDb(t *testing.T) {
	bdb, err := InitDB()
	if err != nil {
		t.Error(err)
	}
	defer bdb.Close()
	key := "client testclient"
	clientConf := []string{} //oauth.Registration{Client_id: "test2", Redirect_uris: []string{"https://test.de"}}
	b := StructToDB(clientConf)
	n := map[string][]byte{key: b.Bytes()}
	for i, v := range n {
		err := AddEntryToDb(bdb, i, v)
		if err != nil {
			t.Error(err)
		}
	}

	clientlist, err := PrefixDb(bdb, "client")
	if len(clientlist) < 1 || err != nil {
		t.Error(err)
		for i, v := range clientlist {
			fmt.Println(i)
			entry, _ := ByteToStruct(v, []string{})
			fmt.Println(entry)
		}
	}
}
