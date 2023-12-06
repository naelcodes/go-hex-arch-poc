package middleware

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/naelcodes/ab-backend/pkg/types"

	CustomErr "github.com/naelcodes/ab-backend/pkg/errors"
)

type DtoValidator interface {
	Validate() error
}

func PayloadValidator(createDTO DtoValidator, updateDTO DtoValidator) fiber.Handler {
	return func(c *fiber.Ctx) error {
		switch c.Method() {

		case fiber.MethodPost:

			payload := createDTO
			if err := validatePayload(c, payload); err != nil {
				return err
			}

			c.Locals("payload", payload)
			return c.Next()

		case fiber.MethodPatch:
			payload := updateDTO
			if err := validatePayload(c, payload); err != nil {
				return err
			}
			c.Locals("payload", payload)
			return c.Next()

		}
		return c.Next()
	}
}

func QueryValidator() fiber.Handler {
	return func(c *fiber.Ctx) error {
		queryParams := new(types.GetQueryParams)
		err := c.QueryParser(queryParams)

		if err != nil {
			return CustomErr.ServiceError(err, "Parsing query params")
		}

		if c.Method() == fiber.MethodGet {
			if queryParams.PageNumber != nil && queryParams.PageSize == nil {
				return CustomErr.ValidationError(errors.New("page size should be provided with page number"))
			}

			if queryParams.PageSize != nil && queryParams.PageNumber == nil {
				return CustomErr.ValidationError(errors.New("page number should be provided with page size"))
			}

			if queryParams.PageSize != nil && queryParams.PageNumber != nil {
				if *queryParams.PageSize <= 0 {
					return CustomErr.ValidationError(errors.New("page size should be greater than 0"))
				}

				if *queryParams.PageNumber < 0 {
					return CustomErr.ValidationError(errors.New("page number should be greater than or equal to 0"))
				}
			}
		}

		c.Locals("queryParams", queryParams)

		return c.Next()
	}
}

func validatePayload(c *fiber.Ctx, payload DtoValidator) error {
	if err := c.BodyParser(payload); err != nil {
		return CustomErr.ServiceError(err, "JSON Parsing")
	}

	//validate payload
	if err := payload.Validate(); err != nil {
		return CustomErr.ValidationError(err)
	}

	return nil
}
