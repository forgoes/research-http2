package main

import (
	"net"
	"net/http"

	"golang.org/x/net/http2"

	"github.com/forgoes/logging"
	lh "github.com/forgoes/logging/handler"
)

var (
	l *logging.Logger
)

func init() {
	l = logging.GetLogger("main")
	l.SetLevel(logging.DEBUG)
	l.AddHandler(lh.NewStdoutHandler(lh.StdFormatter))
}

type Handler struct {
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	println(r.URL.String())
	w.Header().Set("Foo", "Bar")
}

func main() {
	handler := &Handler{}

	httpServer := &http.Server{
		Addr:              "",
		Handler:           handler,
		TLSConfig:         nil,
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}

	http2Server := &http2.Server{
		MaxHandlers:                  0,
		MaxConcurrentStreams:         0,
		MaxReadFrameSize:             0,
		PermitProhibitedCipherSuites: false,
		IdleTimeout:                  0,
		MaxUploadBufferPerConnection: 0,
		MaxUploadBufferPerStream:     0,
		NewWriteScheduler:            nil,
		CountError:                   nil,
	}

	err := http2.ConfigureServer(httpServer, http2Server)
	if err != nil {
		l.Fatal().E(err).Log()
	}

	ln, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		l.Fatal().E(err).Log()
	}

	err = httpServer.ServeTLS(ln, "./cert/hongkang.local.crt", "./cert/hongkang.local.key")
	if err != nil {
		l.Fatal().E(err).Log()
	}
}
