package firebase

import (
	"github.com/appleboy/go-fcm"
	"github.com/rs/zerolog/log"
)

type FCM struct {
	*fcm.Client
}

func NewFCM(key string) (*FCM, error) {
	client, err := fcm.NewClient(key)
	if err != nil {
		return nil, err
	}
	log.Info().Msg("fcm create new client success")
	return &FCM{client}, nil
}

func (fcmClient *FCM) SendMessage(to string, title string, body string) error {
	msg := &fcm.Message{
		To: to,
		Data: map[string]interface{}{
			"foo": "bar",
		},
		Notification: &fcm.Notification{
			Title: title,
			Body:  body,
		},
	}

	_, err := fcmClient.Send(msg)
	if err != nil {
		log.Error().Msg("failed to send FCM")
		return err
	}

	log.Info().Msg("fcm send message success")

	return nil
}
