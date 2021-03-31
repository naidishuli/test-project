package utils

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"sync"
)

// RootPath return the root path of the project, independently where the call started.
var RootPath = func() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d) + "/"
}

// GetFunctionName return the name of the function/method that calls this implementation.
// This function will panic if something goes wrong in the system.
// Is used for writing mocking logics in the testing framework.
var GetFunctionName = func(skip int) string {
	pc, _, _, ok := runtime.Caller(skip)
	if !ok {
		panic(fmt.Errorf("cannot get the caller function name"))
	}

	function := runtime.FuncForPC(pc)
	parts := strings.Split(function.Name(), ".")
	return parts[len(parts)-1]
}

// CaptureStdOutput takes a function as parameter where all the calls to stdout/stderr should be done.
// It will return the combined string being put in those pipes.
// All the initializations that use the system print functionality should be done inside the function as well,
// only this way the function can intercept what was printed.
var CaptureStdOutput = func(f func()) string {
	mux := sync.Mutex{}

	reader, writer, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	mux.Lock()

	stdout := os.Stdout
	stderr := os.Stderr
	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
		log.SetOutput(os.Stderr)
		mux.Unlock()
	}()

	os.Stdout = writer
	os.Stderr = writer
	log.SetOutput(writer)

	out := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		var buf bytes.Buffer
		wg.Done()
		io.Copy(&buf, reader)
		out <- buf.String()
	}()

	wg.Wait()
	f()
	writer.Close()
	return <-out
}

// ToSnakeCase return the snake_case version of a string from CamelCase
var ToSnakeCase = func(s string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchFirstCap.ReplaceAllString(s, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
