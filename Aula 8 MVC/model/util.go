package model

import (
	"time"
)

//Pessoa ...
type Pessoa struct {
	Nome           string
	Idade          int
	DataNascimento time.Time
}
