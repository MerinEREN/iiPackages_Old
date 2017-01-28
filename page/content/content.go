package content

import (
	"encoding/json"
	"github.com/MerinEREN/iiPackages/account"
	"golang.org/x/net/context"
	"google.golang.org/appengine/memcache"
	// "github.com/MerinEREN/iiPackages/cookie"
	usr "github.com/MerinEREN/iiPackages/user"
	// "io/ioutil"
	//"net/http"
)

// II Language and page Sturcts
// Languages colection
// STORE ALL OF THOSE IN TO THE DATASTORE !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
type Languages []Language

type Language struct {
	// MAYBE I SHOULD USE []byte INSTEAD OS string TYPE !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	Id    string `datastore:"" json:"id"` // EN, TR ...
	Pages Pages  `datastore:"-" json:"pages"`
}

type Pages []Page

type Page struct {
	C C `datastore:"-" json:"content"`
	D D `datastore:"-" json:"data"`
}

type C struct {
	Title string `datastore:"" json:"title"`
	// Body  Body   `datastore:"-" json:"body"`
}

/* type Body struct {
	Header  Header  `datastore:"-" json:"header"`
	Nav     Nav     `datastore:"-" json:"nav"`
	Partial Partial `datastore:"-" json:"partial"`
	Footer  Footer  `datastore:"-" json:"footer"`
}

type Header struct {
	Logo      Logo      `datastore:"-" json:"logo"`
	Links     Links     `datastore:"Links:"-" json:"links"`
	Dropdowns Dropdowns `datastore:"Dropdowns:"-" json:"dropdowns"`
}

type Logo struct {
	A   A   `datastore:"-" json:"a"`
	Img Img `datastore:"-" json:"logo"`
}

type Img struct {
	Src string `bson:"Img" json:"img"`
	Alt string `bson:"Alt" json:"alt"`
}

type Footer struct {
	// Should be created their own types in the future !!!!!!!!!!!!!!!!!!!!
	SearchPlaceHolder []byte `datastore:"" json:"searchPlaceHolder"`
	MenuButtonText    []byte `datastore:"" json:"menuButtonText"`
} */

type D struct {
	User      *usr.User        `datastore:"-" json:"user"`
	Account   *account.Account `datastore:"-" json:"account"`
	LoginURL  string           `datastore:"-" json:"login_url"`
	LogoutURL string           `datastore:"-" json:"logout_url"`
	URLUUID   string           `datastore:"-" json:"url_uuid"`
}

/* func (p *page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
} */

// GET PAGE CONTENTS FROM DATASTORE WITH USERS SELECTED LANGUAGE !!!!!!!!!!!!!!!!!!!!!!!!!!
func Get(ctx context.Context, title string) (*Page, error) {
	// filename := title + ".html"
	//USE CURRENT WORKING DIRECTORY IN PATH !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	// BECAUSE ioutil.ReadFile USES CALLAR PACKAGE'S DIRECTORY AS CURRENT WORKING
	// DIRECTORY !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	/* body, err := ioutil.ReadFile("../page/templates/" + filename)
	if err != nil {
		return nil, err
	} */
	// IF PAGE ON MEMCACHE GET FROM THERE, OTHERWISE GET FROM DATASTORE =) !!!!!!!!!!!!
	p := new(Page)
	pageItem, err := memcache.Get(ctx, title)
	if err == memcache.ErrCacheMiss {
		content := C{
			Title: title,
		}
		p.C = content
		bs, err1 := json.Marshal(p)
		if err1 != nil {
			return p, err1
		}
		pageItem = &memcache.Item{
			Key:   title,
			Value: bs,
		}
		if err2 := memcache.Set(ctx, pageItem); err2 != nil {
			return p, err2
		}
	} else {
		err = json.Unmarshal(pageItem.Value, p)
		if err != nil {
			return nil, err
		}
	}
	return p, nil
}
