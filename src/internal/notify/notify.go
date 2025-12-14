package notify

import (
	"fmt"

	"github.com/andre-felipe-wonsik-alves/advisor-go/internal/task"
)

func NotifyToTerminal(t task.Task) {
	fmt.Printf("[ATENÇÃO NERD] %s (prioridade: %s)\n", t.Title, t.Priority)

	if t.Description != "" {
		fmt.Printf(" %s\n", t.Description)
	}
}
