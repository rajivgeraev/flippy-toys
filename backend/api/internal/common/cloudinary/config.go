package cloudinary

type Config struct {
	CloudName    string
	APIKey       string
	APISecret    string
	UploadPreset string
}

func NewConfig(cloudName, apiKey, apiSecret, uploadPreset string) *Config {
	return &Config{
		CloudName:    cloudName,
		APIKey:       apiKey,
		APISecret:    apiSecret,
		UploadPreset: uploadPreset,
	}
}
