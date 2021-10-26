package main

import (
	"fmt"
	"strings"
)

type log struct {
	module         string
	activeFunction string
	message        []string
}

func initialize(m string, af string, msg []string) log {
	activeLog := log{
		module:         m,
		activeFunction: af,
		message:        msg,
	}
	return activeLog
}

func (l log) print(msg []string) {
	l.message = msg
	fmt.Println("...............................................................................")
	fmt.Println("Module: " + l.module)
	fmt.Println("function: " + l.activeFunction)
	fmt.Println("Message: " + strings.Join(l.message, " || "))
	fmt.Println("...............................................................................")
}

func (l log) setModule(name string) log {
	l.module = name
	l.message = []string{""}
	return l
}

func (l log) setFunction(name string) log {
	fmt.Print("function is set")
	l.activeFunction = name
	l.message = []string{""}
	return l
}
