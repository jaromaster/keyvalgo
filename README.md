# Keyvalgo
### simple, lightweight key-value database written in Go

## Features
- Encrypted connection
- Fast operations (get, set, delete)
- Persist and load data from csv files
- Extremely lightweight


# Get started

## Build
`go build -o kvg`

## RSA Keys
- create server.key <br>
`openssl genrsa -out server.key 2048` <br>
`openssl ecparam -genkey -name secp384r1 -out server.key`

- create server.crt <br>
`openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650`

- move to ./tls dir

## Start database
- set password <br>
`export KEYVALGO_PW=password` <br>

- start db <br>
`./kvg`
