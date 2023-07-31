package history

type History interface {
	Ephemeral(player string, message string, callback func(message string))
	Report(player string, message string)
	System(message string)
	Error(err error)
	List() <-chan string
}
