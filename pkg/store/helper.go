package store

import (
	"bytes"
	"encoding/gob"
	"errors"

	"github.com/rs/zerolog/log"
)

func ByteToStruct[T any](b []byte, v T) (T, error) {
	if b != nil {
		//var clientReg oauth.Registration

		d := gob.NewDecoder(bytes.NewReader(b))
		if err := d.Decode(v); err != nil {
			log.Error().Err(err).Msg("decode failed")
		}
		return v, nil
	}
	return v, errors.New("parameter empty")
}
