package contracts

type Exception interface {
	error
	FieldsProvider
}

type ExceptionHandler interface {
	// Handle 处理异常
	Handle(exception Exception)

	// ShouldReport 判断是否需要上报
	ShouldReport(exception Exception) bool

	// Report 上报异常
	Report(exception Exception)
}
