package httpsvc

import (
	"log"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/samuelmjn/go-library/repository/model"
)

func (s *Service) createBook(c echo.Context) (err error) {
	book := new(model.Book)
	if err = c.Bind(book); err != nil {
		return
	}

	err = s.bookRepository.Create(book)
	if err != nil {
		log.Println(err)
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, book)
}

func (s *Service) issueBook(c echo.Context) (err error) {
	issue := new(model.Issue)
	if err = c.Bind(issue); err != nil {
		return
	}

	err = s.bookRepository.Issue(issue)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.String(http.StatusNotFound, "book not found or being issued")
		}
		log.Println(err)
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, issue)
}

func (s *Service) findBookByID(c echo.Context) (err error) {
	var resp bookResponse

	bookID, _ := strconv.Atoi(c.Param("book_id"))
	book, err := s.bookRepository.FindByID(int64(bookID))
	if err != nil {
		log.Println(err)
		if err == gorm.ErrRecordNotFound {
			return c.String(http.StatusNotFound, "book not found")
		}

		return c.String(http.StatusBadRequest, err.Error())
	}

	issue, err := s.bookRepository.FindCurrentIssueByBookID(int64(bookID))
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println(err)

		return c.String(http.StatusBadRequest, err.Error())
	}

	if issue != nil {
		user, err := s.userRepository.FindByID(issue.IssuedBy)
		if err != nil {
			log.Println(err)
			if err == gorm.ErrRecordNotFound {
				return c.String(http.StatusNotFound, "user not found")
			}

			return c.String(http.StatusBadRequest, err.Error())
		}
		resp.IssuedBy = user
	}

	resp.Book = *book

	return c.JSON(http.StatusOK, resp)
}

func (s *Service) findMostIssuedBook(c echo.Context) (err error) {
	book, err := s.bookRepository.FindMostIssued()
	if err != nil {
		log.Println(err)
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, book)
}

func (s *Service) deleteBook(c echo.Context) (err error) {
	bookID, _ := strconv.Atoi(c.Param("book_id"))
	err = s.bookRepository.Delete(int64(bookID))
	if err != nil {
		log.Println(err)
		if err == gorm.ErrRecordNotFound {
			return c.String(http.StatusNotFound, "book not found")
		}

		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "book deleted succesfuly")
}

func (s *Service) unissueBook(c echo.Context) (err error) {
	issueID, _ := strconv.Atoi(c.Param("issue_id"))
	err = s.bookRepository.Unissue(int64(issueID))
	if err != nil {
		log.Println(err)
		if err == gorm.ErrRecordNotFound {
			return c.String(http.StatusNotFound, "issue not found")
		}

		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "book unissued succesfuly")
}

func (s *Service) findAllBooks(c echo.Context) (err error) {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	size, _ := strconv.Atoi(c.QueryParam("size"))
	log.Println(page, size)
	book, err := s.bookRepository.FindAllBooks(int64(page), int64(size))
	if err != nil {
		log.Println(err)
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, book)
}
