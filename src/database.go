package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
)

const (
	START_DATA_CAP int    = 10000 // start capacity of database
	DB_START_TEXT  string = "Starting database...done"
	PATH_CSV       string = "data.csv"

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
	fmt.Println(DB_START_TEXT)

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

// ExportCsv exports keys and values to csv file
func (d Database) ExportCsv() error {
	file, err := os.Create(PATH_CSV)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	// generate table
	data := [][]string{{"KEY", "VALUE"}}
	for k, v := range d.data {
		row := []string{k, v}
		data = append(data, row)
	}

	// export to file
	err = writer.WriteAll(data)
	if err != nil {
		return err
	}

	return nil
}

// ImportCsv imports keys and values from csv file
func (d *Database) ImportCsv() error {
	file, err := os.Open(PATH_CSV)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return err
	}

	// insert into database
	for i, row := range rows {
		// skip header
		if i == 0 {
			continue
		}

		err := d.Set(row[0], row[1])
		if err != nil {
			return err
		}
	}

	return nil
}
