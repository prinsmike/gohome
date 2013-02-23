package user

import (
	"code.google.com/p/go.crypto/bcrypt"
	"labix.org/v2/mgo/bson"
)

type User struct {
	Id       bson.ObjectId `bson:"_id,omitempty"`
	Username string
	Password []byte
	Posts    int
}

func (u *User) SetPassword(password string) {

}

func Login(db *DB, username, password string) (u *User, err error) {

}

func login(w http.ResponseWriter, req *http.Request, db *DB) (err error) {

}

func userExists(username string) bool {

}
