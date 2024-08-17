package server

import (
	"api-example/constants"
	"api-example/internal/domain"
	"api-example/internal/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	validate    *validator.Validate
	userService service.IUserService
}

func NewUserHandler(validate *validator.Validate, userService service.IUserService) *UserHandler {
	return &UserHandler{
		validate:    validate,
		userService: userService,
	}
}

type CreateUserRequest struct {
	FullName string `json:"fullname"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

// CreateUser handling create user request
func (h *UserHandler) CreateUser(c echo.Context) error {
	ctx := c.Request().Context()
	// parse body to struct
	var request CreateUserRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request body"})
	}

	// validate request with validator
	err := h.validate.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request body", "error": err.Error()})
	}

	user := domain.NewUser(request.FullName, request.Email)
	user.SetPassword(request.Password)

	// call user service method
	err = h.userService.CreateUser(ctx, user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal server error", "error": err.Error()})
	}

	// mask password before return it to client
	user.Maskfields()
	return c.JSON(http.StatusCreated, map[string]any{"message": "User created successfully", "data": user})
}

type UpdateUserRequest struct {
	FullName string `json:"fullname"`
}

// UpdateUser handling update user
func (h *UserHandler) UpdateUser(c echo.Context) error {
	ctx := c.Request().Context()

	userID := c.Param("id")
	if userID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid id"})
	}

	var request UpdateUserRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request body"})
	}

	user := domain.User{
		ID:       uuid.MustParse(userID),
		FullName: request.FullName,
	}

	if err := h.userService.UpdateUser(ctx, &user); err != nil {
		if err == constants.ErrNoRows {
			return c.JSON(http.StatusNotFound, map[string]string{"message": "User not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal server error", "error": err.Error()})
	}
	c.JSON(http.StatusOK, map[string]string{
		"message": "User updated successfully",
	})

	return nil
}

// ListUser handling get list user
func (h *UserHandler) ListUser(c echo.Context) error {
	type Request struct {
		Fullname string `query:"fullname"`
		Email    string `query:"email"`
		Status   string `query:"status"`
	}
	ctx := c.Request().Context()

	var request Request
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request body"})
	}

	domainFilter := domain.User{
		FullName: request.Fullname,
		Email:    request.Email,
	}

	if request.Status != "" {
		switch request.Status {
		case "active":
			domainFilter.Status = domain.UserStatusActive
		case "inactive":
			domainFilter.Status = domain.UserStatusInactive
		}
	}

	users, err := h.userService.ListUser(ctx, &domainFilter)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal server error", "error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]any{"message": "User list", "data": users})
}
