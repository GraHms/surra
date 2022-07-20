package TMFramework

type TMFError struct {
	Code    string `json:"code"`
	Reason  string `json:"reason"`
	Message string `json:"message"`
}

var (
	default404Body = []byte("404 page not found")
	default405Body = []byte("405 method not allowed")
)
