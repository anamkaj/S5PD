package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"ps-direct/internal/models"
	"ps-direct/internal/utils"
	"strings"
	"time"
)

type ActiveClientList struct {
	Login     string `json:"login"`
	DateStart string `json:"date_start"`
	DateEnd   string `json:"date_end"`
}

func GetAgencyClients() (*[]models.Client, error) {
	token, err := utils.GetToken()
	if err != nil {
		fmt.Printf("Error: %s", err)
		return nil, err
	}

	url := token.UrlApi

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-yametrika+json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var responseBody []models.Client
	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		return nil, err
	}

	return &responseBody, nil

}

func GetStatApi() (*[]models.CampaignData, error) {
	token, err := utils.GetToken()
	if err != nil {
		fmt.Printf("Error: %s", err)
		return nil, err
	}

	clientList, err := GetAgencyClients()
	if err != nil {
		fmt.Printf("Error: %s", err)
		return nil, err
	}

	url := fmt.Sprintf("https://%s/json/v5/reports", token.UrlApi)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept-Language", "ru")
	req.Header.Set("processingMode", "auto")
	req.Header.Set("returnMoneyInMicros", "false")
	req.Header.Set("skipReportSummary", "true")
	req.Header.Set("IncludeVAT", "true")
	req.Header.Set("Client-Login", "blizko-stroidar")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token.AccessToken)

	activeClients := make([]ActiveClientList, 0)

	for _, c := range *clientList {
		if c.StatusAds == true && c.StatusClient == true && c.DirectLogin != "" {
			activeClients = append(activeClients, ActiveClientList{
				Login:     strings.TrimSpace(c.DirectLogin),
				DateStart: utils.TimeTransform(c.DateStart),
				DateEnd:   utils.TimeTransform(c.DateEnd),
			})
		}

	}

	fmt.Println(activeClients)

	var campaignData []models.CampaignData

	for _, g := range activeClients {
		req.Header.Set("Client-Login", g.Login)
		jsonReq := map[string]interface{}{
			"params": map[string]interface{}{
				"SelectionCriteria": map[string]interface{}{
					"DateFrom": g.DateStart,
					"DateTo":   g.DateEnd,
				},
				"FieldNames": []string{
					"Clicks",
					"Cost",
					"AvgImpressionPosition",
					"AvgTrafficVolume",
					"AvgCpc",
					"AvgPageviews",
					"BounceRate",
					"ClientLogin",
				},
				"ReportName":      fmt.Sprintf("%s"+"%s", time.Now().Format("2006-01-02 15:04:05"), g.Login),
				"ReportType":      "CUSTOM_REPORT",
				"DateRangeType":   "CUSTOM_DATE",
				"Format":          "TSV",
				"IncludeVAT":      "YES",
				"IncludeDiscount": "YES",
			},
		}

		jsonBody, err := json.Marshal(jsonReq)
		if err != nil {
			return nil, err
		}
		req.Body = io.NopCloser(bytes.NewBuffer(jsonBody))
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}

		if resp.StatusCode > 300 {
			fmt.Println(resp.Status)
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			fmt.Println(string(body))
			return nil, err
		}

		if resp.StatusCode == 201 {

			fmt.Println("Created report", g.Login)
			time.Sleep(5 * time.Second)

			req.Body = io.NopCloser(bytes.NewBuffer(jsonBody))

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				return nil, err
			}

			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}

			result, err := utils.TvsTransform(body)
			if err != nil {
				fmt.Printf("Error: %s", err)
				continue
			}

			campaignData = append(campaignData, *result)

		} else {

			fmt.Println("Report already created", g.Login)
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}

			defer resp.Body.Close()

			result, err := utils.TvsTransform(body)
			if err != nil {
				fmt.Printf("Error: %s\n", err)
				continue
			}
			campaignData = append(campaignData, *result)
		}

	}
	return &campaignData, err

}
