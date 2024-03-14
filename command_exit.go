package main

import "os"

func commandExit(_ *commandConfig) error {
	os.Exit(0)
	return nil
}
