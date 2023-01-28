package usecase

import (
	"errors"
	"strings"
	"time"

	"haioo.id/cart/config"
	"haioo.id/cart/domain/v1/user/repo"
	"haioo.id/cart/model"
	"haioo.id/cart/utl/password"
)

func UserList(conf config.Configuration) (users []model.User, err error) {
	db := conf.MysqlDB
	return repo.GetUserList(db)
}

// get user detail
func UserDetail(conf config.Configuration, userId string) (result *model.UserDetail, err error) {
	db := conf.MysqlDB
	getuser, err := repo.GetUserDetail(db, userId)
	if err != nil {
		return nil, err
	}

	getroles, err := repo.GetUserRoleMaps(db, userId)
	if err != nil {
		return nil, err
	}

	roleIds := make([]string, 0)
	for _, role := range *getroles {
		roleIds = append(roleIds, role.RoleId)
	}

	getaccess, err := repo.GetUserRoleAccess(db, roleIds)
	if err != nil {
		return nil, err
	}

	gAccess := make([]model.UserRoleAccess, 0)
	for _, access := range *getaccess {
		if access.IsAccess {
			gAccess = append(gAccess, access)
		}
	}

	return &model.UserDetail{
		User:   getuser,
		Roles:  getroles,
		Access: &gAccess,
	}, nil
}

// register new user
func UserRegister(conf config.Configuration, usr *model.User) (user model.User, err error) {
	var (
		payload = model.UserPayload{}
	)
	tx, err := conf.MysqlDB.Begin()
	if err != nil {
		return user, err
	}

	if strings.TrimSpace(usr.Name) == "" || strings.TrimSpace(usr.Password) == "" || strings.TrimSpace(usr.Email) == "" {
		return user, errors.New("data can not null")
	}

	hash, err := password.Encrypt(usr.Password)
	if err != nil {
		return *usr, err
	}

	payload.Password = string(hash)
	payload.Name = usr.Name
	payload.Email = usr.Email

	_, err = repo.RegisterUser(tx, &payload)
	if err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()

	return *usr, nil
}

// update user
func UserUpdate(conf config.Configuration, usr *model.User) (user model.User, err error) {
	var (
		payload = model.UserPayload{}
	)
	tx, err := conf.MysqlDB.Begin()
	if err != nil {
		return user, err
	}

	if strings.TrimSpace(usr.Name) == "" || strings.TrimSpace(usr.Email) == "" {
		return user, errors.New("data can not null")
	}

	payload.Name = usr.Name
	payload.Email = usr.Email
	payload.Id = usr.Id

	_, err = repo.UpdateUser(tx, &payload)
	if err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()

	return *usr, nil
}

// delete user
func UserDelete(conf config.Configuration, usr *model.User) (user model.User, err error) {
	var (
		payload = model.UserPayload{}
	)
	tx, err := conf.MysqlDB.Begin()
	if err != nil {
		return user, err
	}

	now := time.Now().Format("2006-01-02 15:04:05")
	payload.Id = usr.Id
	payload.Deleted = &now

	_, err = repo.DeleteUser(tx, &payload)
	if err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()

	return *usr, nil
}
