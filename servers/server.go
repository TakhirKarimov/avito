package servers

import "log"

type Server interface {
	Run() error
	Enabled() bool
	Stop() error
}

type ServerManager struct {
	Servers []Server
}

func (s *ServerManager) AddServer(server Server) {
	s.Servers = append(s.Servers, server)
}

func (s *ServerManager) Run() error {
	for _, server := range s.Servers {
		if server.Enabled() {
			continue
		}
		go func(srv Server) {
			if err := srv.Run(); err != nil {
				log.Fatalf("failed start service manager. %v", err)
			}
		}(server)
	}
	return nil
}

func (s *ServerManager) Stop() error {
	for i := len(s.Servers) - 1; i >= 0; i-- {
		server := s.Servers[i]
		if !server.Enabled() {
			continue
		}
		if err := server.Stop(); err != nil {
			return err
		}
	}

	return nil
}
