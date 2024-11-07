package models

type CampaignData struct {
	UpdateDate            string  `json:"update_date" db:"update_date"`
	Clicks                int     `json:"clicks" db:"clicks"`
	Cost                  float64 `json:"cost" db:"cost"`
	AvgImpressionPosition float64 `json:"avg_impression_position" db:"avg_impression_position"`
	AvgTrafficVolume      float64 `json:"avg_traffic_volume" db:"avg_traffic_volume"`
	AvgCpc                float64 `json:"avg_cpc" db:"avg_cpc"`
	AvgPageviews          float64 `json:"avg_pageviews" db:"avg_pageviews"`
	BounceRate            float64 `json:"bounce_rate" db:"bounce_rate"`
	ClientLogin           string  `json:"client_login" db:"client_login"`
}
