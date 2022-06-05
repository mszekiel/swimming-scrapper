package pools

import (
	"mszekiel/swimming-scrapper/common"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type PoolValidator struct {
	Pool struct {
		Name          string  `validate:"reqyured"`
		Identificator int     `validate:"reqyured"`
		Lat           float64 `validate:"reqyured"`
		Lon           float64 `validate:"reqyured"`
	} `json:"pool"`
	poolModel Pool `json:"-"`
}

func NewPoolValidator() PoolValidator {
	return PoolValidator{}
}

func (s *PoolValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, s)
	err = validate.Struct(s.Pool)

	if err != nil {
		return err
	}

	s.poolModel.Name = s.Pool.Name
	s.poolModel.Identificator = s.Pool.Identificator
	s.poolModel.Location.Lat = s.Pool.Lat
	s.poolModel.Location.Lon = s.Pool.Lon

	return nil
}
