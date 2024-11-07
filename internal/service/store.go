package service

import (
	"context"
	"encoding/json"
	"ps-direct/internal/models"
	"time"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type Store struct {
	db    *sqlx.DB
	redis *redis.Client
}

func NewStore(db *sqlx.DB, redis *redis.Client) *Store {
	return &Store{
		db:    db,
		redis: redis,
	}
}

func (s *Store) InsertData(data *[]models.CampaignData) error {

	qInsert := `INSERT INTO campaign_data 
    (update_date, clicks, cost, avg_impression_position, avg_traffic_volume, avg_cpc, avg_pageviews, bounce_rate, client_login) 
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);`

	for _, item := range *data {
		_, err := s.db.Exec(qInsert, item.UpdateDate, item.Clicks, item.Cost, item.AvgImpressionPosition, item.AvgTrafficVolume, item.AvgCpc, item.AvgPageviews, item.BounceRate, item.ClientLogin)
		if err != nil {
			return err
		}
	}

	return nil

}

func (s *Store) InsertStatistics(data *[]models.CampaignData) error {

	for _, item := range *data {

		data, err := json.Marshal(item)
		if err != nil {
			return err
		}
		expiration := 24 * time.Hour

		err = s.redis.Set(context.Background(), item.ClientLogin, data, expiration).Err()
		if err != nil {
			return err
		}

	}

	return nil
}
