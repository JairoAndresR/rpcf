package fixtures

import (
	uuid "github.com/satori/go.uuid"
	"rpcf/app/accounts/adapters/repositories/entities"
	"time"
)

func GetAccount() *entities.Account {
	now := time.Now()
	return &entities.Account{
		ID:        uuid.NewV1().String(),
		Name:      "John Doe",
		Email:     "johndoe@gmail.com",
		Password:  "12345",
		CreatedAt: &now,
		UpdatedAt: &now,
	}
}

func GetAccountEntity(id string) []map[string]interface{} {
	now := time.Now()
	return []map[string]interface{}{
		{
			"id":         id,
			"name":       "John Doe",
			"created_at": now.String(),
			"updated_at": now.String(),
		},
	}
}
