package account

import (
	"github.com/MerinEREN/iiPackages/user"
	"time"
)

// type Accounts []Account

type Account struct {
	Photo string `datastore: "" json:"photo"`
	Name  string `datastore: "" json:"name"`
	// Company       Company   `datastore: "-" json:"company"`
	CurrentStatus string `datastore: "" json:"current_status"`
	AccountStatus string `datastore: "" json:"account_status"`
	About         string `datastore: "" json:"about"`
	// Tags          Tags      `datastore: "-" json:"tags"`
	// Ranks         Ranks     `datastore: "-" json:"ranks"`
	// Cards         Cards     `datastore: "-" json:"card" valid:"creditCard"`
	Users        user.Users `datastore: "-" json:"users"`
	Registered   time.Time  `datastore: "" json:"registered"`
	LastModified time.Time  `datastore: "" json:"last_modified"`
}

type Company struct {
	Name string `datastore: "" json:"name"`
	Type string `datastore: "" json:"type"`
	// Addresses Addresses `datastore: "-" json:"addresses"`
}

// type Addresses []Address

type Address struct {
	Description string      `datastore: "" json:"description"`
	Borough     string      `datastore: "" json:"borough"`
	City        string      `datastore: "" json:"city"`
	Country     string      `datastore: "" json:"country"`
	Postcode    string      `datastore: "" json:"postcode"`
	Geolocation Geolocation `datastore: "" json:"geolocation"`
}

type Geolocation struct {
	Lat  string `datastore: "" json:"lat"`  // type could be differnt !!!
	Long string `datastore: "" json:"Long"` // type could be differnt !!!
}

// type Tags []Tag

type Tag struct {
	ID     string            `datastore: "" json:"id"`
	Values map[string]string `datastore: "" json:"values"`
}

// type Ranks []Rank

type Rank struct {
	ID     string            `datastore: "" json:"id"`
	Values map[string]string `datastore: "" json:"values"`
}

// type Cards struct {
// CreditCards CreditCards `datastore: "" json:"creditCards"`
// DebitCards  DebitCards  `datastore: "" json:"debitCards"`
// }

// type CreditCards []CreditCard

type CreditCard struct {
	HolderName string `datastore: "" json:"holder_name"`
	No         string `datastore: "" json:"no"`
	ExpMonth   string `datastore: "" json:"exp_month"`
	ExpYear    string `datastore: "" json:"exp_year"`
	CVV        string `datastore: "" json:"cvv"`
}

// type DebitCards []DebitCard

type DebitCard struct {
	HolderName string `datastore: "" json:"holder_name"`
	No         string `datastore: "" json:"no"`
	ExpMonth   string `datastore: "" json:"exp_month"`
	ExpYear    string `datastore: "" json:"exp_year"`
	CVV        string `datastore: "" json:"cvv"`
}

type Entity interface {
	// Use this for all structs
	// Update()
	// Upsert()
	// Delete()
}
