package server

import (
	"log"
	"net/http"
	"sync"

	"github.com/audit-brands/assurcast-reader/src/storage"
	"github.com/audit-brands/assurcast-reader/src/worker"
)

type Server struct {
	Addr        string
	db          *storage.Storage
	worker      *worker.Worker
	cache       map[string]interface{}
	cache_mutex *sync.Mutex

	BasePath string

	// auth
	Username string
	Password string
	// https
	CertFile string
	KeyFile  string
}

func NewServer(db *storage.Storage, addr string) *Server {
	return &Server{
		db:          db,
		Addr:        addr,
		worker:      worker.NewWorker(db),
		cache:       make(map[string]interface{}),
		cache_mutex: &sync.Mutex{},
	}
}

func (h *Server) GetAddr() string {
	proto := "http"
	if h.CertFile != "" && h.KeyFile != "" {
		proto = "https"
	}
	return proto + "://" + h.Addr + h.BasePath
}

func (s *Server) Start() {
	refreshRate := s.db.GetSettingsValueInt64("refresh_rate")
	s.worker.FindFavicons()
	s.worker.StartFeedCleaner()
	s.worker.SetRefreshRate(refreshRate)
	// Always trigger an initial refresh on startup so users see items immediately
	// — important for the pre-seeded Assurcast feed bundle. RefreshFeeds is a
	// no-op when the feed list is empty, so this is safe regardless.
	s.worker.RefreshFeeds()

	httpserver := &http.Server{Addr: s.Addr, Handler: s.handler()}

	var err error
	if s.CertFile != "" && s.KeyFile != "" {
		err = httpserver.ListenAndServeTLS(s.CertFile, s.KeyFile)
	} else {
		err = httpserver.ListenAndServe()
	}
	if err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
