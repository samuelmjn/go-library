package httpsvc

import (
	"github.com/samuelmjn/go-library/repository/model"
)

type bookResponse struct {
	Book     model.Book  `json:"book"`
	IssuedBy *model.User `json:"issued_by,omitempty"`
}
