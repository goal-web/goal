package contracts

type Scheduler interface {
	Call()
	Command()
	Exec()
	//Job()
	//DispatchToQueue()
}