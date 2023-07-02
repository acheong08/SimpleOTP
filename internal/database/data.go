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
	return utilities.LoadFile(e, constants.SaveFile)
}

var FileStore fileStore = fileStore{
	Entries: Entries{},
}
