package main

import "errors"

const (
	START_DATA_CAP int = 10000

	// error messages
	MSG_KEY_EMPTY     = "key must not be empty"
	MSG_VALUE_EMPTY   = "value must not be empty"
	MSG_KEY_NOT_EXIST = "key does not exist"
)

// Database contains data (key-value pairs)
type Database struct {
	data map[string]string
	port int
}

// New creates and returns new instance of Database
func New(port int) Database {

	// init data map
	data := make(map[string]string, START_DATA_CAP)

	// create database
	var database Database = Database{data: data, port: port}

	return database
}

// Set assigns value to key
func (d *Database) Set(key, value string) error {
	if key == "" {
		return errors.New(MSG_KEY_EMPTY)
	}
	if value == "" {
		return errors.New(MSG_VALUE_EMPTY)
	}

	// set value
	d.data[key] = value
	return nil
}

// Get gets value of key
func (d Database) Get(key string) (string, error) {
	if key == "" {
		return "", errors.New(MSG_KEY_EMPTY)
	}

	// get value
	value := d.data[key]
	if value == "" {
		return "", errors.New(MSG_KEY_NOT_EXIST)
	}

	return value, nil
}

// Delete deletes key from database
func (d *Database) Delete(key string) error {
	if key == "" {
		return errors.New(MSG_KEY_EMPTY)
	}

	// check if key has value -> exists
	value := d.data[key]
	if value == "" {
		return errors.New(MSG_KEY_NOT_EXIST)
	}

	delete(d.data, key)

	return nil
}

// Size get number of elements (data)
func (d Database) Size() int {
	return len(d.data)
}
