package server

import (
	"context"
	"fmt"
	handler "github.com/Apartkktrain/go-file-server/router/RequestHandler"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func ServerSetup() {
	server := &http.Server{
		Addr:    ":80",
		Handler: http.HandlerFunc(RequestHandler),
	}
	go StartServer(server)
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	server.Shutdown(ctx)
}

func StartServer(server *http.Server) {
	err := server.ListenAndServe()
	if err != nil {
		fmt.Errorf("creating server is failed")
	}
}

func RequestHandler(Writer http.ResponseWriter, Request *http.Request) {
	if Request.URL.Path == "/" {
		Writer.WriteHeader(403)
		Writer.Header().Add("Content-Type", "application/xml")
		fmt.Fprintf(Writer, "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<Error>\n<Code>AccessDenied</Code>\n<Message>Access Denied</Message>\n<RequestId>XNMQRM20V567CS0Z</RequestId>\n<HostId>JvQydNg8KCurodNJe39rSXLZo6YEsQPv0l7EzujlQJ7shhiEiAY5phjG3HnhQupSuktEHwsm3Ic=</HostId>\n</Error>")
	}
	handler.Handler(Writer, Request)
}
