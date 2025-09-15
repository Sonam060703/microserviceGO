package account

import (
	"context"

	"github.com/segmentio/ksuid"
)

// Server -> Service -> Repository -> database

// Service -> interface as contract
type Service interface {
	PostAccount(ctx context.Context, name string) (*Account, error)               // ------ >  PutAccount of repository interface implemented by
	GetAccount(ctx context.Context, id string) (*Account, error)                  // ------ >  GetAccountByID
	GetAccounts(ctx context.Context, skip uint64, take uint64) ([]Account, error) // ------ >  ListAccounts
}

// Account struct ( ID , Name) but in actual db value are (id , name)
type Account struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// accountService is the implementation of Service contract
type accountService struct {
	// Instance of Repository
	repository Repository
}

// Invoked in main.go
func NewService(r Repository) Service {
	// Function to create new accountService by passing an Repository
	return &accountService{r}
}

// s -> Reciver ( work as this keyword)
func (s *accountService) PostAccount(ctx context.Context, name string) (*Account, error) {
	// create an account with name and id
	a := &Account{
		Name: name,
		ID:   ksuid.New().String(), // Generate globally unique identifier (1st 4 byte -> timestamp)
	}
	// pass 'a' to 'PutAccount' func to create an Account
	if err := s.repository.PutAccount(ctx, *a); err != nil {
		return nil, err
	}
	return a, nil
}

func (s *accountService) GetAccount(ctx context.Context, id string) (*Account, error) {
	return s.repository.GetAccountByID(ctx, id)
}

func (s *accountService) GetAccounts(ctx context.Context, skip uint64, take uint64) ([]Account, error) {
	if take > 100 || (skip == 0 && take == 0) {
		take = 100
	}
	return s.repository.ListAccounts(ctx, skip, take)
}
