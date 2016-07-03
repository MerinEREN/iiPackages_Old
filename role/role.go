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
func Put(r *http.Request) error {
	ctx := appengine.NewContext(r)
	key := datastore.NewIncompleteKey(ctx, "Roles", nil)
	role := new(Role)
	role.ID = r.PostFormValue("role")
	role.Values[r.PostFormValue("lang")] = r.PostFormValue("role")
	_, err := datastore.Put(ctx, key, nil)
	return err
}
