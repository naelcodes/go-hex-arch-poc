package errors

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

const (
	notFoundMessage        = "Record not found"
	validationErrorMessage = "Validation error"
	repositoryErrorMessage = "Error in repository operation"
	serviceErrorMessage    = "Error in service operation"
	DomainErrorMessage     = "Error in domain operation"
	serverErrorMessage     = "Internal server error"
	unknownErrorMessage    = "Something went wrong"
)

type CustomError struct {
	Type string
	Err  error
}

func (c *CustomError) Error() string {
	return c.Err.Error()
}

func NewServiceError(err error) *CustomError {
	return &CustomError{
		Err:  err,
		Type: "ServiceError",
	}
}

func NewDomainError(err error) *CustomError {
	return &CustomError{
		Err:  err,
		Type: "DomainError",
	}
}

func NewRepositoryError(err error) *CustomError {
	return &CustomError{
		Err:  err,
		Type: "RepositoryError",
	}
}

func NewValidationError(err error) *CustomError {
	return &CustomError{
		Err:  err,
		Type: "ValidationError",
	}
}

func NewUnknownError(err error) *CustomError {
	return &CustomError{
		Err:  err,
		Type: "UnknownError",
	}
}

func NewServerError(err error) *CustomError {
	return &CustomError{
		Err:  err,
		Type: "ServerError",
	}
}

func GlobalErrorHandler(ctx *fiber.Ctx, err error) error {

	var fiberError *fiber.Error
	if customError, ok := err.(*CustomError); ok {

		switch customError.Type {
		case "ValidationError":
			return ctx.Status(fiber.StatusBadRequest).JSON(customError.Err)
		case "RepositoryError":
			return ctx.Status(fiber.StatusInternalServerError).JSON(customError.Err)
		case "DomainError":
			return ctx.Status(fiber.StatusBadRequest).JSON(customError.Err)
		case "ServiceError":
			return ctx.Status(fiber.StatusInternalServerError).JSON(customError.Err)
		case "ServerError":
			return ctx.Status(fiber.StatusInternalServerError).JSON(customError.Err)
		case "UnknownError":
			return ctx.Status(fiber.StatusInternalServerError).JSON(customError.Err)
		default:
			return ctx.Status(fiber.StatusInternalServerError).JSON(customError.Err)
		}

	} else if errors.As(err, &fiberError) {

		return ctx.Status(fiberError.Code).JSON(fiberError)
	} else {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}
}
