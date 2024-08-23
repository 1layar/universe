package service

import (
	"context"
	"fmt"

	"github.com/1layar/universe/pkg/shared/dto"
	"github.com/openfga/go-sdk/client"
)

type AuthorizeService struct {
	client *client.OpenFgaClient
}

func NewAuthorizeService(client *client.OpenFgaClient) *AuthorizeService {
	return &AuthorizeService{
		client: client,
	}
}

func (s *AuthorizeService) Grant(permission dto.PermissionDto) error {

	user := fmt.Sprintf("%s:%v", permission.Role, permission.UserId)
	object := fmt.Sprintf("%s:%v", permission.Subject, permission.SubjectID)

	access := []client.ClientTupleKey{}

	for _, p := range permission.Permissions {
		access = append(access, client.ClientTupleKey{
			User:     user,
			Relation: p,
			Object:   object,
		})
	}

	body := client.ClientWriteRequest{
		Writes: access,
	}

	_, err := s.client.Write(context.Background()).
		Body(body).
		Execute()

	if err != nil {
		return err
	}

	return nil
}

func (s *AuthorizeService) Revoke(permission dto.PermissionDto) error {
	user := fmt.Sprintf("%s:%v", permission.Role, permission.UserId)
	object := fmt.Sprintf("%s:%v", permission.Subject, permission.SubjectID)

	access := []client.ClientTupleKeyWithoutCondition{}

	for _, p := range permission.Permissions {
		access = append(access, client.ClientTupleKeyWithoutCondition{
			User:     user,
			Relation: p,
			Object:   object,
		})
	}

	body := client.ClientWriteRequest{
		Deletes: access,
	}

	_, err := s.client.Write(context.Background()).
		Body(body).
		Execute()

	if err != nil {
		return err
	}

	return nil
}

func (s *AuthorizeService) CheckPermission(permission dto.PermissionCheckDTO) (bool, error) {
	user := fmt.Sprintf("%s:%v", permission.Role, permission.UserId)
	object := fmt.Sprintf("%s:%v", permission.Subject, permission.SubjectID)

	body := client.ClientCheckRequest{
		User:     user,
		Relation: permission.Permission,
		Object:   object,
	}

	result, err := s.client.Check(context.Background()).
		Body(body).
		Execute()

	if err != nil {
		return false, err
	}

	return *result.Allowed, nil
}
