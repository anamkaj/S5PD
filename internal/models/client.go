package models

import "time"

type ClientInterface interface{}

type Client struct {
	ID               int64     `json:"id" db:"id"`
	Name             string    `json:"name" db:"name" validate:"required,min=3"`
	URLSite          string    `json:"url_site" db:"url_site" validate:"required,url"`
	DateStart        string    `json:"date_start" db:"date_start" validate:"required,min=3"`
	DateEnd          string    `json:"data_end" db:"data_end" validate:"required,min=3"`
	URLCRM           string    `json:"url_crm" db:"url_crm" validate:"required,url"`
	RegionClient     string    `json:"region_client" db:"region_client"`
	PayCompany       string    `json:"pay_company" db:"pay_company"`
	SpecificClient   string    `json:"specific_client" db:"specific_client"`
	AccountManager   string    `json:"account_manager" db:"account_manager"`
	SpecialistAds    string    `json:"specialist_ads" db:"specialist_ads"`
	StatusAds        bool      `json:"status_ads" db:"status_ads"`
	StatusClient     bool      `json:"status_client" db:"status_client"`
	CountMetrika     int64     `json:"count_metrika" db:"count_metrika"`
	DirectLogin      string    `json:"direct_login" db:"direct_login"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
	Plan             string    `json:"plan" db:"plan"`
	CenterAccounting string    `json:"center_accounting" validate:"required,min=2" db:"center_accounting"`
	PercentageLead   float32   `json:"percentage_lead" db:"percentage_lead"`
	CallTrackingID   int64     `json:"call_tracking_id" db:"call_tracking_id"`
	UniqID           string    `json:"uniq_id" db:"uniq_id"`
}
