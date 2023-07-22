package validator

import (
	"context"
	"net/http"
	"nginx/pkg/responser"

	"github.com/go-playground/validator/v10"
)

// Use a single instance of Validate, it caches struct info
var validate *validator.Validate

func init() {
	validate = validator.New()
}

// Validate struct fields
func ValidateStruct(ctx context.Context, s interface{}) error {
	return validate.StructCtx(ctx, s)
}

func ValidateRequestDto(ctx context.Context, s interface{}) *responser.ErrorResponse {
	e := responser.NewErrorBuilder().SetStatusCode(http.StatusBadRequest)
	if err := validate.StructCtx(ctx, s); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			errorBuilder := e.SetMessage("Validation failed")
			for _, validationError := range validationErrors {
				errorBuilder.SetDetail(validationError.Field(), validationError.Tag())
			}
			return errorBuilder.Build()
		}
		return e.SetMessage(err.Error()).Build()
	}
	return nil
}
