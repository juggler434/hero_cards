package app

func (s *Server) routes() {
	s.router.HandleFunc("/health", s.handleHealth())
}
