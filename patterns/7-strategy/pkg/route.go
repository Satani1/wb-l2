package pkg

import "fmt"

type RoadStrategy struct {
}

func (r *RoadStrategy) Route(startPoint, endPoint int) {
	avgSpeed := 30
	trafficJam := 2
	total := endPoint - startPoint
	totalTime := total * 40 * trafficJam
	fmt.Printf("Road A:[%d] to B:[%d] Avg Speed:[%d] Traffic Jam:[%d], Total:[%d] Total Time:[%d] min\n",
		startPoint, endPoint, avgSpeed, trafficJam, total, totalTime)
}
