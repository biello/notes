package db

import "github.com/boltdb/bolt"

// Tx represents a BoltDB transaction
type Tx struct {
	*bolt.Tx
}

// Page retrieves a Page from the database with the given user and page name.
func (tx *Tx) Page(userName, pageName []byte) (*Page, error) {
	p := &Page{
		Tx:   tx,
		User: userName,
		Name: pageName,
	}

	return p, p.Load()
}

// User retrieves a User from the database with the given user name.
func (tx *Tx) User(userName []byte) (*User, error) {
	u := &User{
		Tx:   tx,
		Name: userName,
	}

	return u, u.Load()
}

// Notes list a user's notes from the database.
func (tx *Tx) Notes(userName []byte) (*Notes, error) {
	n := &Notes{
		Tx:   tx,
		User: userName,
	}

	return n, n.LoadNotes()
}
