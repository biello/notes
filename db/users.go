package db

import "github.com/sirupsen/logrus"

// Errors
var (
	ErrUserNotFound   = &Error{"user not found", nil}
	ErrNoUserName     = &Error{"no user name", nil}
	ErrUserNameExists = &Error{"user name already exist", nil}
	ErrWrongPassword  = &Error{"wrong password", nil}
)

// User represents a user
type User struct {
	Tx       *Tx
	Name     []byte
	Password []byte
}

func (u *User) bucket() []byte {
	return []byte("users")
}

func (u *User) GetPassword() ([]byte, error) {
	bucket := u.Tx.Bucket(u.bucket())
	if bucket == nil {
		return nil, ErrUserNotFound
	}

	password := bucket.Get(u.Name)

	if password == nil {
		return nil, ErrUserNotFound
	}

	return password, nil
}

// Load retrieves a user from the database.
func (u *User) Load() error {
	text, err := u.GetPassword()
	if err != nil {
		return err
	}

	u.Password = text

	return nil
}

// Save commits the User to the database.
func (u *User) Save() error {
	if len(u.Name) == 0 {
		return ErrNoUserName
	}

	logrus.Infof("save user, name: %s, text: %s", u.Name, u.Password)

	return u.Tx.Bucket(u.bucket()).Put(u.Name, u.Password)
}
