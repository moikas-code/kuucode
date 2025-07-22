package status

import (
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/lipgloss/v2"
	"github.com/charmbracelet/lipgloss/v2/compat"
	"github.com/moikas-code/kuucode/internal/app"
	"github.com/moikas-code/kuucode/internal/commands"
	"github.com/moikas-code/kuucode/internal/styles"
	"github.com/moikas-code/kuucode/internal/theme"
)

type StatusComponent interface {
	tea.Model
	tea.ViewModel
}

type statusComponent struct {
	app   *app.App
	width int
	cwd   string
}

func (m statusComponent) Init() tea.Cmd {
	return nil
}

func (m statusComponent) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		return m, nil
	}
	return m, nil
}

func (m statusComponent) logo() string {
	t := theme.CurrentTheme()
	lightPurple := styles.NewStyle().
		Foreground(compat.AdaptiveColor{
			Dark:  lipgloss.Color("#B19CD9"),
			Light: lipgloss.Color("#8B5A9F"),
		}).
		Background(t.BackgroundElement()).
		Bold(true).
		Render
	version := styles.NewStyle().Foreground(t.TextMuted()).Background(t.BackgroundElement()).Render

	kuu := lightPurple("kuu")
	code := lightPurple("code ")
	versionText := version(m.app.Version)
	return styles.NewStyle().
		Background(t.BackgroundElement()).
		Padding(0, 1).
		Render(kuu + code + versionText)
}

func (m statusComponent) View() string {
	t := theme.CurrentTheme()
	logo := m.logo()

	cwd := styles.NewStyle().
		Foreground(t.TextMuted()).
		Background(t.BackgroundPanel()).
		Padding(0, 1).
		Render(m.cwd)

	var modeBackground compat.AdaptiveColor
	var modeForeground compat.AdaptiveColor
	switch m.app.ModeIndex {
	case 0:
		modeBackground = t.BackgroundElement()
		modeForeground = t.TextMuted()
	case 1:
		modeBackground = t.Secondary()
		modeForeground = t.BackgroundPanel()
	case 2:
		modeBackground = t.Accent()
		modeForeground = t.BackgroundPanel()
	case 3:
		modeBackground = t.Success()
		modeForeground = t.BackgroundPanel()
	case 4:
		modeBackground = t.Warning()
		modeForeground = t.BackgroundPanel()
	case 5:
		modeBackground = t.Primary()
		modeForeground = t.BackgroundPanel()
	case 6:
		modeBackground = t.Error()
		modeForeground = t.BackgroundPanel()
	default:
		modeBackground = t.Secondary()
		modeForeground = t.BackgroundPanel()
	}

	command := m.app.Commands[commands.SwitchModeCommand]
	kb := command.Keybindings[0]
	key := kb.Key
	if kb.RequiresLeader {
		key = m.app.Config.Keybinds.Leader + " " + kb.Key
	}

	modeStyle := styles.NewStyle().Background(modeBackground).Foreground(modeForeground)
	modeNameStyle := modeStyle.Bold(true).Render
	modeDescStyle := modeStyle.Render
	mode := modeNameStyle(strings.ToUpper(m.app.Mode.Name)) + modeDescStyle(" MODE")
	mode = modeStyle.
		Padding(0, 1).
		BorderLeft(true).
		BorderStyle(lipgloss.ThickBorder()).
		BorderForeground(modeBackground).
		BorderBackground(t.BackgroundPanel()).
		Render(mode)

	mode = styles.NewStyle().
		Faint(true).
		Background(t.BackgroundPanel()).
		Foreground(t.TextMuted()).
		Render(key+" ") +
		mode

	space := max(
		0,
		m.width-lipgloss.Width(logo)-lipgloss.Width(cwd)-lipgloss.Width(mode),
	)
	spacer := styles.NewStyle().Background(t.BackgroundPanel()).Width(space).Render("")

	status := logo + cwd + spacer + mode

	blank := styles.NewStyle().Background(t.Background()).Width(m.width).Render("")
	return blank + "\n" + status
}

func NewStatusCmp(app *app.App) StatusComponent {
	statusComponent := &statusComponent{
		app: app,
	}

	homePath, err := os.UserHomeDir()
	cwdPath := app.Info.Path.Cwd
	if err == nil && homePath != "" && strings.HasPrefix(cwdPath, homePath) {
		cwdPath = "~" + cwdPath[len(homePath):]
	}
	statusComponent.cwd = cwdPath

	return statusComponent
}
