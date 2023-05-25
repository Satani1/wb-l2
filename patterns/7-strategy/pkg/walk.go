package pkg

import "fmt"

type WalkStrategy struct {
}

func (r *WalkStrategy) Route(startPoint, endPoint int) {
	avgSpeed := 4
	total := endPoint - startPoint
	totalTime := total * 60
	fmt.Printf("Walk A:[%d] to B:[%d] Avg Speed:[%d] Total:[%d] Total Time:[%d] min\n",
		startPoint, endPoint, avgSpeed, total, totalTime)
}
