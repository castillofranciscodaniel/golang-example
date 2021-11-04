package dto

//MessageChannel -
type MessageChannel struct {
	Did        string `json:"did"`
	Msisdn     string `json:"msisdn"`
	IdUser     int    `json:"userId,omitempty"`
	CampaignId int    `json:"campaignId,omitempty"`
}
