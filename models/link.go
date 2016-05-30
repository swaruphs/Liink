package models

type Link struct {
	Id           int    `json:"id" schema:"-" redis:"id"`
	WebURL       string `schema:"web_url" json:"web_url" redis:"web_url"`
	AppstoreURL  string `schema:"appstore_url" json:"appstore_url" redis:"appstore_url"`
	PlayStoreURL string `schema:"playstore_url" json:"playstore_url" redis:"playstore_url"`
}
