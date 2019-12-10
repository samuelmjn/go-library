package httpsvc

// InitializeRoutes :nodoc:
func (s *Service) InitializeRoutes() {
	// Book Handlers
	s.echoHandler.POST("/book/create", s.createBook)
	s.echoHandler.POST("/book/issue", s.issueBook)
	s.echoHandler.GET("/book/:book_id", s.findBookByID)
	s.echoHandler.GET("/book/popular", s.findMostIssuedBook)
	s.echoHandler.GET("/book/unissue/:issue_id", s.unissueBook)
	s.echoHandler.DELETE("/book/:book_id", s.deleteBook)

	// User Handlers
	s.echoHandler.POST("/user/create", s.createUser)
}
