package usecase

import (
	"context"

	"github.com/krostyle/auth-systme-go/src/application/dto"
	"github.com/krostyle/auth-systme-go/src/application/interfaces"
	"github.com/krostyle/auth-systme-go/src/domain/entity"
	"github.com/krostyle/auth-systme-go/src/domain/interfaces/service"
	"github.com/krostyle/auth-systme-go/src/domain/repository"
)

type PermissionUseCase struct {
	permissionRepository repository.PermissionRepositoryInterface
	identifierGenerator  service.IdentifierGeneratorInterface
}

func NewPermissionUseCase(permissionRepository repository.PermissionRepositoryInterface, identifierGenerator service.IdentifierGeneratorInterface) interfaces.PermissionUseCaseInterface {
	return &PermissionUseCase{
		permissionRepository: permissionRepository,
		identifierGenerator:  identifierGenerator,
	}
}

func (p *PermissionUseCase) CreatePermission(ctx context.Context, permission *dto.PermissionCreateDTO) error {
	permissionEntity := &entity.Permission{
		ID:   p.identifierGenerator.GenerateUUID(),
		Name: permission.Name,
	}
	return p.permissionRepository.CreatePermission(ctx, permissionEntity)
}

func (p *PermissionUseCase) GetPermissionByID(ctx context.Context, id string) (*dto.PermissionDTO, error) {
	permission, err := p.permissionRepository.GetPermissionByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &dto.PermissionDTO{
		ID:        permission.ID,
		Name:      permission.Name,
		CreatedAt: permission.CreatedAt,
		UpdatedAt: permission.UpdatedAt,
	}, nil

}

func (p *PermissionUseCase) UpdatePermission(ctx context.Context, id string, permission *dto.PermissionUpdateDTO) error {
	permissionEntity := &entity.Permission{
		Name: permission.Name,
	}
	return p.permissionRepository.UpdatePermission(ctx, id, permissionEntity)
}

func (p *PermissionUseCase) DeletePermission(ctx context.Context, id string) error {
	return p.permissionRepository.DeletePermission(ctx, id)
}
