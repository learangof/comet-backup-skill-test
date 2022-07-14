package main

import (
	"fmt"
)

func main() {

	Ncpm := readCPM("Road 1")
	Ecpm := readCPM("Road 2")
	Scpm := readCPM("Road 3")
	Wcpm := readCPM("Road 4")

	scores := map[string][]string{
		"Road 1": efficiencyScore(Ncpm),
		"Road 2": efficiencyScore(Ecpm),
		"Road 3": efficiencyScore(Scpm),
		"Road 4": efficiencyScore(Wcpm),
	}

	for key, value := range scores {
		fmt.Printf("Suggested Traffic Control for %s is %s with a efficiency score of: %s\n", key, value[1], value[0])
	}
}
func readCPM(road string) int {
	var cpm int;
	fmt.Printf("%s CPM:", road)
	_, err := fmt.Scanf("%d",&cpm)
	fmt.Printf("%s CPM:", err)
	if err != nil {
		for{
			fmt.Printf("%s CPM:", road)
			_, err := fmt.Scanf("%d",&cpm)
			if err != nil {
				break
			}
		}
	}
	return cpm
}
func cpmLevel(cpm int) (map[string]float64, float64) {

	HT := map[string]float64{
		"Roundabout": 0.5,
		"Stop Signs": 0.2,
		"Traffic Lights": 0.9,
	}
	MT := map[string]float64{
		"Roundabout": 0.75,
		"Stop Signs": 0.3,
		"Traffic Lights": 0.75,
	}
	LT := map[string]float64{
		"Roundabout": 0.9,
		"Stop Signs": 0.4,
		"Traffic Lights": 0.3,
	}

	var cpmAvg float64
	var controlPerformance map[string]float64

	switch {
	case cpm < 10:
		controlPerformance = LT
		cpmAvg = 9.0
	case cpm < 20:
		controlPerformance = MT
		cpmAvg = 19.0
	default:
		controlPerformance = HT
		cpmAvg = 20.0
	}
	return controlPerformance, cpmAvg
}
func efficiencyScore(cpm int)[]string{
	var scoreMax float64 = 0.0
	var controlMethod string = ""
	controlPerformance, cpmAvg := cpmLevel(cpm)

	for key, value := range controlPerformance {
		//fmt.Printf("%v : %s -> %.2f\n",cpm, key, float64(cpm)*value/float64(cpmAvg))
		if score := float64(cpm)*value/cpmAvg ; score > scoreMax {
			if cpm >= 25 && key != "Roundabout"{
				scoreMax = score
				controlMethod = key
				continue
			}
			if cpm >= 20 && key != "Traffic Lights"{
				scoreMax = score
				controlMethod = key
				continue
			}
			scoreMax = score
			controlMethod = key
		}
	}
	if scoreMax > 1.0 {
		scoreMax = 1.0
	}
	scoreString := []string{fmt.Sprintf("%.2f", (scoreMax*100))+"%", controlMethod}
	return scoreString
}
