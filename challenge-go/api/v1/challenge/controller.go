package challenge

import (
	"challenge-go/api/v1/commons"
	"challenge-go/pkg/tools"
	"github.com/gofiber/fiber/v2"
)

type Controller struct{}

func NewChallengeController() *Controller {
	return &Controller{}
}

// Challenge function
func (Controller) Challenge(c *fiber.Ctx) error {

	var body = new(ChallengeRequest)
	if err := c.BodyParser(&body); err != nil {
		return commons.BadRequestAndMessage(c, "Invalid request")
	}

	if !tools.ValidateMatrixRectangle(body.Matrix) {
		return commons.BadRequestAndMessage(c, "The matrix is not a rectangle")
	}

	matrix := tools.ConvertIntToFloat64(body.Matrix)

	var matrixTranspose = tools.Transpose(matrix)
	Q := tools.ApplyGramSchmidt(matrixTranspose)
	QTranspose := tools.Transpose(Q)
	//Q, R := tools.QRFactorization(m)
	R := tools.MultiplyMatrices(Q, matrix)

	Qx := tools.ConvertFloat64ToFloat32(removeScientificNotation(QTranspose))
	Rx := tools.ConvertFloat64ToFloat32(removeScientificNotation(R))
	//var data = map[string]interface{}{"Qx": Qx, "Rx": Rx, "Q": QTranspose, "R": RTranspose}

	var data = ChallengeResponse{
		Q:  QTranspose,
		R:  R,
		Qx: Qx,
		Rx: Rx,
	}

	return commons.OkAndData(c, data)
}

func removeScientificNotation(A [][]float64) [][]float64 {
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[i]); j++ {
			A[i][j] = float64(int(A[i][j]*1000000)) / 1000000
		}
	}
	return A
}
