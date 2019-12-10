package httpsvc

// InitializeRoutes :nodoc:
func (s *Service) InitializeRoutes() {
	// Book Handlers
	s.echoHandler.POST("/books/create", s.createBook)
	s.echoHandler.GET("/books", s.findAllBooks)
	s.echoHandler.POST("/books/issue", s.issueBook)
	s.echoHandler.GET("/books/:book_id", s.findBookByID)
	s.echoHandler.GET("/books/popular", s.findMostIssuedBook)
	s.echoHandler.GET("/books/unissue/:issue_id", s.unissueBook)
	s.echoHandler.DELETE("/books/:book_id", s.deleteBook)

	// User Handlers
	s.echoHandler.POST("/users/create", s.createUser)
}
