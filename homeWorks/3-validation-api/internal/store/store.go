package store

import (
	"encoding/json"
	"os"
)

type DB struct {
	nameFile string
}

type EmailNotFoundError struct {
	Msg string
}

func (e *EmailNotFoundError) Error() string {
	return e.Msg
}

type dbEntry struct {
	Email string `json:"email"`
	Hash  string `json:"hash"`
}

func NewDB(nameDB string) (*DB, error) {
	file, err := os.OpenFile(nameDB, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	file.Close()
	return &DB{nameFile: nameDB}, nil
}

func (db *DB) AddHash(email string, hash string) error {
	file, err := os.OpenFile(db.nameFile, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	var dbEntrys []dbEntry
	err = json.NewDecoder(file).Decode(&dbEntrys)
	if err != nil {
		return err
	}
	dbEntrys = append(dbEntrys, dbEntry{Email: email, Hash: hash})

	// Переместить указатель и очистить файл
	if _, err := file.Seek(0, 0); err != nil {
		return err
	}
	if err := file.Truncate(0); err != nil {
		return err
	}
	err = json.NewEncoder(file).Encode(&dbEntrys)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) GetEmailByHash(hash string) (string, error) {
	file, err := os.OpenFile(db.nameFile, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}
	defer file.Close()
	var dbEntrys []dbEntry
	err = json.NewDecoder(file).Decode(&dbEntrys)
	if err != nil {
		return "", err
	}
	for _, i := range dbEntrys {
		if hash == i.Hash {
			return i.Email, nil
		}
	}
	return "", &EmailNotFoundError{Msg: "Email not found"}
}

func (db *DB) DelHash(email string, hash string) error {
	file, err := os.OpenFile(db.nameFile, os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	var dbEntrys []dbEntry
	err = json.NewDecoder(file).Decode(&dbEntrys)
	if err != nil {
		return err
	}
	serchEntry := dbEntry{Email: email, Hash: hash}
	filtered := dbEntrys[:0]
	for _, i := range dbEntrys {
		if i != serchEntry {
			filtered = append(filtered, i)
		}
	}
	// Переместить указатель и очистить файл
	if _, err := file.Seek(0, 0); err != nil {
		return err
	}
	if err := file.Truncate(0); err != nil {
		return err
	}
	err = json.NewEncoder(file).Encode(&dbEntrys)
	if err != nil {
		return err
	}
	return nil
}
