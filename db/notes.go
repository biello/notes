package db

// Notes represents a user's note list
type Notes struct {
	Tx    *Tx
	User  []byte
	Notes []Note
}

type Note struct {
	Name    string
	Preview string
}

func (n *Notes) bucket() []byte {
	if len(n.User) == 0 {
		return []byte("pages")
	}
	return n.User
}

// LoadNotes retrieves all note of a user from the database.
func (n *Notes) LoadNotes() error {
	bucket := n.Tx.Bucket(n.bucket())
	if bucket == nil {
		return ErrPageNotFound
	}

	c := n.Tx.Bucket(n.bucket()).Cursor()
	for k, v := c.First(); k != nil; k, v = c.Next() {
		if len(v) <= 100 {
			n.Notes = append(n.Notes, Note{Name: string(k), Preview: string(v)})
		} else {
			n.Notes = append(n.Notes, Note{Name: string(k), Preview: string(v)[0:99] + "..."})
		}
	}

	return nil
}
