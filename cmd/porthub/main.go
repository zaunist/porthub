package porthub

import (
	"fmt"

	"github.com/zaunist/porthub/pkg/api"
)

func Main() {
	fmt.Println("hello world")
	r := api.SetUpRouter()
	r.Run(":3000")
}
