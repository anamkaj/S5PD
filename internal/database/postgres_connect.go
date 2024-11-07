package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // PostgreSQL driver
	"log"
	"ps-direct/internal/utils"
)

const createStatTable = `
	CREATE TABLE IF NOT EXISTS public.campaign_data (
            id bigserial NOT NULL,
            update_date varchar(50) NOT NULL,
            clicks int8 NOT NULL,
            avg_traffic_volume float8 NOT NULL,
            cost float8 NOT NULL,
            avg_impression_position float8 NOT NULL,
            avg_cpc float8 NOT NULL,
            avg_pageviews float8 NOT NULL,
            bounce_rate float8 NOT NULL,
            client_login varchar(255) NOT NULL,
            CONSTRAINT campaign_data_pkey PRIMARY KEY (id));`

func PostgresConnect() (*sqlx.DB, error) {
	token, err := utils.GetToken()
	if err != nil {
		fmt.Printf("Error: %s", err)
		return nil, err
	}

	pool, err := sqlx.Connect("postgres", token.DirectTable)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = pool.Exec(createStatTable)
	if err != nil {
		return nil, fmt.Errorf("cannot create table: %w", err)
	}

	fmt.Println("Postgres connected")

	return pool, nil
}
