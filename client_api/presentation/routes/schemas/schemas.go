package schemas

import "encoding/json"

type Schema func(decoder *json.Decoder) (interface{}, error)
