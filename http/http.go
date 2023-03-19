package http

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type EOLProduct string

const (
	EOLNode   EOLProduct = "nodejs"
	EOLPython EOLProduct = "python"
	EOLDjango EOLProduct = "django"
	EOLReact  EOLProduct = "react"
)

type ProductEOLDetails struct {
	Cycle             string      `json:"cycle"`
	Support           interface{} `json:"support"`
	EOL               interface{} `json:"EOL"`
	Latest            string      `json:"latest"`
	LatestReleaseDate string      `json:"latestReleaseDate"`
	ReleaseDate       string      `json:"releaseDate"`
	LTS               interface{} `json:"LTS"`
}

func EOLProvider(product EOLProduct) ([]ProductEOLDetails, error) {
	resp, err := http.Get("https://endoflife.date/api/" + string(product) + ".json")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("EOL provider did not return a status code of 200")
	}

	b, bErr := io.ReadAll(resp.Body)
	if bErr != nil {
		return nil, bErr
	}
	resp.Body.Close()
	eolDetails := make([]ProductEOLDetails, 0)
	if err := json.Unmarshal(b, &eolDetails); err != nil {
		return nil, err
	}

	return eolDetails, nil

}
