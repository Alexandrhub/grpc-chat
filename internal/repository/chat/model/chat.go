package model

type Chat struct {
	ID   int64    `db:"id"`
	Name []string `db:"username"`
	Msg  string   `db:"msg"`
}
