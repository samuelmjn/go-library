package httpsvc

import (
	"github.com/labstack/echo"
	"github.com/samuelmjn/go-library/repository"
)

// Service :nodoc:
type Service struct {
	bookRepository repository.BookRepository
	userRepository repository.UserRepository
	echoHandler    *echo.Echo
}

// NewService :nodoc:
func NewService() *Service {
	return new(Service)
}

// RegisterBookRepository :nodoc:
func (s *Service) RegisterBookRepository(b repository.BookRepository) {
	s.bookRepository = b
}

// RegisterUserRepository :nodoc:
func (s *Service) RegisterUserRepository(u repository.UserRepository) {
	s.userRepository = u
}

// RegisterEcho :nodoc:
func (s *Service) RegisterEcho(e *echo.Echo) {
	s.echoHandler = e
}
