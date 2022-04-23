package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

//connection interface to the database
type PixKeyRespositoryInterface interface {
	RegisterKey(pixKey *PixKey) (*PixKey, error)
	FindKeyByKind(key string, kind string) (*PixKey, error)
	AddBank(bank *Bank) error
	AddAccount(account *Account) error
	FindAccountById(id string) (*Account, error)
}

//main aggregate
type PixKey struct {
	Base      `valid:"required"`
	Kind      string   `json:"kind" valid: "notnull"`
	Key       string   `json:"key" valid: "notnull"`
	AccountID string   `json:"account_id" valid: "notnull"`
	Account   *Account `valid:"-"`
	Status    string   `json:"status" valid: "notnull"`
}

func (pixKey *PixKey) isValid() error {
	_, err := govalidator.ValidateStruct(pixKey)

	if pixKey.Kind != "email" && pixKey.Kind != "cpf" && pixKey.Kind != "phone" {
		return errors.New("kind must be email, cpf or phone")
	}
	if pixKey.Status != "active" && pixKey.Status != "inactive" {
		return errors.New("status must be active or inactive")
	}
	if err != nil {
		return err
	}
	return nil
}

func NewPixKey(kind string, account *Account, key string) (*PixKey, error) {
	pixKey := PixKey{
		Kind:    kind,
		Key:     key,
		Account: account,
		Status:  "active",
	}
	pixKey.ID = uuid.NewV4().String()
	pixKey.CreatedAt = time.Now()

	err := pixKey.isValid()
	if err != nil {
		return nil, err
	}
	return &pixKey, nil
}
