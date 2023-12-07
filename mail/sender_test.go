package mail

import (
	"github.com/longIdt2502/pharmago_be/utils"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSendEmailWithGmail(t *testing.T) {
	config, err := utils.LoadConfig("..")
	require.NoError(t, err)

	sender := NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)

	subject := "A test email"
	content := `
	<h1>Pharmago xin chào</h1>
	<p>Đây là code</p>
	`
	to := []string{"nhi.nv@idtinc.co", "minh.nv@idtinc.co"}
	attachFiles := []string{}

	err = sender.SendEmail(subject, content, to, nil, nil, attachFiles)
	require.NoError(t, err)
}
