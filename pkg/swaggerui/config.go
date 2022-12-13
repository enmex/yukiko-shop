package swaggerui

const (
	DefaultStaticRoot = "./static"
	DefaultURLPath    = "spec"
)

type Config struct {
	StaticRoot string `json:"static_root"`
	URLPatch   string `json:"url_patch"`
}
