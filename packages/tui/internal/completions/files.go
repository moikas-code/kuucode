package completions

import (
	"context"
	"log/slog"
	"strings"

	"github.com/moikas-code/kuucode/internal/app"
	"github.com/moikas-code/kuucode/internal/compat"
	"github.com/moikas-code/kuucode/internal/styles"
)

type filesContextGroup struct {
	app      *app.App
	gitFiles []CompletionSuggestion
}

func (cg *filesContextGroup) GetId() string {
	return "files"
}

func (cg *filesContextGroup) GetEmptyMessage() string {
	return "no matching files"
}

func (cg *filesContextGroup) getGitFiles() []CompletionSuggestion {
	items := make([]CompletionSuggestion, 0)

	// TODO: Implement file status via compat layer when available
	// For now, return empty list since the Status method doesn't exist in the compat layer
	return items
}

func (cg *filesContextGroup) GetChildEntries(
	query string,
) ([]CompletionSuggestion, error) {
	items := make([]CompletionSuggestion, 0)

	query = strings.TrimSpace(query)
	if query == "" {
		items = append(items, cg.gitFiles...)
	}

	files, err := cg.app.Client.Find().Files(
		context.Background(),
		compat.FindFilesParams{Query: query},
	)
	if err != nil {
		slog.Error("Failed to get completion items", "error", err)
		return items, err
	}
	if files == nil {
		return items, nil
	}

	for _, file := range *files {
		exists := false
		for _, existing := range cg.gitFiles {
			if existing.Value == file.Path {
				if query != "" {
					items = append(items, existing)
				}
				exists = true
			}
		}
		if !exists {
			displayFunc := func(s styles.Style) string {
				return s.Render(file.Path)
			}

			item := CompletionSuggestion{
				Display:    displayFunc,
				Value:      file.Path,
				ProviderID: cg.GetId(),
				RawData:    file,
			}
			items = append(items, item)
		}
	}

	return items, nil
}

func NewFileContextGroup(app *app.App) CompletionProvider {
	cg := &filesContextGroup{
		app: app,
	}
	go func() {
		cg.gitFiles = cg.getGitFiles()
	}()
	return cg
}
