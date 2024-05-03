package b2

import (
	"github.com/kothar/go-backblaze"
	"github.com/longIdt2502/pharmago_be/utils"
)

type B2Bucket struct {
	*backblaze.Bucket
}

func NewB2Bucket(b2AccountId string, b2ApplicationKey string, b2KeyId string, b2Bucket string) (*B2Bucket, error) {
	b2, err := backblaze.NewB2(backblaze.Credentials{
		AccountID:      b2AccountId,
		ApplicationKey: b2ApplicationKey,
		KeyID:          b2KeyId,
	})
	if err != nil {
		return nil, err
	}

	bucket, err := b2.Bucket(b2Bucket)
	if err != nil {
		return nil, err
	}

	return &B2Bucket{
		bucket,
	}, nil
}

func (b2 *B2Bucket) UploadFileToB2(data []byte) (url string, err error) {
	file, _ := utils.NewFileFromImage(data)
	_, err = b2.UploadFile(file.Name, file.Meta, file.File)
	if err != nil {
		return url, err
	}
	url, err = b2.FileURL(file.Name)
	if err != nil {
		return url, err
	}

	return url, nil
}
