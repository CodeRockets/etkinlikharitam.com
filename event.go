package main

import (
	"time"
)

type Event struct {
	Id                            int       `db:"id" json:"id"`
	VenueId                       int       `db:"venue_id" json:"venue_id"`
	EventName                     string    `db:"event_name" json:"event_name"`
	EventLink                     string    `db:"event_link" json:"event_link"`
	EventImage                    string    `db:"event_image" json:"event_image"`
	EventDescription              string    `db:"event_description" json:"event_description"`
	EventDate                     time.Time `db:"event_date" json:"event_date"`
	EventPrice                    string    `db:"event_price" json:"event_price"`
	EventPurchaseLink             string    `db:"event_purchase_link" json:"event_purchase_link"`
	EventPriceList                []string  `db:"price_list" json:"price_list"`
	EventTagList                  []string  `db:"tag_list" json:"tag_list"`
	VenueName                     string    `db:"venue_name" json:"venue_name"`
	VenuePhone                    string    `db:"venue_phone" json:"venue_phone"`
	VenueTwitter                  string    `db:"venue_twitter" json:"venue_twitter"`
	VenueFacebookId               string    `db:"venue_facebook_id" json:"venue_facebook_id"`
	VenueUserName                 string    `db:"venue_facebook_username" json:"venue_facebook_username"`
	VenueUrl                      string    `db:"venue_url" json:"venue_url"`
	VenueCity                     string    `db:"venue_city" json:"venue_city"`
	VenueCountry                  string    `db:"venue_country" json:"venue_country"`
	VenueFormattedAddress         string    `db:"venue_formatted_address" json:"venue_formatted_address"`
	VenueFacebookUsername         string    `db:"venue_facebook_username" json:"venue_facebook_username"`
	VenueFoursquareTags           string    `db:"venue_foursquare_tags" json:"venue_foursquare_tags"`
	VenueFoursquareRatings        float64   `db:"venue_foursquare_rating" json:"venue_foursquare_rating"`
	VenueFoursquareRatingsSignals int64     `db:"venue_foursquare_rating_signals" json:"venue_foursquare_rating_signals"`
	VenueAddress                  string    `db:"venue_address" json:"venue_address"`
	VenueLat                      float64   `db:"venue_latitude" json:"venue_latitude"`
	VenueLon                      float64   `db:"venue_longitude" json:"venue_longitude"`
	VenueBiletixTag               string    `db:"venue_biletix_tag" json:"venue_biletix_tag"`
	VenueBiletixUrlstring         string    `db:"venue_biletix_url" json:"venue_biletix_url"`
	VenueBiletixName              string    `db:"venue_biletix_name" json:"venue_biletix_name"`
	VenueBiletixImage             string    `db:"venue_biletix_image" json:"venue_biletix_image"`
	VenueBiletixDescription       string    `db:"venue_biletix_description" json:"venue_biletix_description"`
	VenueBiletixDirections        string    `db:"venue_biletix_directions" json:"venue_biletix_directions"`
}
