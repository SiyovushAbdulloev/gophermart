package httpserver

type Option func(*Server)

func WithAddress(addres string) Option {
	return func(s *Server) {
		s.Address = addres
	}
}
