package errno

// ErrPostNotFound 文章未找到.
var ErrPostNotFound = &Errno{HTTP: 404, Code: "ResourceNotFound.PostNotFound", Message: "Post was not found"}
