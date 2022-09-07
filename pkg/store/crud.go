package store

import (
	"fmt"

	badger "github.com/dgraph-io/badger/v3"
	"github.com/rs/zerolog/log"
)

func AddEntryToDb(db *badger.DB, clientname string, m []byte) error {
	err := db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(clientname), m)
		return err
	})
	return err

}

func DeleteEntryToDb(db *badger.DB, keys []string) []error {
	var errorlist []error
	for _, v := range keys {
		err := db.Update(func(txn *badger.Txn) error {
			err := txn.Delete([]byte(v))
			return err
		},
		)
		errorlist = append(errorlist, err)

	}
	return errorlist
}

func ViewEntryDb(db *badger.DB, key string) ([]byte, error) {
	var val []byte
	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		val, err = item.ValueCopy(nil)

		if err != nil {
			return err
		}
		/*
			var studentDecode auth.ClientData
			d := gob.NewDecoder(bytes.NewReader(val))
			if err := d.Decode(&studentDecode); err != nil {
				panic(err)
			}

			log.Printf("Decoded Struct from badger : name [%s] age [%s]\n", studentDecode.ClientID, studentDecode.RedirectURI) */

		return nil
	})
	return val, err

}

func PrefixDb(db *badger.DB, key string) (map[string][]byte, error) {
	clientList := make(map[string][]byte)
	err := db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		prefix := []byte(key)
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()
			k := item.Key()

			err := item.Value(func(v []byte) error {
				fmt.Printf("key=%s, value=%s\n", k, v)
				return nil
			})
			if err != nil {
				return err
			}
			keyc := string(item.KeyCopy(nil))
			val, err := item.ValueCopy(nil)
			if err != nil {
				log.Error().Err(err).Msg("Unable to copy db item")
			}
			clientList[keyc] = val

		}
		return nil
	})
	if err != nil {
		log.Error().Err(err).Msg("Unmarshal body")
	}
	return clientList, nil
}
