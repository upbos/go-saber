package http

import (
	"github.com/upbos/go-base/log"
	"net/http"
	"time"
)

type Server struct {
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func Setup(cfg Server, router http.Handler) {
	server := &http.Server{
		Addr:         cfg.Addr,
		Handler:      router,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	}

	log.Infof("Start http server listening %s", cfg.Addr)
	err := server.ListenAndServe()
	log.Errorf(err, "Start http server error, detail: %v")
}
