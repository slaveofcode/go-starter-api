package hashids

import (
	hashids "github.com/speps/go-hashids"
	"os"
)

// New will return new instance of HashID generator
func New(minLength int) *hashids.HashID {
	hashIds := hashids.NewData()
	hashIds.Salt = os.Getenv("HASHIDS_SALT")
	hashIds.MinLength = minLength
	inst, _ := hashids.NewWithData(hashIds)
	return inst
}
