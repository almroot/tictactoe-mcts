package general

import (
	"fmt"
	"time"
)

type History struct {
	backlog []string
}

func NewHistory() *History {
	return &History{}
}

func (h *History) Ephemeral(player string, message string, callback func(message string)) {
	final := fmt.Sprintf("[%s] <%s> $ %s", time.Now().Format(time.DateTime), player, message)
	callback(final)
}

func (h *History) Report(player string, message string) {
	h.backlog = append(h.backlog, fmt.Sprintf("[%s] <%s> $ %s", time.Now().Format(time.DateTime), player, message))
}

func (h *History) System(message string) {
	h.backlog = append(h.backlog, fmt.Sprintf("[%s]     $ %s", time.Now().Format(time.DateTime), message))
}

func (h *History) Error(err error) {
	h.backlog = append(h.backlog, fmt.Sprintf("[%s] ERR $ %s", time.Now().Format(time.DateTime), err.Error()))
}

func (h *History) List() <-chan string {
	var output = make(chan string)
	go func(output chan string) {
		defer close(output)
		for _, msg := range h.backlog {
			output <- msg
		}
	}(output)
	return output
}
