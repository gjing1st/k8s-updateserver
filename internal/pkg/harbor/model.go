package harbor

type ArtList struct {
	RepositoryName string `json:"repository_name" form:"repository_name" binding:"required"`
}

// type VersionInfoSum struct {
// 	AllInfo []string `json:"allVersionSum"`
// }

type VersionInfo struct {
	Name    string               `json:"name"`
	Content []VersionInfoContent `json:"content"`
}

type VersionInfoContent struct {
	Title  string   `json:"title"`
	Detail []string `json:"detail"`
}
