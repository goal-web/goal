package hashing

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/utils"
	"golang.org/x/crypto/bcrypt"
)

type Bcrypt struct {
	cost int
	salt string
}

func (this Bcrypt) mixWithSalt(value string) string {
	return value + this.salt
}

func (this *Bcrypt) Info(hashedValue string) contracts.Fields {
	cost, _ := bcrypt.Cost([]byte(hashedValue))
	return contracts.Fields{
		"cost": cost,
	}
}

func (this *Bcrypt) getCost(options contracts.Fields) int {
	return utils.GetIntField(options, "cost", this.cost)
}

func (this *Bcrypt) Make(value string, options contracts.Fields) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(this.mixWithSalt(value)), this.getCost(options))
	return string(bytes)
}

func (this *Bcrypt) Check(value, hashedValue string, _ contracts.Fields) bool {
	hashedBytes := []byte(hashedValue)
	err := bcrypt.CompareHashAndPassword(hashedBytes, []byte(this.mixWithSalt(value)))
	return err == nil
}
