//go:build !windows && !macos
// +build !windows,!macos

package platform

import (
	"github.com/audit-brands/assurcast-reader/src/server"
)

func Start(s *server.Server) {
	s.Start()
}
