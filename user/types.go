package user

import (
	"errors"
	"google.golang.org/appengine/datastore"
	"time"
)

/*
Go's declaration syntax allows grouping of declarations. A single doc comment can introduce
a group of related constants or variables. Since the whole declaration is presented, such
a comment can often be perfunctory.
*/
// Errors
var (
	ErrEmailNotExist = errors.New("Email Not Exist")
	// ErrExistingEmail   = errors.New("Existing Email")
	ErrInvalidEmail    = errors.New("Invalid Email")
	ErrInvalidPassword = errors.New("Invalid Password")
	ErrPutUser         = errors.New("Error while putting user into the datastore.")
	ErrFindUser        = errors.New("Error while checking email existincy.")
)

type Users []User

type User struct {
	UUID  string `datastore: "" json:"uuid"`
	Email string `datastore: "" json:"email"`
	// Password string `datastore: "" json:"password"`
	Photo string `datastore: ",noindex" json:"photo"`
	Name  Name   `datastore: ",noindex" json:"name"`
	// Phones    Phones `datastore: "" json:"phones"`
	// Online, offline, frozen
	Status       string    `datastore: "" json:"status"`
	Type         string    `datastore: "" json:"type"`
	Roles        []string  `datastore: ",noindex" json:"roles"`
	BirthDate    time.Time `datastore: ",noindex" json:"birth_date"`
	Registered   time.Time `datastore: ",noindex" json:"registered"`
	LastModified time.Time `datastore: ",noindex" json:"last_modified"`
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
	Title       string    `datastore: ",noindex" json:"title"`
	Description string    `datastore: ",noindex" json:"description"`
	StartTime   time.Time `datastore: "" json:"start_time"`
	Duration    string    `datastore: ",noindex" json:"duration"`
	Price       Price     `datastore: "" json:"price"`
	// Photos         Photos            `datastore: "-" json:"photos"`
	// Videos         Videos            `datastore: "-" json:"videos"`
	// Tags           Tags              `datastore: "-" json:"tags"`
	Created time.Time `datastore: "" json:"created"`
	// IF THERE IS AT LEAST ONE OFFER DO NOT LET USER TO CHANGE DEMAND !!!!!!!!!!!!!!!!
	LastModified time.Time `datastore: ",noindex" json:"last_modified"`
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
	Description string    `datastore: ",noindex" json:"description"`
	StartTime   time.Time `datastore: ",noindex" json:"start_time"`
	Duration    string    `datastore: ",noindex" json:"duration"`
	Price       Price     `datastore: ",noindex" json:"price"`
	Created     time.Time `datastore: ",noindex" json:"created"`
	// INFORM DEMAND OWNER WHEN AN OFFER MODIFIED !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	LastModified time.Time `datastore: ",noindex" json:"last_modified"`
	// available, accepted, notAccepted, changed, removed, successful, unsuccessful
	// backup (ONLY AUTHORIZED ACCOUNTS WHO ACCEPTED TO BE BACKUP) !!!!!!!!!!!!!!!!!!!!
	Status         string     `datastore: ",noindex" json:"status"`
	Evaluation     Evaluation `datastore: ",noindex" json:"evaluation"`
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
