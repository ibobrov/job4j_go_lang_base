package base

type ValidateRequest struct {
	UserID      string
	Title       string
	Description string
}

func Validate(req *ValidateRequest) []string {
	res := make([]string, 0)
	if req == nil {
		res = append(res, "req is nil")
		return res
	}
	if req.UserID == "" || req.Description == "" || req.Title == "" {
		res = append(res, "value is empty")
	}
	return res
}
