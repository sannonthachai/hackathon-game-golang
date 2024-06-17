package transport

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	model "gitlab.com/sannonthachai/find-the-hidden-backend/model/user"
	"gitlab.com/sannonthachai/find-the-hidden-backend/util"
)

func (h *Handler) Register(c echo.Context) error {
	payload := model.User{}

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, util.CreateErrorResponse(nil, "11", "BadRequest", "BadRequest"))
	}

	if err := h.userService.Register(payload); err != nil {
		return c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(nil, "99", "ErrorUnexpected", "ErrorUnexpected"))
	}

	return c.JSON(http.StatusOK, util.CreateSuccessResponse(nil))
}

func (h *Handler) Login(c echo.Context) error {
	payload := model.Login{}

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, util.CreateErrorResponse(nil, "11", "BadRequest", "BadRequest"))
	}

	result, err := h.userService.Login(payload)
	if err != nil {
		if err.Error() == "401" {
			return c.JSON(http.StatusUnauthorized, util.CreateErrorResponse(nil, "12", "Unauthorized", "Unauthorized"))
		}
		return c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(nil, "99", "ErrorUnexpected", "ErrorUnexpected"))
	}

	return c.JSON(http.StatusOK, util.CreateSuccessResponse(result))
}

func (h *Handler) UpdateUserPoint(c echo.Context) error {
	payload := model.UserPointByChapter{}

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, util.CreateErrorResponse(nil, "11", "BadRequest", "BadRequest"))
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*model.JwtCustomClaims)
	userID := claims.UserID

	if err := h.userService.UpdateUserPoint(userID, payload); err != nil {
		return c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(nil, "99", "ErrorUnexpected", "ErrorUnexpected"))
	}

	return c.JSON(http.StatusOK, util.CreateSuccessResponse(nil))
}

func (h *Handler) GetUserPointByChapter(c echo.Context) error {
	chapter := c.QueryParam("chapter")
	chapterInt, err := strconv.Atoi(chapter)
	if err != nil {
		fmt.Println("Error strconv")
		return c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(nil, "99", "ErrorUnexpected", "ErrorUnexpected"))
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*model.JwtCustomClaims)
	userID := claims.UserID

	result, err := h.userService.GetUserPointByChapter(userID, chapterInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(nil, "99", "ErrorUnexpected", "ErrorUnexpected"))
	}

	return c.JSON(http.StatusOK, util.CreateSuccessResponse(result))
}

func (h *Handler) GetUserPoint(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*model.JwtCustomClaims)
	userID := claims.UserID

	result, err := h.userService.GetUserPoint(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(nil, "99", "ErrorUnexpected", "ErrorUnexpected"))
	}

	return c.JSON(http.StatusOK, util.CreateSuccessResponse(result))
}

func (h *Handler) GetLeaderBoardByChapter(c echo.Context) error {
	chapter := c.QueryParam("chapter")
	chapterInt, err := strconv.Atoi(chapter)
	if err != nil {
		fmt.Println("Error strconv")
		return c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(nil, "99", "ErrorUnexpected", "ErrorUnexpected"))
	}

	result, err := h.userService.GetLeaderBoardByChapter(chapterInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(nil, "99", "ErrorUnexpected", "ErrorUnexpected"))
	}

	return c.JSON(http.StatusOK, util.CreateSuccessResponse(result))
}

func (h *Handler) GetLeaderBoard(c echo.Context) error {
	result, err := h.userService.GetLeaderBoard()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(nil, "99", "ErrorUnexpected", "ErrorUnexpected"))
	}

	return c.JSON(http.StatusOK, util.CreateSuccessResponse(result))
}
