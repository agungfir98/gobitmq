package lib

import (
	"encoding/json"
	"io"
)

func Json(w io.Writer, message any) {
	json.NewEncoder(w).Encode(message)
}
