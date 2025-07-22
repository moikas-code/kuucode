package app

import (
	"time"

	"github.com/moikas-code/kuucode-sdk-go"
	"github.com/moikas-code/kuucode/internal/attachment"
	"github.com/moikas-code/kuucode/internal/id"
)

type Prompt struct {
	Text        string                   `toml:"text"`
	Attachments []*attachment.Attachment `toml:"attachments"`
}

func (p Prompt) ToMessage(
	messageID string,
	sessionID string,
) Message {
	message := kuucode.UserMessage{
		ID:        messageID,
		SessionID: sessionID,
		Role:      kuucode.UserMessageRoleUser,
		Time: kuucode.UserMessageTime{
			Created: float64(time.Now().UnixMilli()),
		},
	}

	text := p.Text
	textAttachments := []*attachment.Attachment{}
	for _, attachment := range p.Attachments {
		if attachment.Type == "text" {
			textAttachments = append(textAttachments, attachment)
		}
	}
	for i := 0; i < len(textAttachments)-1; i++ {
		for j := i + 1; j < len(textAttachments); j++ {
			if textAttachments[i].StartIndex < textAttachments[j].StartIndex {
				textAttachments[i], textAttachments[j] = textAttachments[j], textAttachments[i]
			}
		}
	}
	for _, att := range textAttachments {
		if source, ok := att.GetTextSource(); ok {
			text = text[:att.StartIndex] + source.Value + text[att.EndIndex:]
		}
	}

	parts := []kuucode.PartUnion{kuucode.TextPart{
		ID:        id.Ascending(id.Part),
		MessageID: messageID,
		SessionID: sessionID,
		Type:      kuucode.TextPartTypeText,
		Text:      text,
	}}
	for _, attachment := range p.Attachments {
		text := kuucode.FilePartSourceText{
			Start: int64(attachment.StartIndex),
			End:   int64(attachment.EndIndex),
			Value: attachment.Display,
		}
		source := &kuucode.FilePartSource{}
		switch attachment.Type {
		case "text":
			continue
		case "file":
			if fileSource, ok := attachment.GetFileSource(); ok {
				source = &kuucode.FilePartSource{
					Text: text,
					Path: fileSource.Path,
					Type: kuucode.FilePartSourceTypeFile,
				}
			}
		case "symbol":
			if symbolSource, ok := attachment.GetSymbolSource(); ok {
				source = &kuucode.FilePartSource{
					Text: text,
					Path: symbolSource.Path,
					Type: kuucode.FilePartSourceTypeSymbol,
					Kind: int64(symbolSource.Kind),
					Name: symbolSource.Name,
					Range: kuucode.SymbolSourceRange{
						Start: kuucode.SymbolSourceRangeStart{
							Line:      float64(symbolSource.Range.Start.Line),
							Character: float64(symbolSource.Range.Start.Char),
						},
						End: kuucode.SymbolSourceRangeEnd{
							Line:      float64(symbolSource.Range.End.Line),
							Character: float64(symbolSource.Range.End.Char),
						},
					},
				}
			}
		}
		parts = append(parts, kuucode.FilePart{
			ID:        id.Ascending(id.Part),
			MessageID: messageID,
			SessionID: sessionID,
			Type:      kuucode.FilePartTypeFile,
			Filename:  attachment.Filename,
			Mime:      attachment.MediaType,
			URL:       attachment.URL,
			Source:    *source,
		})
	}
	return Message{
		Info:  message,
		Parts: parts,
	}
}

