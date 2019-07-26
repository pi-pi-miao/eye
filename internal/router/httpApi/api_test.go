package httpApi

import (
	"github.com/julienschmidt/httprouter"
	"testing"
)

func TestRegister(t *testing.T) {
	Register(httprouter.New())
}
