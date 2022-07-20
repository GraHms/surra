package json

import "encoding/json"

var (
	// Marshal is exported by gin/json package.
	Marshal = json.Marshal
	// Unmarshal is exported by gin/json package.
	Unmarshal = json.Unmarshal

	// NewDecoder is exported by gin/json package.
	NewDecoder = json.NewDecoder
)
