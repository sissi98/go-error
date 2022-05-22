package dao

import (
	"database/sql"
	"github.com/pkg/errors"
	_ "github.com/go-sql-driver/mysql"
)

type UsersDao struct {
	Table   string      
	Group   string       
	Columns UsersColumns 
}

type UsersColumns struct {
	UserId             string //
	Username           string //
	Nickname           string //
}

var usersColumns = UsersColumns{
	UserId:             "user_id",
	Username:           "username",
	Nickname:           "nickname",
}

var DB *sql.DB

func NewUsersDao() *UsersDao {
	return &UsersDao{
		Group:   "default",
		Table:   "users",
		Columns: usersColumns,
	}
}
func init() {
	var err error
	DB, err = sql.Open("mysql", "mysql:public:123456@tcp(192.168.1.140:3306)/eastblue_new")
	if err != nil {
	   panic(err)
	}
 }

func QueryUserById(id string) (UsersColumns, error) {
	var user UsersColumns
	row := DB.QueryRow("select user_id ,user_name from users where user_id = ?" ,id )
	err := row.Scan(&user.UserId,&user.Username)
	if err != nil{
	   return user,errors.Wrap(err,"dao#QueryUserById err")
	}
	return user,nil
 }