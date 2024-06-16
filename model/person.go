package model

import (
	"time"
)

type Person struct {
	ID          string    `json:"id" db:"id"`
	Apelido     string    `json:"apelido" db:"apelido"`
	Nome        string    `json:"nome" db:"nome"`
	Nascimento  time.Time `json:"nascimento" db:"nascimento"`
	Stack       []string  `json:"stack" db:"stack"`
}
