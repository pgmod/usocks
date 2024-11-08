package main

import (
	"os"
	"strconv"

	"github.com/haxii/socks5"
)

func main() {
	counter := 0
	port := "1080"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	creds := socks5.StaticCredentials{}
	for {
		counter++
		if os.Getenv("USER_"+strconv.Itoa(counter)) != "" {
			creds[os.Getenv("USER_"+strconv.Itoa(counter))] = os.Getenv("PASS_" + strconv.Itoa(counter))
		} else {
			break
		}
	}
	conf := &socks5.Config{
		AuthMethods: []socks5.Authenticator{
			socks5.UserPassAuthenticator{
				Credentials: creds,
			},
		},
	}
	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}

	if err := server.ListenAndServe("tcp", "0.0.0.0:"+port); err != nil {
		panic(err)
	}
}
