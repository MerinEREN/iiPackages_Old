package user

import (
	"errors"
	// "fmt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/user"
	// "github.com/MerinEREN/iiPackages/cookie"
	// "github.com/MerinEREN/iiPackages/role"
	valid "github.com/asaskevich/govalidator"
	"github.com/nu7hatch/gouuid"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	// "io"
	// "log"
	"net/http"
	"time"
)

var (
	EmailNotExist   = errors.New("Email Not Exist")
	ExistingEmail   = errors.New("Existing Email")
	InvalidEmail    = errors.New("Invalid Email")
	InvalidPassword = errors.New("Invalid Password")
	PutUserError    = errors.New("Error while putting user into the datastore.")
	FindUserError   = errors.New("Error while checking email existincy.")
)

// type Users []User

type User struct {
	UUID     string `datastore: "" json:"uuid"`
	Email    string `datastore: "" json:"email"`
	Password string `datastore: "" json:"password"`
	Photo    string `datastore: "" json:"photo"`
	Name     Name   `datastore: "" json:"name"`
	// Phones    Phones `datastore: "" json:"phones"`
	Status       string           `datastore: "" json:"status"`
	Roles        []*datastore.Key `datastore: "" json:"roles"`
	BirthDate    time.Time        `datastore: "" json:"birth_date"`
	Registered   time.Time        `datastore: "" json:"registered"`
	LastModified time.Time        `datastore: "" json:"last_modified"`
	// User could be deactivated by superiors !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	IsActive bool `datastore: "" json:"is_active"`
	// Demands Demands `datastore: "-" json:"demands""`
	// Offers Offers `datastore: "-" json:"offers""`
	// ServicePacks ServicePacks `datastore: "-" json:"service_packs""`
	PurchasedServices []string `datastore: "-" json:"purchased_srvices"`
}

type Name struct {
	First string `datastore: "" json:"first"`
	Last  string `datastore: "" json:"last"`
}

// type Phones []Phone

type Phone string

// type Demands []Demand

type Demand struct {
	ID string `datastore: "" json:"id"`
	// remote or inPlace
	Type        string    `datastore: "" json:"type"`
	Title       string    `datastore: "" json:"title"`
	Description string    `datastore: "" json:"description"`
	StartTime   time.Time `datastore: "" json:"start_time"`
	Duration    string    `datastore: "" json:"duration"`
	Price       Price     `datastore: "" json:"price"`
	// Photos         Photos            `datastore: "-" json:"photos"`
	// Videos         Videos            `datastore: "-" json:"videos"`
	// Tags           Tags              `datastore: "-" json:"tags"`
	Created time.Time `datastore: "" json:"created"`
	// IF THERE IS AT LEAST ONE OFFER DO NOT LET USER TO CHANGE DEMAND !!!!!!!!!!!!!!!!
	LastModified time.Time `datastore: "" json:"last_modified"`
	// underConsideration, active, rejected, changed, removed, finished, disaproved
	Status string `datastore: "" json:"status"`
	// Person In Charge
	Pic string `datastore: "" json:"pic"`
	// IF THERE IS A WAY TO CREATE AN ENTITY WITH TWO ANCESTOR REMOVE Offers PROPERTY !
	Offers []*datastore.Key `datastore:"" json:"offers"`
}

// type Offers []Offer

type Offer struct {
	ID string `datastore: "" json:"id"`
	// remote or inPlace
	// Type        string    `datastore: "" json:"type"` NOT NECESSARY !!!!!!!!!!!!!!!!
	// Title       string    `datastore: "" json:"title"` NOT NECESSARY !!!!!!!!!!!!!!!
	Description string    `datastore: "" json:"description"`
	StartTime   time.Time `datastore: "" json:"start_time"`
	Duration    string    `datastore: "" json:"duration"`
	Price       Price     `datastore: "" json:"price"`
	Created     time.Time `datastore: "" json:"created"`
	// INFORM DEMAND OWNER WHEN AN OFFER MODIFIED !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	LastModified time.Time `datastore: "" json:"last_modified"`
	// available, accepted, notAccepted, changed, removed, successful, unsuccessful
	// backup (ONLY AUTHORIZED ACCOUNTS WHO ACCEPTED TO BE BACKUP) !!!!!!!!!!!!!!!!!!!!
	Status         string     `datastore: "" json:"status"`
	Evaluation     Evaluation `datastore: "" json:"evaluation"`
	CustomerReview string     `datastore: "" json:"customer_review"`
}

// type ServicePacks []ServicePack

