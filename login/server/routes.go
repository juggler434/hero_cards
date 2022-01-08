package app

func (s *server) routes() {
	s.router.HandleFunc("/health", s.handleHealth())
}
