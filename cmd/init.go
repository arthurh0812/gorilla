package main

import "os"

// Env stores the environment variables passed in at runtime.
var Env map[string]string

var envList = []string{"PORT", "HOST"}

func init() {
	for _, e := range envList {
		Env[e] = os.Getenv(e)
	}
}
