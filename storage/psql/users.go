package psql

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/fn-code/hexagonal-arcitec/storage"
)

func (db *postgresConn) fetchUser(query string, args ...interface{}) ([]storage.UserInfo, error) {
	ctx, cancel := context.WithTimeout(db.ctx, 10*time.Second)
	defer cancel()
	stmt, err := db.conn.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	res := make([]storage.UserInfo, 0)
	for rows.Next() {
		us := storage.UserInfo{}
		err := rows.Scan(&us.IDUser, &us.Name, &us.Email, &us.CreateAt, &us.ChangeAt, &us.LastLogin, &us.Level)
		if err != nil {
			return nil, err
		}
		us.Name = strings.TrimSpace(us.Name)
		us.Email = strings.TrimSpace(us.Email)
		res = append(res, us)
	}
	return res, nil
}

func (db *postgresConn) GetUsers() ([]storage.UserInfo, error) {
	us, err := db.fetchUser("SELECT id_user, name, email, created_at, changed_at, last_login, level FROM users")
	if err != nil {
		return nil, err
	}
	return us, nil
}

func (db *postgresConn) GetUserByID(id string) (*storage.UserInfo, error) {
	us, err := db.fetchUser("SELECT id_user, name, email, created_at, changed_at, last_login, level FROM users WHERE id_user=$1", id)
	if err != nil {
		return nil, err
	}
	rus := storage.UserInfo{}
	if len(us) == 0 {
		return nil, ErrDataNotFound
	}
	rus = us[0]
	return &rus, nil
}

func (db *postgresConn) AddUser(u storage.User, l int8) (string, error) {
	ctx, cancel := context.WithTimeout(db.ctx, 10*time.Second)
	defer cancel()

	stmt, err := db.conn.PrepareContext(ctx, "INSERT INTO users (id_user, name, email, password, salt, papper, created_at, changed_at, last_login, level) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id_user")
	if err != nil {
		return "", err
	}
	var id string
	err = stmt.QueryRowContext(ctx, u.Info.IDUser, u.Info.Name, u.Info.Email, u.Info.Password, u.Salt, u.Papper, u.Info.CreateAt, "-", "-", u.Info.Level).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (db *postgresConn) UpdateUser(id string, u storage.UserInfo) error {
	ctx, cancel := context.WithTimeout(db.ctx, 10*time.Second)
	defer cancel()

	stmt, err := db.conn.PrepareContext(ctx, "UPDATE users SET name=$1, email=$2, changed_at=$3, level=$4 WHERE id_user=$5 RETURNING id_user")
	if err != nil {
		log.Println(err)
		return err
	}
	var ids string
	err = stmt.QueryRowContext(ctx, u.Name, u.Email, u.ChangeAt, u.Level, id).Scan(&ids)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (db *postgresConn) DeleteUser(id string) (string, error) {
	ctx, cancel := context.WithTimeout(db.ctx, 10*time.Second)
	defer cancel()

	stmt, err := db.conn.PrepareContext(ctx, "DELETE FROM users where id_user=$1 RETURNING id_user")
	if err != nil {
		return "", err
	}
	var ids string
	err = stmt.QueryRowContext(ctx, id).Scan(&ids)
	if err != nil {
		return "", err
	}
	return ids, nil
}

func (db *postgresConn) GetUserByUsername(user string) (*storage.User, error) {
	ctx, cancel := context.WithTimeout(db.ctx, 10*time.Second)
	defer cancel()

	stmt, err := db.conn.PrepareContext(ctx, "SELECT id_user, email, name,  level, password, salt, papper FROM users WHERE email=$1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRowContext(ctx, user)
	usr := storage.User{}
	err = row.Scan(&usr.Info.IDUser, &usr.Info.Email, &usr.Info.Name, &usr.Info.Level, &usr.Info.Password, &usr.Salt, &usr.Papper)
	usr.Info.Email = strings.TrimSpace(usr.Info.Email)
	usr.Info.Name = strings.TrimSpace(usr.Info.Name)
	if err != nil {
		return nil, err
	}
	return &usr, nil
}
