package web

type WebServer struct {
	enabled bool
}

func (w *WebServer) Run() error {
	w.enabled = true
	e := Init()
	e.Logger.Fatal(e.Start(":8089"))
	return nil
}

func (w *WebServer) Stop() error {
	w.enabled = false
	return nil
}

func (w *WebServer) Enabled() bool {
	return w.enabled
}

func NewWebServer() *WebServer {
	return &WebServer{}
}