func (m Message) ToSessionChatParams() []kuucode.SessionChatParamsPartUnion {
	parts := []kuucode.SessionChatParamsPartUnion{}
	for _, part := range m.Parts {
		switch p := part.(type) {
		case kuucode.TextPart:
			parts = append(parts, kuucode.TextPartInputParam{
				ID:        kuucode.F(p.ID),
				Type:      kuucode.F(kuucode.TextPartInputTypeText),
				Text:      kuucode.F(p.Text),
				Synthetic: kuucode.F(p.Synthetic),
				Time: kuucode.F(kuucode.TextPartInputTimeParam{
					Start: kuucode.F(p.Time.Start),
					End:   kuucode.F(p.Time.End),
				}),
			})
		case kuucode.FilePart:
			var source kuucode.FilePartSourceUnionParam
			switch p.Source.Type {
			case "file":
				source = kuucode.FileSourceParam{
					Type: kuucode.F(kuucode.FileSourceTypeFile),
					Path: kuucode.F(p.Source.Path),
					Text: kuucode.F(kuucode.FilePartSourceTextParam{
						Start: kuucode.F(int64(p.Source.Text.Start)),
						End:   kuucode.F(int64(p.Source.Text.End)),
						Value: kuucode.F(p.Source.Text.Value),
					}),
				}
			case "symbol":
				source = kuucode.SymbolSourceParam{
					Type: kuucode.F(kuucode.SymbolSourceTypeSymbol),
					Path: kuucode.F(p.Source.Path),
					Name: kuucode.F(p.Source.Name),
					Kind: kuucode.F(p.Source.Kind),
					Range: kuucode.F(kuucode.SymbolSourceRangeParam{
						Start: kuucode.F(kuucode.SymbolSourceRangeStartParam{
							Line:      kuucode.F(float64(p.Source.Range.(kuucode.SymbolSourceRange).Start.Line)),
							Character: kuucode.F(float64(p.Source.Range.(kuucode.SymbolSourceRange).Start.Character)),
						}),
						End: kuucode.F(kuucode.SymbolSourceRangeEndParam{
							Line:      kuucode.F(float64(p.Source.Range.(kuucode.SymbolSourceRange).End.Line)),
							Character: kuucode.F(float64(p.Source.Range.(kuucode.SymbolSourceRange).End.Character)),
						}),
					}),
					Text: kuucode.F(kuucode.FilePartSourceTextParam{
						Value: kuucode.F(p.Source.Text.Value),
						Start: kuucode.F(p.Source.Text.Start),
						End:   kuucode.F(p.Source.Text.End),
					}),
				}
			}
			parts = append(parts, kuucode.FilePartInputParam{
				ID:       kuucode.F(p.ID),
				Type:     kuucode.F(kuucode.FilePartInputTypeFile),
				Mime:     kuucode.F(p.Mime),
				URL:      kuucode.F(p.URL),
				Filename: kuucode.F(p.Filename),
				Source:   kuucode.F(source),
			})
		}
	}
	return parts
}

func (p Prompt) ToSessionChatParams() []kuucode.SessionChatParamsPartUnion {
	parts := []kuucode.SessionChatParamsPartUnion{
		kuucode.TextPartInputParam{
			Type: kuucode.F(kuucode.TextPartInputTypeText),
			Text: kuucode.F(p.Text),
		},
	}
	for _, att := range p.Attachments {
		filePart := kuucode.FilePartInputParam{
			Type:     kuucode.F(kuucode.FilePartInputTypeFile),
			Mime:     kuucode.F(att.MediaType),
			URL:      kuucode.F(att.URL),
			Filename: kuucode.F(att.Filename),
		}
		switch att.Type {
		case "file":
			if fs, ok := att.GetFileSource(); ok {
				filePart.Source = kuucode.F(
					kuucode.FilePartSourceUnionParam(kuucode.FileSourceParam{
						Type: kuucode.F(kuucode.FileSourceTypeFile),
						Path: kuucode.F(fs.Path),
						Text: kuucode.F(kuucode.FilePartSourceTextParam{
							Start: kuucode.F(int64(att.StartIndex)),
							End:   kuucode.F(int64(att.EndIndex)),
							Value: kuucode.F(att.Display),
						}),
					}),
				)
			}
		case "symbol":
			if ss, ok := att.GetSymbolSource(); ok {
				filePart.Source = kuucode.F(
					kuucode.FilePartSourceUnionParam(kuucode.SymbolSourceParam{
						Type: kuucode.F(kuucode.SymbolSourceTypeSymbol),
						Path: kuucode.F(ss.Path),
						Name: kuucode.F(ss.Name),
						Kind: kuucode.F(int64(ss.Kind)),
						Range: kuucode.F(kuucode.SymbolSourceRangeParam{
							Start: kuucode.F(kuucode.SymbolSourceRangeStartParam{
								Line:      kuucode.F(float64(ss.Range.Start.Line)),
								Character: kuucode.F(float64(ss.Range.Start.Char)),
							}),
							End: kuucode.F(kuucode.SymbolSourceRangeEndParam{
								Line:      kuucode.F(float64(ss.Range.End.Line)),
								Character: kuucode.F(float64(ss.Range.End.Char)),
							}),
						}),
						Text: kuucode.F(kuucode.FilePartSourceTextParam{
							Start: kuucode.F(int64(att.StartIndex)),
							End:   kuucode.F(int64(att.EndIndex)),
							Value: kuucode.F(att.Display),
						}),
					}),
				)
			}
		}
		parts = append(parts, filePart)
	}
	return parts
}
