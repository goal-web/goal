package supports

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type User struct {
	Name string
}

func TestClass(t *testing.T) {
	class := GetClass(User{})

	userInstance := class.New(map[string]interface{}{
		"name": "goal",
	}).(User)

	fmt.Println(userInstance)

	assert.True(t, userInstance.Name == "goal")
}
