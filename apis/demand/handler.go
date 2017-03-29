/*
Every package should have a package comment, a block comment preceding the package clause.
For multi-file packages, the package comment only needs to be present in one file, and any
one will do. The package comment should introduce the package and provide information
relevant to the package as a whole. It will appear first on the godoc page and should set
up the detailed documentation that follows.
*/
package demand

import (
	'encoding/json'
	// 'github.com/MerinEREN/iiPackages/account'
	api 'github.com/MerinEREN/iiPackages/apis'
	'github.com/MerinEREN/iiPackages/demand'
	usr 'github.com/MerinEREN/iiPackages/user'
	// 'google.golang.org/appengine'
	'google.golang.org/appengine/datastore'
	'google.golang.org/appengine/user'
	// 'io/ioutil'
	'golang.org/x/net/context'
	'log'
	// 'mime/multipart'
	'net/http'
	// 'regexp'
	// 'time'
)

func Handler(ctx context.Context, w http.ResponseWriter, r *http.Request, ug *user.User) {
	tags, err := usr.GetTags(ctx, ug.Email)
	if err != nil {
		log.Printf('Path: %s, Error: %v\n', r.URL.Path, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resBody := new(api.ResponseBody)
	if tags != nil {
		switch r.FormValue('direction') {
		case 'next':
			demands, cursor, err := demand.GetAllNew(
				r.FormValue('cursor'), 
				tags
			)
	if err != nil {
		log.Printf('Path: %s, Error: %v\n', r.URL.Path, err)
		// Inform client.
	}
		case 'prev':
			demands, cursor, err := demand.GetAllOld(
				r.FormValue('cursor'), 
				tags
			)
	if err != nil {
		log.Printf('Path: %s, Error: %v\n', r.URL.Path, err)
		// Inform client.
	}
	} else {
		resBody.Result = nil
	}
	bs, err := json.Marshal(resBody)
	if err != nil {
		log.Printf('Path: %s, Error: %v\n', r.URL.Path, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(bs)
}
