package main

import (
	"fmt"
	"testing"
)

// TestNew tests New function
func TestNew(t *testing.T) {
	port := 8000
	database := New(port)

	if database.data == nil {
		t.Log("data should not be nil")
		t.Fail()
	}
	if database.Size() != 0 {
		t.Log("number of keys should be 0")
		t.Fail()
	}
	if database.port != port {
		t.Log("port is incorrect")
		t.Fail()
	}
}

// TestSet tests Set function
func TestSet(t *testing.T) {
	database := New(8000)

	key, value := "key", "value"
	err := database.Set(key, value)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}

// TestGet tests Get function
func TestGet(t *testing.T) {
	database := New(8000)

	// insert value
	key, value := "key", "value"
	err := database.Set(key, value)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	// get value
	value_get, err := database.Get(key)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	if value_get != value {
		t.Log("inserted value is incorrect")
		t.Fail()
	}
}

// TestDelete tests Delete function
func TestDelete(t *testing.T) {
	database := New(8000)

	// add value
	key, value := "key", "value"
	err := database.Set(key, value)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	// delete value
	err = database.Delete(key)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if database.Size() != 0 {
		t.Log("database size should be 0")
		t.Log(database.Size())
		t.Fail()
	}

	// get value
	value_get, err := database.Get(key)
	if err == nil {
		t.Log("should throw error because key should not exist")
		t.Fail()
	}
	if value_get != "" {
		t.Log("value should be empty")
		t.Fail()
	}
}

// TestSize tests Size function
func TestSize(t *testing.T) {
	database := New(8000)
	count := 1

	// add
	key, value := "key", "value"
	err := database.Set(key, value)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if database.Size() != count {
		t.Log("count incorrect")
		t.Fail()
	}

	// add 10
	for i := 0; i < 10; i++ {
		database.Set(fmt.Sprintf("key%d", i), "value")
		count++
	}

	if database.Size() != count {
		t.Log("count incorrect")
		t.Fail()
	}
}
