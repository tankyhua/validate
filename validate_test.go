package validatestruct

import (
	"fmt"
	"testing"
	"time"
)

type TestModel struct {
	CreateTime time.Time `timeval:"required"`
	Name       string
	Age        int
	Money      float64
}

func TestValidate(t *testing.T) {
	var strTime = "2016-01-02 21:00:00"
	var ti, er = time.Parse("2006-01-02 15:04:05", strTime)
	if er != nil {
		return
	}
	var model TestModel
	model.Age = 18
	model.CreateTime = ti
	model.Money = 132.12
	model.Name = "hello"
	
	fmt.Println(Validate(&model))
}
