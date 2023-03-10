package user

import (
	"net/http"

	"haioo.id/cart/domain/v1/user/usecase"
	"haioo.id/cart/model"
	"haioo.id/cart/utl/response"

	"github.com/labstack/echo/v4"
)

func (m *Module) HandleRest(group *echo.Group) {
	group.GET("/list", m.userList).Name = "user-list"
	group.GET("/detail/:param", m.userDetail).Name = "user-detail"
	group.POST("/register", m.userRegister).Name = "user-register"
	group.PUT("/update/:param", m.userUpdate).Name = "user-update"
	group.DELETE("/delete/:param", m.userDelete).Name = "user-delete"
}

// @Summary Get List User
// @Description get the status of server.
// @Tags Identity Access Management - User
// @Accept */*
// @Produce json
// @Success 200 {interface} model.Response{}
// @Router /api/v1/user/list [get]
func (m *Module) userList(c echo.Context) error {
	var (
		requestId = c.Get("request_id").(string)
	)

	// usecase get user list
	resp, err := usecase.UserList(m.Config)
	if err != nil {
		return response.Error(c, model.Response{
			LogId:   requestId,
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Error:   err,
		})
	}

	return response.Success(c, model.Response{
		LogId:   requestId,
		Status:  http.StatusOK,
		Message: nil,
		Data:    resp,
	})
}

// @Summary Register User
// @Description get the status of server.
// @Tags Identity Access Management - User
// @Accept */*
// @Produce json
// @Success 200 {interface} model.Response{}
// @Router /api/v1/user/register [post]
func (m *Module) userRegister(c echo.Context) error {

	var (
		requestId = c.Get("request_id").(string)
		email     = c.Get("email").(string)
		usr       = model.User{}
	)

	err := c.Bind(&usr)
	if err != nil {
		return response.Error(c, model.Response{
			LogId:   requestId,
			Status:  http.StatusBadRequest,
			Message: nil,
			Error:   err,
		})
	}

	usr.CreatedBy = &email
	usr.UpdatedBy = &email

	resp, err := usecase.UserRegister(m.Config, &usr)
	if err != nil {
		return response.Error(c, model.Response{
			LogId:   requestId,
			Status:  http.StatusBadRequest,
			Message: "error",
			Error:   err.Error(),
		})
	}

	return response.Success(c, model.Response{
		LogId:   requestId,
		Status:  http.StatusOK,
		Message: "user has been registered",
		Data:    resp,
	})
}

// @Summary Update User
// @Description get the status of server.
// @Tags Identity Access Management - User
// @Accept */*
// @Produce json
// @Param param path int true "User Id"
// @Success 200 {interface} model.Response{}
// @Router /api/v1/user/update/{param} [put]
func (m *Module) userUpdate(c echo.Context) error {
	var (
		requestId = c.Get("request_id").(string)
		usr       = model.User{}
	)

	id := c.Param("param")
	usr.Id = id

	err := c.Bind(&usr)
	if err != nil {
		return response.Error(c, model.Response{
			LogId:   requestId,
			Status:  http.StatusBadRequest,
			Message: nil,
			Error:   err,
		})
	}

	resp, err := usecase.UserUpdate(m.Config, &usr)
	if err != nil {
		return response.Error(c, model.Response{
			LogId:   requestId,
			Status:  http.StatusBadRequest,
			Message: nil,
			Error:   err.Error(),
		})
	}

	return response.Success(c, model.Response{
		LogId:   requestId,
		Status:  http.StatusOK,
		Message: "your data has been updated",
		Data:    resp,
	})
}

// @Summary Delete the User
// @Description get the status of server.
// @Tags Identity Access Management - User
// @Accept */*
// @Produce json
// @Param param path int true "User Id"
// @Success 200 {interface} model.Response{}
// @Router /api/v1/user/delete/{param} [delete]
func (m *Module) userDelete(c echo.Context) error {
	var (
		requestId = c.Get("request_id").(string)
		usr       = model.User{}
	)

	id := c.Param("param")
	usr.Id = id

	err := c.Bind(&usr)
	if err != nil {
		return response.Error(c, model.Response{
			LogId:   requestId,
			Status:  http.StatusBadRequest,
			Message: nil,
			Error:   err,
		})
	}

	_, err = usecase.UserDelete(m.Config, &usr)
	if err != nil {
		return response.Error(c, model.Response{
			LogId:   requestId,
			Status:  http.StatusBadRequest,
			Message: nil,
			Error:   err.Error(),
		})
	}

	return response.Success(c, model.Response{
		LogId:   requestId,
		Status:  http.StatusOK,
		Message: "your data has been updated",
		Data:    nil,
	})
}

// @Summary Detail of User
// @Description get the status of server.
// @Tags Identity Access Management - User
// @Accept */*
// @Produce json
// @Param param path int true "User Id"
// @Success 200 {interface} model.Response{}
// @Router /api/v1/user/detail/{param} [get]
func (m *Module) userDetail(c echo.Context) error {
	var (
		requestId = c.Get("request_id").(string)
	)

	resp, err := usecase.UserDetail(m.Config, c.Param("param"))
	if err != nil {
		return response.Error(c, model.Response{
			LogId:   requestId,
			Status:  http.StatusNotFound,
			Message: err.Error(),
			Data:    nil,
			Error:   err.Error(),
		})
	}

	return response.Success(c, model.Response{
		LogId:   requestId,
		Status:  http.StatusOK,
		Message: nil,
		Data:    resp,
		Error:   err,
	})
}
