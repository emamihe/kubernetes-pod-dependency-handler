package main

import (
	"errors"
	"fmt"
	"os"
)

func validate() error {
	if *host == "" {
		return errors.New("hostname cannot set empty")
	}

	if *selfSigned {
		if !fileExists(*cacert) {
			return errors.New(
				fmt.Sprintf("cacert path %s not exists", *cacert) )
		}
	}

	if !fileExists(*token) {
		return errors.New(
			fmt.Sprintf("token path %s not exists", *token) )
	}

	return nil
}

func fileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
