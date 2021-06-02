package wecom

type Project struct {
	Name   string `json:"name"`
	WebURL string `json:"web_url"`
}

type ProjectMessage struct {
	ObjectKind string   `json:"object_kind"`
	UserName   string   `json:"user_name"`
	UserEmail  string   `json:"user_email"`
	Project    *Project `json:"project"`
}

type TextResponse struct {
	MsgType string               `json:"msgtype"`
	Text    *TextResponseContent `json:"text"`
}

type TextResponseContent struct {
	Content string `json:"content"`
}

func NewTextResponse(content string) *TextResponse {
	return &TextResponse{
		MsgType: "text",
		Text: &TextResponseContent{
			Content: content,
		},
	}
}
