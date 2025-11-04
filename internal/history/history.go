package history

type History struct {
	commands []string
	capacity int
}

func NewHistory(capacity int) History {
	return History{
		commands: make([]string, 0, capacity),
		capacity: capacity,
	}
}

func (h *History) Append(command string) {
	if command == "" {
		return
	}

	if len(h.commands) > 0 && command == h.commands[len(h.commands)-1] {
		return
	}

	h.commands = append(h.commands, command)

	if len(h.commands) > h.capacity {
		h.commands = h.commands[1:]
	}
}

func (h *History) List() []string {
	return h.commands
}
