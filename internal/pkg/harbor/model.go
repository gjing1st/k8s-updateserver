package harbor

type ArtList struct {
	RepositoryName string `json:"repository_name" form:"repository_name" binding:"required"`
}