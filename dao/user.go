package dao

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/iiinsomnia/goadmin/models"

	"github.com/jinzhu/gorm"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/shenghui0779/yiigo"
)

type User struct {
	db  *sqlx.DB
	orm *gorm.DB
}

func NewUser() *User {
	return &User{
		db:  yiigo.DB(),
		orm: yiigo.Orm(),
	}
}

type UserQuery struct {
	Name  string
	Role  int
	Page  int
	Limit int
}

func (u *User) CountList(params *UserQuery) (int64, error) {
	var count int64

	where := make([]string, 0, 2)
	binds := make([]interface{}, 0, 2)

	if params.Name != "" {
		where = append(where, "`name` = ?")
		binds = append(binds, params.Name)
	}

	if params.Role != 0 {
		where = append(where, "`role` = ?")
		binds = append(binds, params.Role)
	}

	query := "SELECT COUNT(*) FROM `go_user`"

	if len(where) > 0 {
		query = fmt.Sprintf("%s WHERE %s", query, strings.Join(where, " AND "))
	} else {
		query += " WHERE 1=1"
	}

	if err := u.db.Get(&count, query, binds...); err != nil {
		return 0, errors.Wrap(err, "count user list error")
	}

	return count, nil
}

func (u *User) FindList(params *UserQuery) ([]*models.User, error) {
	records := make([]*models.User, 0)
	where := make([]string, 0, 2)
	binds := make([]interface{}, 0, 4)

	if params.Name != "" {
		where = append(where, "`name` like "+"'%"+params.Name+"%'")
		// binds = append(binds, params.Name)
	}

	if params.Role != 0 {
		where = append(where, "`role` = ?")
		binds = append(binds, params.Role)
	}

	query := "SELECT `id`, `name`, `email`, `role`, `last_login_ip`, `last_login_time`, `created_at`, `updated_at` FROM `go_user`"

	if len(where) > 0 {
		query = fmt.Sprintf("%s WHERE %s", query, strings.Join(where, " AND "))
	} else {
		query += " WHERE 1=1"
	}

	binds = append(binds, (params.Page-1)*params.Limit, params.Limit)
	query += " LIMIT ?, ?"

	if err := u.db.Select(&records, query, binds...); err != nil {
		return nil, errors.Wrap(err, "find user list error")
	}

	return records, nil
}

// GetUserByID 根据ID获取用户信息
func (u *User) FindByID(id int64) (*models.User, error) {
	record := new(models.User)

	if err := u.db.Get(record, "SELECT `id`, `name`, `email`, `role`, `last_login_ip`, `last_login_time`, `created_at`, `updated_at` FROM `go_user` WHERE `id` = ?", id); err != nil {
		if err != sql.ErrNoRows {
			return nil, errors.Wrap(err, "find user by id error")
		}

		return nil, nil
	}

	return record, nil
}

func (u *User) FindByName(username string) (*models.User, error) {
	record := new(models.User)

	if err := u.db.Get(record, "SELECT `id`, `name`, `email`, `role`, `password`, `salt`, `last_login_ip`, `last_login_time` FROM `go_user` WHERE `name` = ?", username); err != nil {
		if err != sql.ErrNoRows {
			return nil, errors.Wrap(err, "find user by name error")
		}

		return nil, nil
	}

	return record, nil
}

type UserAddData struct {
	Name      string `db:"name"`
	EMail     string `db:"email"`
	Role      int    `db:"role"`
	Password  string `db:"password"`
	Salt      string `db:"salt"`
	CreatedAt int64  `db:"created_at"`
	UpdatedAt int64  `db:"updated_at"`
}

func (u *User) Add(data *UserAddData) error {
	query, binds := yiigo.InsertSQL("go_user", data)

	if _, err := u.db.Exec(query, binds...); err != nil {
		return errors.Wrap(err, "add new user error")
	}

	return nil
}

func (u *User) UpdateByID(id int64, data yiigo.X) error {
	fmt.Println("data: ", data)
	query, binds := yiigo.UpdateSQL("UPDATE `go_user` SET ? WHERE `id` = ?", data, id)

	if _, err := u.db.Exec(query, binds...); err != nil {
		return errors.Wrap(err, "update user by id error")
	}

	return nil
}

func (u *User) DeleteByID(id int64) error {
	if _, err := u.db.Exec("DELETE FROM `go_user` WHERE `id` = ?", id); err != nil {
		return errors.Wrap(err, "delete user by id error")
	}

	return nil
}

func (u *User) FindForUniqueCheck(name string, id ...int64) (*models.User, error) {
	record := new(models.User)
	binds := make([]interface{}, 0, 2)

	query := "SELECT `id` FROM `go_user` WHERE `name` = ?"
	binds = append(binds, name)

	if len(id) > 0 {
		query += " AND `id` <> ?"
		binds = append(binds, id[0])
	}

	query += " LIMIT 1"

	if err := u.db.Get(record, query, binds...); err != nil {
		if err != sql.ErrNoRows {
			return nil, errors.Wrap(err, "find user for unique check error")
		}

		return nil, nil
	}

	return record, nil
}
