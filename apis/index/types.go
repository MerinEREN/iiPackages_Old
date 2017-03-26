package index

import (
	"github.com/MerinEREN/iiPackages/account"
	"github.com/MerinEREN/iiPackages/user"
)

// Properties has to be kapitalized
// Otherwise they they can't be accessable at the client side.
type userAccount struct {
	User    map[string]*user.User       `json:"user"`
	Account map[string]*account.Account `json:"account"`
}
