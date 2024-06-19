package challenge

import (
	"challenge-go/api/v1/commons"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"reflect"
)

type ChallengeRequest struct {
	Matrix [][]int `json:"matrix" validate:"required,challenge_matrix"`
}

type ChallengeResponse struct {
	Q  [][]float64 `json:"q"`
	R  [][]float64 `json:"r"`
	Qx [][]float32 `json:"qx,omitempty"`
	Rx [][]float32 `json:"rx,omitempty"`
}

func ValidateChallenge(c *fiber.Ctx) error {
	var validate = validator.New()
	validate.RegisterTagNameFunc(commons.FieldJSONName)
	_ = validate.RegisterValidation("challenge_matrix", CustomTypeSlice)

	var responseErr []*commons.IError
	var body = new(ChallengeRequest)
	_ = c.BodyParser(&body)

	err := validate.Struct(body)
	if err != nil {
		for _, val := range err.(validator.ValidationErrors) {
			var el commons.IError
			el.Field = val.Field()
			el.Tag = val.Tag()
			el.Value = val.Param()
			responseErr = append(responseErr, &el)
		}
		return commons.BadRequestAndData(c, "", 400, responseErr)
	}
	return c.Next()
}

func CustomTypeSlice(fl validator.FieldLevel) bool {
	switch v := fl.Field(); v.Kind() {
	case reflect.Array:
		return false
	case reflect.Slice:
		data := fl.Field().Interface().([][]int)
		if len(data) == 0 {
			return false
		}
		return true
	default:
		return false
	}
}
