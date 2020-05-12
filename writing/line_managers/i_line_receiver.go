package line_managers

type ILineReceiver interface {
	Send(line string)
}
