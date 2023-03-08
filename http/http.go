package http

type ProductDetail struct {
	EOLDate string `json:"eol_date"`
	Latest  string `json:"latest"`
	LTS     bool   `json:"lts"`
}

type EOL struct {
	Node   ProductDetail `json:"node"`
	Python ProductDetail `json:"python"`
	Vue    ProductDetail `json:"vue"`
	Django ProductDetail `json:"django"`
	React  ProductDetail `json:"react"`
}

func GetEOLDetails(product string) ProductDetail {
	// make http request here
	// https://endoflife.date/api/{product}.json
	return ProductDetail{
		EOLDate: "2023-06-01",
		Latest:  "19.7.0",
		LTS:     false,
	} // Test with hard coded values
}
