package account

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

// INTERFACE ->  It is contract which has method declaration .

type Repository interface {
	Close()
	PutAccount(ctx context.Context, a Account) error
	GetAccountByID(ctx context.Context, id string) (*Account, error)
	ListAccounts(ctx context.Context, skip uint64, take uint64) ([]Account, error)
}

// postgresRepository implements Repository bcaz It defines/holds all methods of Repository Interface .

type postgresRepository struct {
	db *sql.DB
}

// This fun invoked in main.go
func NewPostgresRepository(url string) (Repository, error) {
	// Open connection of postgress db with url
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	// Ping the opened connection to check working
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	// Return a postgress repository with the db instance in it
	return &postgresRepository{db}, nil
}

// r is Receiver → declares that this function is a method on postgresRepository. (r work as this keyword in other lang)

func (r *postgresRepository) Close() {
	r.db.Close()
}

func (r *postgresRepository) Ping() error {
	return r.db.Ping()
}

func (r *postgresRepository) PutAccount(ctx context.Context, a Account) error {
	// Create an entry in the db
	_, err := r.db.ExecContext(ctx, "INSERT INTO accounts(id, name) VALUES($1, $2)", a.ID, a.Name)
	return err
}

func (r *postgresRepository) GetAccountByID(ctx context.Context, id string) (*Account, error) {
	// fetch entry from db
	row := r.db.QueryRowContext(ctx, "SELECT id, name FROM accounts WHERE id = $1", id)
	// a is empty Account
	a := &Account{}
	// a is filled by actual values take out from db
	if err := row.Scan(&a.ID, &a.Name); err != nil {
		return nil, err
	}
	return a, nil
}

func (r *postgresRepository) ListAccounts(ctx context.Context, skip uint64, take uint64) ([]Account, error) {
	rows, err := r.db.QueryContext(
		ctx,
		"SELECT id, name FROM accounts ORDER BY id DESC OFFSET $1 LIMIT $2",
		skip,
		take,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// accounts is a list of empty account
	accounts := []Account{}
	for rows.Next() {
		a := &Account{}
		if err = rows.Scan(&a.ID, &a.Name); err == nil {
			accounts = append(accounts, *a)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return accounts, nil
}
