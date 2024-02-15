package service

import (
	"github.com/krostyle/auth-systme-go/src/domain/contract"
	"github.com/krostyle/auth-systme-go/src/domain/contract/service"
	"github.com/krostyle/auth-systme-go/src/domain/domainerror"
	"github.com/krostyle/auth-systme-go/src/domain/entity"
)

type UserService struct {
	User           *entity.User
	PasswordHasher service.PasswordHasherInterface
}

func NewUserService() contract.UserInterface {
	userService := &UserService{
		User: &entity.User{},
	}
	return userService
}

func (u *UserService) GetID() string {
	return u.User.ID
}

func (u *UserService) GetName() string {
	return u.User.Name
}

func (u *UserService) GetLastname() string {
	return u.User.Lastname
}

func (u *UserService) GetEmail() string {
	return u.User.Email
}

func (u *UserService) GetRoles() []contract.RoleInterface {
	rolesInterfaces := make([]contract.RoleInterface, len(u.User.Roles))
	for i := range u.User.Roles {
		rolesInterfaces[i] = &RoleService{Role: &u.User.Roles[i]}
	}
	return rolesInterfaces
}

func (u *UserService) SetID(id string) {
	u.User.ID = id
}

func (u *UserService) SetName(name string) {
	u.User.Name = name
}

func (u *UserService) SetLastname(lastname string) {
	u.User.Lastname = lastname
}

func (u *UserService) SetEmail(email string) {
	u.User.Email = email
}

func (u *UserService) HasRole(role contract.RoleInterface) bool {
	for _, r := range u.User.Roles {
		if r.ID == role.GetID() {
			return true
		}
	}
	return false
}

func (u *UserService) AddRole(role contract.RoleInterface) error {
	if u.HasRole(role) {
		return domainerror.ErrRoleExists
	}
	u.User.Roles = append(u.User.Roles, *role.(*RoleService).Role)
	return nil
}

func (u *UserService) RemoveRole(role contract.RoleInterface) error {
	for i, r := range u.User.Roles {
		if r.ID == role.GetID() {
			u.User.Roles = append(u.User.Roles[:i], u.User.Roles[i+1:]...)
			return nil
		}
	}
	return domainerror.ErrRoleNotFound
}
