package main

type RouteRegister func()

type Logger interface {
	Infof(format string, args ...interface{})
	Error(args ...interface{})
}

func StartServer(l Logger, r RouteRegister) {
}
