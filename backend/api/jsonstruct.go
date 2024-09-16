package api

type Input struct {
	AdCampaignId int    `json:"adCampaignId"`
	CustomerId   int    `json:"customerId"`
	GameName     string `json:"gameName"`
	ImageName    string `json:"imageName"`
	ValidAccount bool   `json:"validAccount"`
}
