package v1

import (
	"net/http"
	"strconv"

	"github.com/kwarabei/Discoverish/internal/handlers/v1/requests"
	"github.com/kwarabei/Discoverish/internal/handlers/v1/responses"
	"github.com/kwarabei/Discoverish/internal/models"
	"github.com/labstack/echo/v4"
)

func (h *Handler) InitUserRoutes(apiGroup *echo.Group) {
	users := apiGroup.Group("/users")
	{
		users.GET("/:id", h.getUser)
		users.GET("/", h.getUserList)
		users.POST("/", h.createUser)
		users.PATCH("/", h.updateUser)
		users.DELETE("/", h.deleteUser)
	}
}

// @Summary User SignUp
// @Tags users
// @Description create user account
// @Accept  json
// @Produce  json
// @Param user body requests.UserSignUpRequest true "User info for registration"
// @Success 201 {string} string
// @Router /v1/users/ [post]
func (h *Handler) createUser(c echo.Context) error {
	u := &requests.UserSignUpRequest{}
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(u); err != nil {
		return err
	}

	h.db.Create(&models.User{
		Username: u.Username,
	})

	return c.JSON(http.StatusCreated, u)
}

// @Summary User Get
// @Tags users
// @Description get one user account
// @Accept  json
// @Produce  json
// @Param id path int true "User id"
// @Success 200 {string} string
// @Router /v1/users/{id} [get]
func (h *Handler) getUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	u := &models.User{}
	h.db.First(&u, id)
	resp := responses.UserGetResponse{
		Id:            u.ID,
		Username:      u.Username,
		FirstName:     u.FirstName,
		LastName:      u.LastName,
		Position:      u.Position,
		Qualification: u.Qualification,
		About:         u.About,
		Skillset:      u.Skillset,
	}

	return c.JSON(http.StatusOK, resp)
}

// @Summary User GetList
// @Tags users
// @Description get a list of users
// @Accept  json
// @Produce  json
// @Param limit query int true "limit"
// @Param offset query int true "offset"
// @Success 200 {object} string
// @Router /v1/users/ [get]
func (h *Handler) getUserList(c echo.Context) error {
	limit, _ := strconv.Atoi(c.Param("limit"))
	offset, _ := strconv.Atoi(c.Param("offset"))

	users := []models.User{}
	h.db.Limit(limit).Offset(offset).Find(&users)

	resp := responses.UserGetListResponse{}
	for _, u := range users {
		resp.Users = append(resp.Users, responses.UserGetResponse{
			Id:            u.ID,
			Username:      u.Username,
			FirstName:     u.FirstName,
			LastName:      u.LastName,
			Position:      u.Position,
			Qualification: u.Qualification,
			About:         u.About,
			Skillset:      u.Skillset,
		})
	}

	return c.JSON(http.StatusOK, resp)
}

// @Summary User Update
// @Tags users
// @Description update user account
// @Accept  json
// @Produce  json
// @Param user body requests.UserUpdateRequest true "User info for update"
// @Success 202 {string} string
// @Router /v1/users/ [patch]
func (h *Handler) updateUser(c echo.Context) error {
	u := &requests.UserUpdateRequest{}
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(u); err != nil {
		return err
	}

	user := models.User{}
	user.ID = u.Id

	h.db.Model(&user).Updates(models.User{
		FirstName:     u.FirstName,
		LastName:      u.LastName,
		Position:      u.Position,
		Qualification: u.Qualification,
		About:         u.About,
		Skillset:      u.Skillset,
	})

	return c.JSON(http.StatusAccepted, u)
}

// @Summary User Delete
// @Tags users
// @Description delete user account
// @Accept  json
// @Produce  json
// @Param user body requests.UserDeleteRequest true "Info required to delete user's account"
// @Success 204 {string} string
// @Router /v1/users/ [delete]
func (h *Handler) deleteUser(c echo.Context) error {
	u := &requests.UserDeleteRequest{}
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(u); err != nil {
		return err
	}

	h.db.Delete(&models.User{}, u.Id)

	return c.String(http.StatusNoContent, "user's account is deleted")
}
