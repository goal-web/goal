package tests

import (
	"fmt"
	class2 "github.com/goal-web/supports/class"
	"github.com/stretchr/testify/assert"
	"testing"
)

type User struct {
	Name string `json:"nickname"`
}

func TestClass(t *testing.T) {
	class := class2.Make(new(User))

	userInstance := class.New(map[string]interface{}{
		"nickname": "goal",
	}).(User)

	fmt.Println(userInstance)

	assert.True(t, userInstance.Name == "goal")
}
