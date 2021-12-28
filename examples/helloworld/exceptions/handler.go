package exceptions

import "github.com/qbhy/goal/contracts"

var DontReportExceptions []contracts.Exception

func init() {
	DontReportExceptions = []contracts.Exception{}
}
