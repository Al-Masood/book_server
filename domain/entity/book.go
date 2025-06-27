package entity

type Book struct {
	UUID        string   `json:"uuid" db:"uuid"`
	Name        string   `json:"name" db:"name"`
	AuthorList  []string `json:"authorList" db:"authorList"`
	PublishDate string   `json:"publishDate" db:"publishDate"`
	ISBN        string   `json:"isbn" db:"isbn"`
}
