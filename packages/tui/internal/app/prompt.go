package app

import (
	"time"

	"github.com/moikas-code/kuuzuki-sdk-go"
	"github.com/moikas-code/kuuzuki/internal/attachment"
	"github.com/moikas-code/kuuzuki/internal/compat"
)

type Prompt struct {
	Text        string                   `toml:"text"`
	Attachments []*attachment.Attachment `toml:"attachments"`
}

func (p Prompt) ToMessage(
	messageID string,
	sessionID string,
) Message {
	// Simplified implementation for compatibility
	message := kuuzuki.UserMessage{
		Id:        messageID,
		SessionID: sessionID,
		Role:      compat.UserMessageRoleUser,
		Time: kuuzuki.PermissionInfoTime{
			Created: float32(time.Now().UnixMilli()),
		},
	}

	// For now, just create a simple text part
	// TODO: Implement full attachment support once SDK types are clarified
	if p.Text != "" {
		// This will need to be implemented once we understand the new SDK structure
		// For now, just return the basic message
	}

	return Message{
		Info:  &message,
		Parts: []compat.PartUnion{}, // TODO: Implement parts when SDK structure is clarified
	}
}
