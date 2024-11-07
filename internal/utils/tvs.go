package utils

import (
	"fmt"
	"ps-direct/internal/models"
	"strconv"
	"strings"
	"time"
)

func TvsTransform(data []byte) (*models.CampaignData, error) {

	lines := strings.Split(string(data), "\n")

	headers := strings.Split(lines[1], "\t")
	values := strings.Split(lines[2], "\t")

	if len(headers) != len(values) {
		return nil, fmt.Errorf("headers and values have different lengths %s", lines[0])
	}

	campaignData := models.CampaignData{}

	for i, header := range headers {

		trim := strings.TrimSpace(header)
		campaignData.UpdateDate = time.Now().Format("2006-01-02 15:04:05")
		switch trim {
		case "Clicks":
			clicks, err := strconv.Atoi(values[i])
			if err != nil {
				fmt.Printf("Ошибка при парсинге Clicks: %v\n", err)
				continue
			}
			campaignData.Clicks = clicks

		case "Cost":
			cost, err := strconv.ParseFloat(values[i], 64)
			if err != nil {
				fmt.Printf("Ошибка при парсинге Cost: %v\n", err)
				continue
			}
			campaignData.Cost = cost

		case "AvgImpressionPosition":
			avgImpressionPosition, err := strconv.ParseFloat(values[i], 64)
			if err != nil {
				fmt.Printf("Ошибка при парсинге AvgImpressionPosition: %v\n", err)
				continue
			}
			campaignData.AvgImpressionPosition = avgImpressionPosition

		case "AvgTrafficVolume":
			avgTrafficVolume, err := strconv.ParseFloat(values[i], 64)
			if err != nil {
				fmt.Printf("Ошибка при парсинге AvgTrafficVolume: %v\n", err)
				continue
			}
			campaignData.AvgTrafficVolume = avgTrafficVolume

		case "AvgCpc":
			avgCpc, err := strconv.ParseFloat(values[i], 64)
			if err != nil {
				fmt.Printf("Ошибка при парсинге AvgCpc: %v\n", err)
				continue
			}
			campaignData.AvgCpc = avgCpc

		case "AvgPageviews":
			avgPageviews, err := strconv.ParseFloat(values[i], 64)
			if err != nil {
				fmt.Printf("Ошибка при парсинге AvgPageviews: %v\n", err)
				continue
			}
			campaignData.AvgPageviews = avgPageviews

		case "BounceRate":
			bounceRate, err := strconv.ParseFloat(values[i], 64)
			if err != nil {
				fmt.Printf("Ошибка при парсинге BounceRate: %v\n", err)
				continue
			}
			campaignData.BounceRate = bounceRate

		case "ClientLogin":
			campaignData.ClientLogin = values[i]
		}

	}

	return &campaignData, nil

}

func TimeTransform(date string) string {
	parseDate, err := time.Parse("2006.01.02", date)
	if err != nil {
		return date
	}

	return parseDate.Format("2006-01-02")
}
