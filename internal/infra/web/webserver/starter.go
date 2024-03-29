package webserver

type Starter struct {
	WebServer WebServer
}

func NewWebServerStarter(webServer WebServer) *Starter {
	return &Starter{
		WebServer: webServer,
	}
}
