package interfaces

type AppScope interface {
	Services
	Log() Logger
	HTTPClient() HTTPClient
}
