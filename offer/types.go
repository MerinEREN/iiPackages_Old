package offer

import (
	"github.com/MerinEREN/iiPackages/price"
	"github.com/MerinEREN/iiPackages/score"
	"time"
)

// INFORM DEMAND OWNER WHEN AN OFFER MODIFIED !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
// Status is available, accepted, notAccepted, changed, removed, successful,
// unsuccessful.
// backup (ONLY AUTHORIZED ACCOUNTS WHO ACCEPTED TO BE BACKUP) !!!!!!!!!!!!!!!!!!!!
// User key is Ancestor
type Offer struct {
	ID             string         `datastore:"-"`
	Description    string         `datastore: ",noindex" json:"description"`
	StartTime      time.Time      `datastore: ",noindex" json:"startTime"`
	Duration       string         `datastore: ",noindex" json:"duration"`
	Price          price.Price    `datastore: ",noindex" json:"price"`
	Created        time.Time      `datastore: ",noindex" json:"created"`
	LastModified   time.Time      `datastore: ",noindex" json:"lastModified"`
	Status         string         `datastore: ",noindex" json:"status"`
	Evaluation     Evaluation     `datastore: ",noindex" json:"evaluation"`
	CustomerReview string         `json:"customerreview"`
	DemandID       *datastore.Key `json:"demandID"`
	Score          score.Score    `json:"score"`
}

type Offers []Offer
