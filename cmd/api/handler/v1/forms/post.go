package forms

import (
	"context"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"github.com/talgat-ruby/multi-step-form-api/cmd/db/model"
	"github.com/talgat-ruby/multi-step-form-api/internal/response"
	"github.com/talgat-ruby/multi-step-form-api/pkg/utils"
)

type AddOns struct {
	OnlineService       bool `xml:"onlineService" json:"onlineService" form:"onlineService"`
	LargerStorage       bool `xml:"largerStorage" json:"largerStorage" form:"largerStorage"`
	CustomizableProfile bool `xml:"customizableProfile" json:"customizableProfile" form:"customizableProfile"`
}

type PostRequestBody struct {
	Name   string  `xml:"name" json:"name" form:"name" validate:"required"`
	Email  string  `xml:"email" json:"email" form:"email" validate:"required,email"`
	Phone  string  `xml:"phone" json:"phone" form:"phone" validate:"required,e164"`
	Plan   string  `xml:"plan" json:"plan" form:"plan" validate:"required,oneof=arcade advanced pro"`
	Period string  `xml:"period" json:"period" form:"period" validate:"required,oneof=monthly yearly"`
	AddOns *AddOns `xml:"addOns" json:"addOns" form:"addOns" validate:"required"`
}

func (h *Handler) Add(c echo.Context) error {
	ctx := c.Request().Context()
	h.log.InfoContext(ctx, "start Add", "path", c.Path())

	reqBody, err := h.postRequestBody(ctx, c)
	if err != nil {
		h.log.ErrorContext(
			ctx,
			"fail Add:: body binding error",
			"path", c.Path(),
			"error", err,
		)
		return c.String(http.StatusBadRequest, "bad request")
	}

	if validationError := h.postRequestValidationErrors(ctx, reqBody); validationError != nil {
		h.log.ErrorContext(
			ctx,
			"fail Add:: validation errors",
			"path", c.Path(),
		)
		return c.JSON(http.StatusBadRequest, response.Error{Error: validationError})
	}

	dbInput := postDBInput(reqBody)
	if err := h.db.AddForm(ctx, dbInput); err != nil {
		h.log.ErrorContext(
			ctx,
			"fail Add:: db add fail",
			"path", c.Path(),
			"error", err,
		)
		return c.JSON(http.StatusBadRequest, response.ErrorWithMessage{Error: response.WithMessage{Message: err.Error()}})
	}

	h.log.InfoContext(ctx, "success Add", "path", c.Path())
	return c.NoContent(http.StatusNoContent)
}

type validationError struct {
	Name   *string `json:"name,omitempty"`
	Email  *string `json:"email,omitempty"`
	Phone  *string `json:"phone,omitempty"`
	Plan   *string `json:"plan,omitempty"`
	Period *string `json:"period,omitempty"`
	AddOns *string `json:"addOns,omitempty"`
}

func (h *Handler) postRequestBody(_ context.Context, c echo.Context) (*PostRequestBody, error) {
	reqBody := new(PostRequestBody)
	if err := c.Bind(reqBody); err != nil {
		return nil, err
	}

	return reqBody, nil
}

func (h *Handler) postRequestValidationErrors(_ context.Context, reqBody *PostRequestBody) *validationError {
	if err := h.validate.Struct(reqBody); err != nil {
		vErr := new(validationError)

		for _, err := range err.(validator.ValidationErrors) {
			switch err.StructField() {
			case "Name":
				vErr.Name = utils.ToPtr("name is required")
			case "Email":
				vErr.Email = utils.ToPtr("email is invalid")
			case "Phone":
				vErr.Phone = utils.ToPtr("phone is invalid")
			case "Plan":
				vErr.Plan = utils.ToPtr("plan is invalid")
			case "Period":
				vErr.Period = utils.ToPtr("period is invalid")
			case "AddOns":
				vErr.AddOns = utils.ToPtr("add-ons are required")
			}
		}

		return vErr
	}

	return nil
}

func postDBInput(reqBody *PostRequestBody) model.Form {
	var inp model.Form

	if reqBody == nil {
		return inp
	}

	inp.Name = reqBody.Name
	inp.Email = reqBody.Email
	inp.Phone = reqBody.Phone
	inp.Plan = reqBody.Plan
	inp.Period = reqBody.Period

	if reqBody.AddOns == nil {
		return inp
	}

	inp.OnlineService = reqBody.AddOns.OnlineService
	inp.LargerStorage = reqBody.AddOns.LargerStorage
	inp.CustomizableProfile = reqBody.AddOns.CustomizableProfile

	return inp
}
