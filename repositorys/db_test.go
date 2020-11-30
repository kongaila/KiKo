package repositorys

import (
	"log"
	"testing"
)

func TestConnection(t *testing.T) {
	e := Db.Error
	log.Println(e)
}
