package middleware

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/gofiber/fiber/v2"
	"github.com/naelcodes/ab-backend/internal/core/dto"
	"github.com/naelcodes/ab-backend/pkg/types"
	"github.com/naelcodes/ab-backend/pkg/utils"

	CustomErr "github.com/naelcodes/ab-backend/pkg/errors"
)

type DtoValidator interface {
	Validate() error
}

func PayloadValidator(createDTO DtoValidator, updateDTO DtoValidator) fiber.Handler {
	return func(c *fiber.Ctx) error {

		utils.Logger.Info("[PayloadValidator] - Validating payload")

		isImputationRoute, _ := regexp.Match(`/api/v1/invoices/\d+/imputations`, []byte(c.Path()))

		utils.Logger.Info(fmt.Sprintf("[PayloadValidator] - path: %v", c.Path()))
		utils.Logger.Info(fmt.Sprintf("[PayloadValidator] - Is imputation route: %v", isImputationRoute))

		if !isImputationRoute {

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
		}

		return c.Next()
	}
}

func ImputationPayloadValidator(payload []*dto.InvoiceImputationDTO) fiber.Handler {
	return func(c *fiber.Ctx) error {

		utils.Logger.Info("[ImputationPayloadValidator] - Validating payload")
		utils.Logger.Info(fmt.Sprintf("[ImputationPayloadValidator] - path: %v", c.Path()))

		switch c.Method() {
		case fiber.MethodPost:
			if err := c.BodyParser(&payload); err != nil {
				return CustomErr.ServiceError(err, "JSON Parsing")
			}

			if len(payload) == 0 {
				return CustomErr.ValidationError(errors.New("no imputations provided"))
			}

			//validate payload
			for _, p := range payload {
				if err := p.Validate(); err != nil {
					return CustomErr.ValidationError(err)
				}
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
