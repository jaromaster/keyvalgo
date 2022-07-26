package main

import (
	"bufio"
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"strings"
)

const (
	OK        = "ok"
	CERT_DIR  = "tls/"
	CERT_NAME = "server.crt"
	KEY_NAME  = "server.key"

	// commands
	SET_VAL = "set"
	GET_VAL = "get"
	DEL_VAL = "delete"
	EXPORT  = "export"
	IMPORT  = "import"
	EXIT    = "exit"
)

// HandleConnection allows clients to connect to database
func (d *Database) HandleConnections() error {

	// setup tls
	cert, err := tls.LoadX509KeyPair(CERT_DIR+CERT_NAME, CERT_DIR+KEY_NAME)
	if err != nil {
		return err
	}
	tls_conf := &tls.Config{Certificates: []tls.Certificate{cert}}

	// listen with tls
	listener, err := tls.Listen("tcp", fmt.Sprintf(":%d", d.port), tls_conf)
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

	reader := bufio.NewReader(conn)

	// authenticate
	password, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	if !d.Auth(strings.TrimSpace(password)) {
		conn.Write([]byte("Auth failed: incorrect password\n"))
		conn.Close()
		return errors.New("password incorrect")
	}
	conn.Write([]byte("Auth successful\n"))

	// read connection (payload)
	message, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	// remove outer whitespaces
	message = strings.TrimSpace(message)

	// close if command is "exit"
	if message == EXIT {
		conn.Write([]byte("Closed connection\n"))
		conn.Close()
		return nil
	}

	// starts with SET_VAL
	if strings.HasPrefix(message, SET_VAL) {
		// extract key and value
		key_val := strings.TrimSpace(strings.ReplaceAll(message, SET_VAL, "")) // should be "key:value"
		args_list := strings.Split(key_val, ":")                               // should be ["key", "val"]

		if len(args_list) != 2 {
			invalid_args := "invalid number of arguments, should look like: set key:value"
			conn.Write([]byte(invalid_args + "\n"))
			return errors.New(invalid_args)
		}

		// set value
		key, val := args_list[0], args_list[1]
		err := d.Set(key, val)
		if err != nil {
			conn.Write([]byte(err.Error() + "\n"))
			return err
		}
		conn.Write([]byte(OK + "\n"))

		// starts with GET_VAL
	} else if strings.HasPrefix(message, GET_VAL) {
		// extract key
		key := strings.TrimSpace(strings.ReplaceAll(message, GET_VAL, "")) // should be "key"

		// get value
		value, err := d.Get(key)
		if err != nil {
			conn.Write([]byte(err.Error() + "\n"))
			return err
		}
		conn.Write([]byte(value + "\n"))

		// starts with DEL_VAL
	} else if strings.HasPrefix(message, DEL_VAL) {
		// extract key
		key := strings.TrimSpace(strings.ReplaceAll(message, DEL_VAL, "")) // "key"

		// delete value
		err := d.Delete(key)
		if err != nil {
			conn.Write([]byte(err.Error() + "\n"))
			return err
		}
		conn.Write([]byte(OK + "\n"))

		// equals EXPORT
	} else if message == EXPORT {
		err := d.ExportCsv()
		if err != nil {
			conn.Write([]byte(err.Error() + "\n"))
			return err
		}
		conn.Write([]byte(OK + "\n"))

		// equals IMPORT
	} else if message == IMPORT {
		err := d.ImportCsv()
		if err != nil {
			conn.Write([]byte(err.Error() + "\n"))
			return err
		}
		conn.Write([]byte(OK + "\n"))

		// unknown command
	} else {
		uk_command := "unknown command"
		conn.Write([]byte(uk_command + "\n"))
		return errors.New(uk_command)
	}

	return nil
}
