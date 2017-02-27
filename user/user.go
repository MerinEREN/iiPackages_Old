/*
Every package should have a package comment, a block comment preceding the package clause.
For multi-file packages, the package comment only needs to be present in one file, and any
one will do. The package comment should introduce the package and provide information
relevant to the package as a whole. It will appear first on the godoc page and should set
up the detailed documentation that follows.
*/
package user

import (
	// "fmt"
	// "google.golang.org/appengine"
	// "google.golang.org/appengine/user"
	// "github.com/MerinEREN/iiPackages/cookie"
	// "github.com/MerinEREN/iiPackages/role"
	// valid "github.com/asaskevich/govalidator"
	"github.com/nu7hatch/gouuid"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	// "io"
	// "log"
	// "net/http"
	"time"
)

func (u *User) IsAdmin() bool {
	for _, r := range u.Roles {
		if r == "admin" {
			return true
		}
	}
	return false
}

func (u *User) IsContentEditor() bool {
	for _, r := range u.Roles {
		if r == "contentEditor" {
			return true
		}
	}
	return false
}

func New(ctx context.Context, parentKey *datastore.Key, email, role string) (u *User,
	key *datastore.Key, err error) {
	var roles []string
	roles = append(roles, role)
	u, _, err = Get(ctx, email)
	if err == datastore.Done {
		u4 := new(uuid.UUID)
		u4, err = uuid.NewV4()
		if err != nil {
			return
		}
		UUID := u4.String()
		u = &User{
			UUID:  UUID,
			Email: email,
			// Password:     GetHmac(password),
			Roles:        roles,
			Photo:        "adele.jpg",
			Status:       "online",
			IsActive:     true,
			Registered:   time.Now(),
			LastModified: time.Now(),
		}
		key = datastore.NewKey(ctx, "User", u.UUID, 0, parentKey)
		_, err = datastore.Put(ctx, key, u)
		if err != nil {
			return
		}
	}
	return
}

func Get(ctx context.Context, email string) (*User, *datastore.Key, error) {
	u := new(User)
	q := datastore.NewQuery("User").Filter("Email =", email)
	it := q.Run(ctx)
	// BUG !!!!! If i made this function as naked return "it.Next" fails because of "u"
	k, err := it.Next(u)
	if err == datastore.Done {
		return u, k, err
	}
	if err != nil {
		err = ErrFindUser
		return nil, nil, err
	}
	return u, k, nil
}

func GetKey(ctx context.Context, email string) (k *datastore.Key, err error) {
	q := datastore.NewQuery("User").Filter("Email =", email).KeysOnly()
	it := q.Run(ctx)
	k, err = it.Next(nil)
	if err == datastore.Done {
		return
	}
	if err != nil {
		err = ErrFindUser
		return
	}
	return
}

func Exist(ctx context.Context, email string) (c int, err error) {
	c, err = datastore.NewQuery("User").Filter("Email =", email).Count(ctx)
	return
}
