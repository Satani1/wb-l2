package pkg

import "fmt"

type PublicTransportStrategy struct {
}

func (r *PublicTransportStrategy) Route(startPoint, endPoint int) {
	avgSpeed := 40
	total := endPoint - startPoint
	totalTime := total * 40
	fmt.Printf("Public Transport A:[%d] to B:[%d] Avg Speed:[%d] Total:[%d] Total Time:[%d] min\n",
		startPoint, endPoint, avgSpeed, total, totalTime)
}
