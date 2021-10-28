package hashing

import (
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/utils"
)

type Md5 struct {
	salt string
}

func (this *Md5) mixWithSalt(value string) string {
	return value + this.salt
}

func (this *Md5) Info(_ string) contracts.Fields {
	return nil
}

func (this *Md5) Make(value string, _ contracts.Fields) string {
	return utils.Md5(this.mixWithSalt(value))
}

func (this *Md5) Check(value, hashedValue string, _ contracts.Fields) bool {
	return this.Make(value, nil) == hashedValue
}

