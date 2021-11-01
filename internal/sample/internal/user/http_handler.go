package user

import (
	"strconv"

	"github.com/arimaulana/point-of-no-return/internal/common/pkg/log"
	"github.com/gofiber/fiber/v2"
)

type ApiHttpHandler interface {
	GetUserByID(c *fiber.Ctx) error
	GetUserList(c *fiber.Ctx) error
}

type apiHttpHandler struct {
	service Service
	logger  log.Logger
}

func NewApiHttpHandler(svc Service, logger log.Logger) ApiHttpHandler {
	return apiHttpHandler{service: svc, logger: logger}
}

func (h apiHttpHandler) GetUserByID(c *fiber.Ctx) error {
	user, err := h.service.GetUserByID(c.UserContext(), c.Params("id"))
	if err != nil {
		return err
	}

	res := UserResponse{
		BaseResponse: BaseResponse{
			Message: "success",
		},
		Data: user,
	}
	return c.Status(fiber.StatusOK).JSON(res)
}

func (h apiHttpHandler) GetUserList(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return err
	}

	perpage, err := strconv.Atoi(c.Query("perpage", "10"))
	if err != nil {
		return err
	}

	data, total, err := h.service.GetUserList(c.UserContext(), page, perpage)
	if err != nil {
		return err
	}

	res := UserListResponse{
		BaseListResponse: BaseListResponse{
			BaseResponse: BaseResponse{
				Message: "success",
			},
			Page:    page,
			Perpage: perpage,
			Total:   total,
		},
		Data: data,
	}
	return c.Status(fiber.StatusOK).JSON(res)
}

type NoAuthHttpHandler interface{}

type noAuthHttpHandler struct {
	service Service
	logger  log.Logger
}

func NewNoAuthHttpHandler(svc Service, logger log.Logger) NoAuthHttpHandler {
	return noAuthHttpHandler{service: svc, logger: logger}
}

type AdminHttpHandler interface{}

type adminHttpHandler struct {
	service Service
	logger  log.Logger
}

func NewAdminHttpHandler(svc Service, logger log.Logger) AdminHttpHandler {
	return adminHttpHandler{service: svc, logger: logger}
}