type ServicePack struct {
	ID string `datastore: "" json:"id"`
	// remote or inPlace
	Type        string `datastore: "" json:"type"`
	Title       string `datastore: "" json:"title"`
	Description string `datastore: "" json:"description"`
	Duration    string `datastore: "" json:"duration"`
	Price       Price  `datastore: "" json:"price"`
	// Extras         ServicePackExtras `datastore: "-" json:"extras"`
	// Photos         Photos            `datastore: "-" json:"photos"`
	// Videos         Videos            `datastore: "-" json:"videos"`
	// Tags           Tags              `datastore: "-" json:"tags"`
	Created      time.Time `datastore: "" json:"created"`
	LastModified time.Time `datastore: "" json:"last_modified"`
	// underConsideration, disaproved, active, passive, changed, removed
	Status string `datastore: "" json:"status"`
	// Person In Charge
	Pic            string     `datastore: "" json:"pic"`
	Evaluation     Evaluation `datastore: "" json:"evaluation"`
	CustomerReview string     `datastore: "" json:"customer_review"`
}

type Price struct {
	Amount   float64 `datastore: "" json:amount"`
	Currency string  `datastore: "" json:currency"`
}

// type ServicePackExtras []ServicePackOption

type ServicePackOption struct {
	ID          string `datastore: "" json:"id"`
	Title       string `datastore: "" json:"title"`
	Description string `datastore: "" json:"description"`
	Duration    string `datastore: "" json:"duration"`
	Price       Price  `datastore: "" json:"price"`
	// Photos      Photos `datastore: "" json:"photos"`
	// Videos      Videos `datastore: "" json:"videos"`
	Created      time.Time `datastore: "" json:"created"`
	LastModified time.Time `datastore: "" json:"last_modified"`
	// underConsideration, disaproved, active, passive, changed, removed
	Status string `datastore: "" json:"status"`
	// Person In Charge
	Pic string `datastore: "" json:"pic"`
	// Evaluation     Evaluation `datastore: "" json:"evaluation"`
	// CustomerReview string     `datastore: "" json:"customer_review"`
}

// type Photos []Photo

type Photo struct {
	Path         string    `datastore: "" json:"path"`
	Title        string    `datastore: "" json:"title"`
	Description  string    `datastore: "" json:"description"`
	Uploaded     time.Time `datastore: "" json:"uploaded"`
	LastModified time.Time `datastore: "" json:"last_modified"`
	// active or deactive
	Status string `datastore: "" json:"status"`
}

// type Videos []Video

type Video struct {
	Path         string    `datastore: "" json:"path"`
	Title        string    `datastore: "" json:"title"`
	Description  string    `datastore: "" json:"description"`
	Uploaded     time.Time `datastore: "" json:"uploaded"`
	LastModified time.Time `datastore: "" json:"last_modified"`
	// active or deactive
	Status string `datastore: "" json:"status"`
}

// type Tags []Tag

type Tag struct {
	ID     string            `datastore: "" json:"id"`
	Values map[string]string `datastore: "" json:"values"`
}

type Evaluation struct {
	Technical     int
	Timing        int
	Communication int
}

type Entity interface {
	// Use this for all structs
	// Update()
	// Upsert()
	// Delete()
}

func Add(r *http.Request, parentKey *datastore.Key) (u *User, err error) {
	ctx := appengine.NewContext(r)
	var email string
	var roleID string
	var roleKey *datastore.Key
	var roles []*datastore.Key
	if r.Method == "POST" {
		k := "email"
		email = r.PostFormValue(k)
		if !valid.IsEmail(email) {
			err = InvalidEmail
			return
		}
		// FIND A WAY TO GET MULTIPLE ROLES FROM FRONTEND !!!!!!!!!!!!!!!!!!!!!!!!!
		k = "role"
		roleID = r.PostFormValue(k)
		roleKey = datastore.NewKey(ctx, "Roles", roleID, 0, nil)
		roles = append(roles, roleKey)
	} else {
		u1 := user.Current(ctx)
		email = u1.Email
		roleKey = datastore.NewKey(ctx, "Roles", "admin", 0, nil)
		roles = append(roles, roleKey)
	}
	/* k = "password"
	password := r.PostFormValue(k)
	// Cahange this control and allow special characters !!!!!!!!!!!!!!!!!!
	if !valid.IsAlphanumeric(password) {
		err = InvalidPassword
		return
	} */
	u, _, err = Exist(ctx, email)
	if err == datastore.Done {
		u4, errUUID := uuid.NewV4()
		if errUUID != nil {
			err = errUUID
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
		key := datastore.NewKey(ctx, "Users", u.UUID, 0, parentKey)
		_, err = datastore.Put(ctx, key, u)
		if err != nil {
			return
		}
	}
	return
}

func Exist(ctx context.Context, email string) (u *User, key *datastore.Key, err error) {
	q := datastore.NewQuery("Users").Filter("Email =", email)
	for t := q.Run(ctx); ; {
		u = new(User)
		key, err = t.Next(u)
		if err == datastore.Done {
			return
		}
		if err != nil {
			err = FindUserError
			return
		}
		if email == u.Email {
			err = ExistingEmail
			return
		}
	}
}
