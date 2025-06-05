package verify

import (
	"fmt"

	"github.com/google/uuid"
)

func (handler VerifyHandler) isHashVerify(hash string) bool {
	err := uuid.Validate(hash)
	if err != nil {
		fmt.Println(err)
	}
	return err == nil
}
