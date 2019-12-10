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

	issue, err := s.bookRepository.FindIssueByBookID(int64(bookID))
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println(err)

		return c.String(http.StatusBadRequest, err.Error())
	}

	if issue != nil {
		user, err := s.userRepository.FindByID(issue.IssuedBy)
		if err != nil {
			log.Println(err)
			if err == gorm.ErrRecordNotFound {
				return c.String(http.StatusNotFound, "book not found")
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
	bookID, _ := strconv.Atoi(c.Param("book_id"))
	err = s.bookRepository.Unissue(int64(bookID))
	if err != nil {
		log.Println(err)
		if err == gorm.ErrRecordNotFound {
			return c.String(http.StatusNotFound, "book not found")
		}

		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "book unissued succesfuly")
}
