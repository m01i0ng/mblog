package errno

var (
	// ErrPostNotFound 文章未找到
	ErrPostNotFound = &Errno{HTTP: 404, Code: "ResourceNotFound.PostNotFound", Message: "Post was not found"}
)
