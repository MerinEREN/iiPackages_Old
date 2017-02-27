package index

import (
	"github.com/MerinEREN/iiPackages/account"
	usr "github.com/MerinEREN/iiPackages/user"
)

type AccountUser struct {
	Account *account.Account `json:"account"`
	User    *usr.User        `json:"user"`
}
