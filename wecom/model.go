package wecom

type Project struct {
	Name   string `json:"name"`
	WebURL string `json:"web_url"`
}

type PushMessage struct {
	ObjectKind string   `json:"object_kind"`
	Ref        string   `json:"ref"`
	UserName   string   `json:"user_name"`
	UserEmail  string   `json:"user_email"`
	Project    *Project `json:"project"`
}

type ObjectAttributes struct {
	Status string `json:"status"`
	Ref    string `json:"ref"`
}
type MergeRequest struct {
	Title        string `json:"title"`
	SourceBranch string `json:"source_branch"`
	TargetBranch string `json:"target_branch"`
}
type User struct {
	Name string `json:"name"`
}
type PipelineMessage struct {
	ObjectAttributes *ObjectAttributes `json:"object_attributes"`
	ObjectKind       string            `json:"object_kind"`
	MergeRequest     *MergeRequest     `json:"merge_request"`
	User             *User             `json:"user"`
	Project          *Project          `json:"project"`
}

type MarkdownResponse struct {
	MsgType  string                   `json:"msgtype"`
	Markdown *MarkdownResponseContent `json:"markdown"`
}

type MarkdownResponseContent struct {
	Content string `json:"content"`
}

func NewMarkdownResponse(content string) *MarkdownResponse {
	return &MarkdownResponse{
		MsgType: "markdown",
		Markdown: &MarkdownResponseContent{
			Content: content,
		},
	}
}
