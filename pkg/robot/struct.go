package robot

type DingMsg struct {
	ConversationID            string    `json:"conversationId"`
	AtUsers                   []AtUsers `json:"atUsers"`
	ChatbotUserID             string    `json:"chatbotUserId"`
	MsgID                     string    `json:"msgId"`
	SenderNick                string    `json:"senderNick"`
	IsAdmin                   bool      `json:"isAdmin"`
	SessionWebhookExpiredTime int64     `json:"sessionWebhookExpiredTime"`
	CreateAt                  int64     `json:"createAt"`
	ConversationType          string    `json:"conversationType"`
	SenderID                  string    `json:"senderId"`
	ConversationTitle         string    `json:"conversationTitle"`
	IsInAtList                bool      `json:"isInAtList"`
	SessionWebhook            string    `json:"sessionWebhook"`
	Text                      Text      `json:"text"`
	RobotCode                 string    `json:"robotCode"`
	Msgtype                   string    `json:"msgtype"`
}
type AtUsers struct {
	DingtalkID string `json:"dingtalkId"`
}
type Text struct {
	Content string `json:"content"`
}

type Config struct {
	URL       string `json:"url"`
	Token     string `json:"token"`
	Secret    string `json:"secret"`
	Name      string `json:"name"`
	Identity  string `json:"identity"`
	AuthToken string `json:"auth_token"`
}
