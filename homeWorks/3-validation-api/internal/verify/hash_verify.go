package verify

import (
	slicehelpers "3-validation-api/pkg/sliceHelpers"

	"github.com/google/uuid"
)

func (handler VerifyHandler) isHashVerify(hash string) bool {
	err := uuid.Validate(hash)
	if err != nil {
		return false
	}
	if slicehelpers.ContainsStringInStringSlicce(handler.dependens.db, hash) {
		return true
	} else {
		return false
	}

}
