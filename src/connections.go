package main

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"strings"
)

const (
	// commands
	SET_VAL = "set"
	GET_VAL = "get"
	DEL_VAL = "delete"
	EXPORT  = "export"
	IMPORT  = "import"
	EXIT    = "exit"
)

// HandleConnection allows clients to connect to database
func HandleConnections(d *Database) error {

	// listen
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", d.port))
	if err != nil {
		return err
	}

	// handle connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		go HandleConn(conn, d)
	}
}

// HandleConn handles single connection (read request, than write response)
func HandleConn(conn net.Conn, d *Database) error {
	defer conn.Close()

	// read connection
	reader := bufio.NewReader(conn)
	message, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	// remove outer whitespaces
	message = strings.TrimSpace(message)

	// close if command is "exit"
	if message == EXIT {
		return nil
	}

	// starts with SET_VAL
	if strings.HasPrefix(message, SET_VAL) {
		// extract key and value
		key_val := strings.TrimSpace(strings.ReplaceAll(message, SET_VAL, "")) // "key:value"
		key, val := strings.Split(key_val, ":")[0], strings.Split(key_val, ":")[1]

		// set value
		err := d.Set(key, val)
		if err != nil {
			conn.Write([]byte(err.Error() + "\n"))
			return err
		}

		// send response
		conn.Write([]byte("ok\n"))

		// starts with GET_VAL
	} else if strings.HasPrefix(message, GET_VAL) {
		// extract key and value
		key_val := strings.TrimSpace(strings.ReplaceAll(message, GET_VAL, "")) // "key:value"
		key := strings.Split(key_val, ":")[0]

		// get value
		value, err := d.Get(key)
		if err != nil {
			conn.Write([]byte(err.Error() + "\n"))
			return err
		}

		// send response
		conn.Write([]byte(value + "\n"))

		// starts with DEL_VAL
	} else if strings.HasPrefix(message, DEL_VAL) {
		// extract key and value
		key_val := strings.TrimSpace(strings.ReplaceAll(message, DEL_VAL, "")) // "key:value"
		key := strings.Split(key_val, ":")[0]

		// delete value
		err := d.Delete(key)
		if err != nil {
			conn.Write([]byte(err.Error() + "\n"))
			return err
		}

		// send response
		conn.Write([]byte("ok\n"))

		// starts with EXPORT
	} else if strings.HasPrefix(message, EXPORT) {
		err := d.ExportCsv()
		if err != nil {
			conn.Write([]byte(err.Error() + "\n"))
			return err
		}

		// send response
		conn.Write([]byte("ok\n"))

		// starts with IMPORT
	} else if strings.HasPrefix(message, IMPORT) {
		err := d.ImportCsv()
		if err != nil {
			conn.Write([]byte(err.Error() + "\n"))
			return err
		}

		// send response
		conn.Write([]byte("ok\n"))

		// unknown command
	} else {
		conn.Write([]byte("unknown command\n"))
		return errors.New("unknown command")
	}

	return nil
}
