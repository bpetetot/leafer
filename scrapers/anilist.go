package scrapers

import (
	"encoding/json"

	"github.com/bpetetot/leafer/db"
	"github.com/go-resty/resty/v2"
)

const anilistURL = "https://graphql.anilist.co"

// https://anilist.gitbook.io/anilist-apiv2-docs

func scrapAnilist(search string) db.Media {
	client := resty.New()

	body, _ := json.Marshal(graphQLBody{
		Query: query,
		Variables: graphQLVariables{
			Search: search,
		},
	})

	resp, _ := client.R().EnableTrace().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(body).
		Post(anilistURL)

	var jsonResult result
	json.Unmarshal(resp.Body(), &jsonResult)

	var media = jsonResult.Data.Media

	return db.Media{
		Title:       media.Title.English,
		TitleNative: media.Title.Native,
		Description: media.Description,
		Status:      media.Status,
		Type:        "MANGA",
		Country:     media.Country,
		CoverImage:  media.CoverImage.Large,
		BannerImage: media.BannerImage,
		Score:       media.Score,
		// Synonyms: media.Synonyms[0]
		// Genre:       media.Genres[0],
		// StartDate:   media.StartDate,
		// EndDate:     media.EndDate,
	}
}

// GraphQL body for Anilist API

type graphQLBody struct {
	Query     string           `json:"query"`
	Variables graphQLVariables `json:"variables"`
}

type graphQLVariables struct {
	Search string `json:"search"`
}

// GraphQL query for Anilist API

const query = `
query ($search: String) {
	Media (search: $search, type: MANGA) {
		title {
			english
			native
			romaji
			userPreferred
		}
		synonyms
		description
		status
		chapters
		volumes
		countryOfOrigin
		coverImage {
			extraLarge
			large
			medium
		}
		bannerImage
		genres
		averageScore
		startDate {
			year
			month
			day
		}
		endDate {
			year
			month
			day
		}
	}
}
`

// Structure of the API response

type result struct {
	Data resultData `json:"data"`
}

type resultData struct {
	Media resultMedia `json:"Media"`
}

type resultMedia struct {
	Title       resultMediaTitle  `json:"title"`
	Synonyms    []string          `json:"synonyms"`
	Description string            `json:"description"`
	Status      string            `json:"status"`
	Chapters    int               `json:"chapters"`
	Volumes     int               `json:"volumes"`
	Country     string            `json:"countryOfOrigin"`
	BannerImage string            `json:"bannerImage"`
	Genres      []string          `json:"genres"`
	Score       int               `json:"averageScore"`
	CoverImage  resultMediaCovers `json:"coverImage"`
	StartDate   resultMediaDate   `json:"startDate"`
	EndDate     resultMediaDate   `json:"endDate"`
}

type resultMediaTitle struct {
	English       string `json:"english"`
	UserPreferred string `json:"userPreferred"`
	Romaji        string `json:"romaji"`
	Native        string `json:"native"`
}

type resultMediaCovers struct {
	ExtraLarge string `json:"extraLarge"`
	Large      string `json:"large"`
	Medium     string `json:"medium"`
}

type resultMediaDate struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"yedayar"`
}
