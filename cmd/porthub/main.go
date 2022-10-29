package porthub

import (
	"sync"

	"github.com/zaunist/porthub/pkg/api"
)

func Main() {
	r := api.SetUpRouter()
	wg := sync.WaitGroup{}
	wg.Add(3)
	go func() {
		defer func() {
			wg.Done()
		}()
		r.Run(":38210")
	}()
	go func() {
		defer func() {
			wg.Done()
		}()
		r.Run(":38211")
	}()
	go func() {
		defer func() {
			wg.Done()
		}()
		r.Run(":38212")
	}()
	wg.Wait()
}
