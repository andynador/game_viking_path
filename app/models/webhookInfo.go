package models

type WebhookInfo struct {
	lastErrorDate int
}

func NewWebhookInfo(lastErrorDate int) WebhookInfo {
	return WebhookInfo{
		lastErrorDate: lastErrorDate,
	}
}

func (webhookInfo WebhookInfo) GetLastErrorDate() int {
	return webhookInfo.lastErrorDate
}
