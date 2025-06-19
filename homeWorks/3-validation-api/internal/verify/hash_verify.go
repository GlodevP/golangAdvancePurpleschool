package verify

import (
	"github.com/google/uuid"
)

func (handler VerifyHandler) isHashVerify(hash string) bool {
	err := uuid.Validate(hash)
	if err != nil {
		return false
	}
	email, err := handler.dependens.DB.GetEmailByHash(hash)
	if err != nil {
		return false
	}
	err = handler.dependens.DB.DelHash(email, hash)
	if err != nil {
		return false
	}
	return true
}
