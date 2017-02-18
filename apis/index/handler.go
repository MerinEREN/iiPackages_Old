/*
Every package should have a package comment, a block comment preceding the package clause.
For multi-file packages, the package comment only needs to be present in one file, and any
one will do. The package comment should introduce the package and provide information
relevant to the package as a whole. It will appear first on the godoc page and should set
up the detailed documentation that follows.
*/
package index

import (
	"encoding/json"
	// "fmt"
	"github.com/MerinEREN/iiPackages/account"
	api "github.com/MerinEREN/iiPackages/apis"
	"github.com/MerinEREN/iiPackages/cookie"
	// "github.com/MerinEREN/iiPackages/page/content"
	// "github.com/MerinEREN/iiPackages/page/template"
	usr "github.com/MerinEREN/iiPackages/user"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/memcache"
	"google.golang.org/appengine/user"
	"io/ioutil"
	// "html/template"
	"log"
	// "mime/multipart"
	"net/http"
	// "regexp"
	// "time"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/favicon.ico" {
		return
	}
	switch r.Method {
	case "POST":
		// remove ctx
		ctx := appengine.NewContext(r)
		wb := new(api.ResponseBody)
		// ug is google user
		ug := user.Current(ctx)
		// Login or get data needed
		if ug == nil {
			gURL, err := user.LoginURL(ctx, r.URL.String())
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			loginURLs := make(map[string]string)
			loginURLs["Google"] = gURL
			loginURLs["LinkedIn"] = gURL
			loginURLs["Twitter"] = gURL
			loginURLs["Facebook"] = gURL
			wb.Data = loginURLs
		} else {
			// Allways close the body
			defer r.Body.Close()
			// r.Body is io.ReadCloser type, so may be closing request body
			// explicitly is not necessary.
			rb, err := ioutil.ReadAll(r.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			var rbum map[string]map[string]interface{}
			err = json.Unmarshal(rb, &rbum)
			if err != nil {
				log.Println("Error while unmarshalling request body:", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			acc := new(account.Account)
			u := new(usr.User)
			// Users own account page or not
			if accName, ok := rbum["data"]["acc"]; !ok {
				aKey := new(datastore.Key)
				uKey := new(datastore.Key)
				item, err := memcache.Get(ctx, "u")
				if err == nil {
					err = json.Unmarshal(item.Value, u)
					if err != nil {
						log.Printf("Path: %s, Error: %v\n",
							r.URL.Path, err)
						http.Error(w, err.Error(),
							http.StatusInternalServerError)
						return
					}
				} else {
					u, uKey, err = usr.Get(ctx, ug.Email)
					switch err {
					case datastore.Done:
						acc, u, uKey, err = account.Create(ctx)
						if err != nil {
							log.Printf("Error while creating "+
								"account: %v\n", err)
							// ALSO LOG THIS WITH DATASTORE LOG
							http.Error(w, err.Error(),
								http.StatusInternalServerError)
							return
						}
						bs, err := json.Marshal(acc)
						if err != nil {
							log.Printf("Path: %s, Error: %v\n",
								r.URL.Path, err)
						}
						item = &memcache.Item{
							Key:   "acc",
							Value: bs,
						}
						err = memcache.Add(ctx, item)
						if err != nil {
							log.Printf("Path: %s, Error: %v\n",
								r.URL.Path, err)
						}
					case usr.ErrFindUser:
						log.Printf("Error while login user: %v\n",
							err)
						// ALSO LOG THIS WITH DATASTORE LOG !!!!!!!
						http.Error(w, err.Error(),
							http.StatusInternalServerError)
						return
					}
					bs, err := json.Marshal(u)
					if err != nil {
						log.Printf("Path: %s, Error: %v\n",
							r.URL.Path, err)
					}
					bsUKey, err := json.Marshal(uKey)
					if err != nil {
						log.Printf("Path: %s, Error: %v\n",
							r.URL.Path, err)
					}
					items := []*memcache.Item{
						{
							Key:   "u",
							Value: bs,
						},
						{
							Key:   "uKey",
							Value: bsUKey,
						},
					}
					err = memcache.AddMulti(ctx, items)
					if err != nil {
						log.Printf("Path: %s, Error: %v\n",
							r.URL.Path, err)
					}
				}
				item, err = memcache.Get(ctx, "acc")
				if err == nil {
					err = json.Unmarshal(item.Value, acc)
					if err != nil {
						log.Printf("Path: %s, Error: %v\n",
							r.URL.Path, err)
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}
				} else {
					item, err = memcache.Get(ctx, "uKey")
					if err == nil {
						err = json.Unmarshal(item.Value, uKey)
						if err != nil {
							log.Printf("Path: %s, Error: %v\n",
								r.URL.Path, err)
							http.Error(w, err.Error(), http.StatusInternalServerError)
							return
						}
						aKey = uKey.Parent()
						err = datastore.Get(ctx, aKey, acc)
						if err != nil {
							log.Printf("Error while getting user's "+
								"account data: %v\n", err)
							// ALSO LOG THIS WITH DATASTORE LOG !!!!!!!
							http.Error(w, err.Error(),
								http.StatusInternalServerError)
							return
						}
					} else {
						uKey, err = usr.GetKey(ctx, ug.Email)
						switch err {
						// Imposible case
						/* case datastore.Done:
						acc, u, uKey, err = account.Create(ctx)
						if err != nil {
							log.Printf("Error while creating "+
								"account: %v\n", err)
							// ALSO LOG THIS WITH DATASTORE LOG
							http.Error(w, err.Error(),
								http.StatusInternalServerError)
							return
						}
						bs, err := json.Marshal(acc)
						if err != nil {
							log.Printf("Path: %s, Error: %v\n",
								s, err)
						}
						itemA = &memcache.Item{
							Key:   "acc",
							Value: bs,
						}
						err = memcache.Add(ctx, itemA)
						if err != nil {
							log.Printf("Path: %s, Error: %v\n",
								s, err)
						} */
						case usr.ErrFindUser:
							log.Printf("Error while login user: %v\n",
								err)
							// ALSO LOG THIS WITH DATASTORE LOG !!!!!!!
							http.Error(w, err.Error(),
								http.StatusInternalServerError)
							return
						default:
							aKey = uKey.Parent()
							err = datastore.Get(ctx, aKey, acc)
							if err != nil {
								log.Printf("Error while getting user's "+
									"account data: %v\n", err)
								// ALSO LOG THIS WITH DATASTORE LOG !!!!!!!
								http.Error(w, err.Error(),
									http.StatusInternalServerError)
								return
							}
						}
						bs, err := json.Marshal(uKey)
						if err != nil {
							log.Printf("Path: %s, Error: %v\n",
								r.URL.Path, err)
						}
						item = &memcache.Item{
							Key:   "uKey",
							Value: bs,
						}
						err = memcache.Add(ctx, item)
						if err != nil {
							log.Printf("Path: %s, Error: %v\n",
								r.URL.Path, err)
						}
					}
					bs, err := json.Marshal(acc)
					if err != nil {
						log.Printf("Path: %s, Error: %v\n",
							r.URL.Path, err)
					}
					item = &memcache.Item{
						Key:   "acc",
						Value: bs,
					}
					err = memcache.Add(ctx, item)
					if err != nil {
						log.Printf("Path: %s, Error: %v\n",
							r.URL.Path, err)
					}
				}
				err = cookie.Set(w, r, u.UUID)
				if err != nil {
					log.Printf("Error while creating session "+
						"cookie: %v\n", err)
				}
			} else {
				// someone elses account
				s, ok := accName.(string)
				if ok {
					acc, err = account.Get(ctx, s)
				} else {
					log.Println("Account name type is not string.")
					http.Error(w, "Account name type is not string.",
						http.StatusBadRequest)
					return
				}
			}
			au := struct {
				Account *account.Account `json:"account"`
				User    *usr.User        `json:"user"`
			}{acc, u}
			wb.Data = au
		}
		/* t := &http.Transport{}
		t.RegisterProtocol("file", http.NewFileTransport(http.Dir("/")))
		c := &http.Client{Transport: t}
		res, err := c.Get("file:///etc/passwd")
		log.Println(res, err) */
		// To respond to request without any data
		// w.WriteHeader(StatusOK)
		// Always send corresponding header values instead of defaults !!!!
		//w.Header().Set("Content-Type", "application/json; charset=utf-8")
		wbm, err := json.Marshal(*wb)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(wbm)
		// http.NotFound(w, r)
		// http.Redirect(w, r, "/MerinEREN", http.StatusFound)
	default:
		// IF DOSN'T HAVE ACCOUNT URL PARAM
		// ELSE GET ACCOUNT AND ALL USERS OF THAT ACCOUNT INFO
	}

	/* temp := template.Must(template.New("fdsfdfdf").Parse(pBody))
	err = temp.Execute(w, p)
	if err != nil {
		log.Print(err)
	} */
	// THE IF CONTROL BELOW IS IMPORTANT
	// WHEN PAGE LOADS THERE IS NO FILE SELECTED AND THIS CAUSE A PROBLEM FOR
	/* if r.Method == "POST" {
		var f multipart.File
		key := "uploadedFile"
		f, _, err := r.FormFile(key)
		if err != nil {
			fmt.Println("File input is empty.")
			return
		}
		defer f.Close()
		var bs []byte
		bs, err = ioutil.ReadAll(f)
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, "File: %s\n Error: %v\n", string(bs), err)
	} */
}
