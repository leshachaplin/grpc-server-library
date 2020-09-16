package types

type Book struct {
	Name          string `db:"name"`
	Author        string `db:"author"`
	Genre         string `db:"genre"`
	Year          int32  `db:"year"`
}
