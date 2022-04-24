package repository

import (
	"context"
	"crypto/md5"
	"database/sql"
	"encoding/hex"

	"github.com/afandi-syaikhu/majoo/model"
)

type User struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &User{
		DB: db,
	}
}

func (_u *User) FindByUsernameAndPassword(ctx context.Context, data model.Auth) (*model.User, error) {
	query := `
			select id, name, user_name, password, created_at, created_by, updated_at, updated_by
			from 
			    users
			where 
			    user_name = $1
				and password = $2
	`
	passHash := md5.Sum([]byte(data.Password))
	passMd5 := hex.EncodeToString(passHash[:])
	rows, err := _u.DB.QueryContext(ctx, query, data.Username, passMd5)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var user *model.User
	for rows.Next() {
		user = &model.User{}
		err = rows.Scan(&user.ID, &user.Name, &user.Username, &user.Password, &user.CreatedAt, &user.CreatedBy, &user.UpdatedAt, &user.UpdatedBy)
		if err != nil {
			return nil, err
		}
	}

	if user == nil {
		return nil, sql.ErrNoRows
	}

	return user, nil
}
