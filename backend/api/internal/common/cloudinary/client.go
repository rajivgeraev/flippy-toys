// internal/common/cloudinary/client.go
package cloudinary

import (
	"context"
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type Client struct {
	cld          *cloudinary.Cloudinary
	config       *Config
	uploadPreset string
}

type UploadParams struct {
	Timestamp int64  `json:"timestamp"`
	Signature string `json:"signature"`
	CloudName string `json:"cloudName"`
	APIKey    string `json:"apiKey"`
	Preset    string `json:"uploadPreset"`
}

func NewClient(config *Config) (*Client, error) {
	cld, err := cloudinary.NewFromParams(config.CloudName, config.APIKey, config.APISecret)
	if err != nil {
		return nil, fmt.Errorf("failed to create cloudinary client: %w", err)
	}

	return &Client{
		cld:          cld,
		config:       config,
		uploadPreset: config.UploadPreset,
	}, nil
}

// GetUploadParams генерирует параметры для безопасной загрузки
func (c *Client) GetUploadParams() (*UploadParams, error) {
	timestamp := time.Now().Unix()

	// Создаем строку для подписи
	signatureStr := fmt.Sprintf("timestamp=%d&upload_preset=%s%s",
		timestamp,
		c.uploadPreset,
		c.config.APISecret)

	// Генерируем подпись
	hash := sha256.New()
	hash.Write([]byte(signatureStr))
	signature := fmt.Sprintf("%x", hash.Sum(nil))

	return &UploadParams{
		Timestamp: timestamp,
		Signature: signature,
		CloudName: c.config.CloudName,
		APIKey:    c.config.APIKey,
		Preset:    c.uploadPreset,
	}, nil
}

// ValidateUploadParams проверяет валидность параметров загрузки
func (c *Client) ValidateUploadParams(timestamp int64, signature string) bool {
	// Проверяем что timestamp не старше часа
	if time.Now().Unix()-timestamp > 3600 {
		return false
	}

	// Воссоздаем строку для подписи
	signatureStr := fmt.Sprintf("timestamp=%d&upload_preset=%s%s",
		timestamp,
		c.uploadPreset,
		c.config.APISecret)

	// Проверяем подпись
	hash := sha256.New()
	hash.Write([]byte(signatureStr))
	expectedSignature := fmt.Sprintf("%x", hash.Sum(nil))

	return signature == expectedSignature
}

func (c *Client) UploadImage(ctx context.Context, data []byte, folder string) (string, string, error) {
	uploadParams := uploader.UploadParams{
		Folder: folder,
		Tags:   []string{"toy"},
	}

	result, err := c.cld.Upload.Upload(ctx, data, uploadParams)
	if err != nil {
		return "", "", err
	}

	return result.SecureURL, result.PublicID, nil
}

func (c *Client) DeleteImage(ctx context.Context, publicID string) error {
	_, err := c.cld.Upload.Destroy(ctx, uploader.DestroyParams{PublicID: publicID})
	return err
}
