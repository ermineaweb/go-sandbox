package shortener

type Redirect struct {
	Code      string `json:"code" bson:"code" msgpack:"code"`
	Url       string `json:"url" bson:"url" msgpack:"url" validate:"url"`
	CreatedAt int64  `json:"created_at" bson:"created_at" msgpack:"created_at"`
}
