package ui

import (
	"fmt"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/schollz/progressbar/v3"
)

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#7B2FBE")).
			Padding(0, 2)

	successStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#04B575")).
			Bold(true)

	errorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF4672")).
			Bold(true)

	labelStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#A0A0A0"))

	valueStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FAFAFA")).
			Bold(true)

	msgBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#7B2FBE")).
			Padding(0, 2).
			MarginTop(1).
			MarginBottom(1)

	doneStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#04B575")).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#04B575")).
			Padding(0, 2)

	subtitleStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#A0A0A0")).
		Italic(true).
		MarginLeft(1)
)

func banner(){
	title :=titleStyle.Render(" AWS Pub/Sub Demo ")
	subTitle := subtitleStyle.Render(" LocalStack • SNS/SQS • Go ")

	fmt.Println()
	fmt.Println(title)
	fmt.Println(subTitle)
	fmt.Println(lipgloss.NewStyle().
		Foreground(lipgloss.Color("#7B2FBE")).
		Render(" ───────────────────────────────────────────────────"))
	fmt.Println()
}

func bar(label string) {
	pb := progressbar.NewOptions(100,
		progressbar.OptionSetDescription(fmt.Sprintf("  %-40s", label)),
		progressbar.OptionSetWidth(30),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "█",
			SaucerPadding: "░",
			BarStart:      "│",
			BarEnd:        "│",
		}),
		progressbar.OptionClearOnFinish(),
		progressbar.OptionSetPredictTime(false),
	)
	for i := 0; i < 100; i++ {
		pb.Add(1)
		time.Sleep(8 * time.Millisecond)
	}
}

func ok(label, value string) {
	fmt.Println(successStyle.Render("  ✓ ") + labelStyle.Render(label+": ") + valueStyle.Render(value))
}

func fail(label, errMsg string) {
	fmt.Println(errorStyle.Render("  ✗ " + label + ": " + errMsg))
}

func MessageBox(direction, message string) {
	fmt.Println(msgBoxStyle.Render("  " + direction + "\n\n  " + message))
}

func Done() {
	fmt.Println()
	fmt.Println(doneStyle.Render(" ✓ SNS → SQS flow completed successfully"))
	fmt.Println()
}

func Spacer() {
	fmt.Println()
}
