package database

import (
	"github.com/acheong08/SimpleOTP/internal/constants"
	"github.com/acheong08/SimpleOTP/internal/utilities"
)

type fileStore struct {
	PasswordHash string
	Entries      Entries
}

func (e *fileStore) Save() error {
	return utilities.SaveFile(e, constants.SaveFile)
}

func (e *fileStore) Load() error {
	err := utilities.LoadFile(e, constants.SaveFile)
	if err != nil {
		e = &fileStore{}
	}
	return nil

}

var FileStore fileStore = fileStore{
	Entries: Entries{
		Entries: make(map[string]string),
	},
}
