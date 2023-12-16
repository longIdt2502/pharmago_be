package b2

import "github.com/kothar/go-backblaze"

func NewB2Bucket(b2AccountId string, b2ApplicationKey string, b2KeyId string, b2Bucket string) (*backblaze.Bucket, error) {
	b2, err := backblaze.NewB2(backblaze.Credentials{
		AccountID:      b2AccountId,
		ApplicationKey: b2ApplicationKey,
		KeyID:          b2KeyId,
	})
	if err != nil {
		return nil, err
	}

	return b2.Bucket(b2Bucket)
}
