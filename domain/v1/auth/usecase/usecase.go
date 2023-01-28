package usecase

import (
	"errors"

	"haioo.id/cart/config"
	"haioo.id/cart/domain/v1/user/repo"
	"haioo.id/cart/model"
	"haioo.id/cart/utl/jwt"
	"haioo.id/cart/utl/password"
)

func Login(cnf config.Configuration, p *model.LoginPayload) (token model.AuthAccess, err error) {

	if p.Email == "" || p.Password == "" {
		err = errors.New("email password can not null")
		return
	}

	getUser, err := repo.GetUserDetailByParam(cnf.MysqlDB, "user_email", p.Email)
	if err != nil {
		err = errors.New("email not found")
		return
	}

	if getUser.Password == "" || getUser.Id == "" {
		err = errors.New("user not found")
		return
	}

	getroles, err := repo.GetUserRoleMaps(cnf.MysqlDB, getUser.Id)
	if err != nil {
		return token, err
	}

	roleIds := make([]string, 0)
	for _, role := range *getroles {
		roleIds = append(roleIds, role.RoleId)
	}

	getaccess, err := repo.GetUserRoleAccess(cnf.MysqlDB, roleIds)
	if err != nil {
		return token, err
	}

	gAccess := make([]string, 0)
	for _, access := range *getaccess {
		if access.IsAccess {
			gAccess = append(gAccess, access.Url)
		}
	}

	if err = password.Decrypt([]byte(getUser.Password), p.Password); err != nil {
		err = errors.New("password not match")
		return
	}

	return jwt.Generate(getUser.Email, gAccess)
}
