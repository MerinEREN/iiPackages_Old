package role

import (
	"google.golang.org/appengine"
	// "golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"net/http"
)

type Roles []*Role

type Role struct {
	ID     string            `datastore: "" json:"id"`
	Values map[string]string `datastore: "" json:"values"`
}

func Get(r *http.Request, ks []*datastore.Key) (roles Roles, err error) {
	ctx := appengine.NewContext(r)
	err = datastore.GetMulti(ctx, ks, roles)
	return
}

// MAKE THIS MULTI !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
func Put(r *http.Request) (*datastore.Key, error) {
	ctx := appengine.NewContext(r)
	role := new(Role)
	if r.PostFormValue("roleLang") == "en_us" {
		role.ID = r.PostFormValue("roleValue")
		key := datastore.NewKey(ctx, "Roles", role.ID, 0, nil)
	}
	role.Values[r.PostFormValue("roleLang")] =
		r.PostFormValue(r.PostFormValue(roleValue))
	_, err := datastore.Put(ctx, key, role)
	return err
}
