package model

type User struct {
	ID       uint64
	Email    string
	PassHash []byte
}
