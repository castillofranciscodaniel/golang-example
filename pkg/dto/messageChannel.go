package dto

//MessageChannel -
type MessageChannel struct {
	Id              string `json:"id,omitempty"`
	IdChat          int64  `json:"idChat,omitempty"`
	ChatType        string `json:"chatType,omitempty"`
	Did             string `json:"did"`
	Msisdn          string `json:"msisdn"`
	IdUser          int64  `json:"idUser,omitempty"`
	Type            string `json:"type"`
	Channel         string `json:"channel"`
	ChannelProvider string `json:"channelProvider"`
	Content         string `json:"content"`
	Name            string `json:"name"`
	Campaign        string `json:"idCampaign,omitempty"`
	Group           string `json:"group,omitempty"`
	IsAttachment    bool   `json:"isAttachment"`
	*Attachment     `json:"attachment,omitempty"`
	Message         interface{} `json:"message,omitempty"`
	IdNode          string      `json:"idNode,omitempty"`
	NameNode        string      `json:"nameNode,omitempty"`
	BotName         string      `json:"botName,omitempty"`
	BotId           string      `json:"botId,omitempty"`
	VersionId       string      `json:"versionId,omitempty"`
	ChoiceId        string      `json:"choiceId,omitempty"`
	LastNodeId      string      `json:"lastNodeId,omitempty"`
	BotProvider     string      `json:"botProvider,omitempty"`
	InteractiveMsg  interface{} `json:"interactiveMsg,omitempty"`
	TypificationId  int64       `json:"typificationId,omitempty"`
}

//Attachment -
type Attachment struct {
	MediaUrl string `json:"mediaUrl,omitempty"`
	MimeType string `json:"mimeType"`
	Base64   string `json:"base64,omitempty"`
}
