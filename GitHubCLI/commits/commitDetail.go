package commits

type CommitDetail struct {
	Author  CommitAuthor `json:"author"`
	Message string       `json:"message"`
}
