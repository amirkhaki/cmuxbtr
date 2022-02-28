package store

import (
	"encoding/gob"
	"fmt"
	"bytes"
)

func Encode(data interface{}) ([]byte, error){
	buf := new(bytes.Buffer)
	e := gob.NewEncoder(buf)
	if err := e.Encode(data); err != nil {
		return nil, fmt.Errorf("Could not encode to byte: %w", err)
	}
	return buf.Bytes(), nil
}

func Decode(data []byte, to interface{}) error {
	buf := bytes.NewBuffer(data)
	de := gob.NewDecoder(buf)
	if err := de.Decode(to); err != nil {
		return fmt.Errorf("Coud not decode from bytes: %w", err)
	}
	return nil
}
