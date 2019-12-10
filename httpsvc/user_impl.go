package httpsvc

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/samuelmjn/go-library/repository/model"
)

func (s *Service) createUser(c echo.Context) (err error) {
	user := new(model.User)
	if err = c.Bind(user); err != nil {
		return
	}

	err = s.userRepository.Create(user)
	if err != nil {
		log.Println(err)
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}
