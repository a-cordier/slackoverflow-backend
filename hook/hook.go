package hook

const (
	AddQuestion = "add_question"
	AddAnswer   = "add_answer"
)

type ShortCutPayload struct {
	ID         string  `json:"message_ts"`
	Channel    Channel `json:"channel"`
	Message    Message `json:"message"`
	ThreadID   string  `json:"thread_ts"`
	CallbackID string  `json:"callback_id"`
}

type Channel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Message struct {
	Blocks []MessageBlock `json:"blocks"`
}

type MessageBlock struct {
	Type     string                `json:"type"`
	Elements []MessageBlockElement `json:"elements"`
}

type MessageBlockElement struct {
	Type   string                   `json:"type"`
	Chunks []map[string]interface{} `json:"elements"`
}

type TextElement struct {
	Text  string    `json:"text"`
	Style TextStyle `json:"style"`
}

type TextStyle struct {
	Code bool `json:"code"`
}

type EmojiElement struct {
	Name string `json:"name"`
}
