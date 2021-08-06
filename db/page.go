package db

import "github.com/sirupsen/logrus"

// Errors
var (
	ErrPageNotFound = &Error{"page not found", nil}
	ErrNoPageName   = &Error{"no page name", nil}
)

// Page represents a page
type Page struct {
	Tx   *Tx
	User []byte
	Name []byte
	Text []byte
}

func (p *Page) bucket() []byte {
	if len(p.User) == 0 {
		return []byte("pages")
	}
	return p.User
}

func (p *Page) get() ([]byte, error) {
	bucket := p.Tx.Bucket(p.bucket())
	if bucket == nil {
		return nil, ErrPageNotFound
	}

	text := p.Tx.Bucket(p.bucket()).Get(p.Name)

	if text == nil {
		return nil, ErrPageNotFound
	}

	return text, nil
}

// Load retrieves a page from the database.
func (p *Page) Load() error {
	text, err := p.get()
	if err != nil {
		return err
	}

	p.Text = text

	return nil
}

// Save commits the Page to the database.
func (p *Page) Save() error {
	if len(p.Name) == 0 {
		return ErrNoPageName
	}

	logrus.Infof("save bucket: %s", p.bucket())
	// Create bucket.
	if _, err := p.Tx.CreateBucketIfNotExists(p.bucket()); err != nil {
		return &Error{"user bucket error", err}
	}

	logrus.Infof("save name: %s, text: %s", p.Name, p.Text)

	return p.Tx.Bucket(p.bucket()).Put(p.Name, p.Text)
}
