package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("\nWelcome!\n======================\nMenu\n======================\nPlease enter one of the following options:\n")
	menu := "\n1: Suggestions based on table\n2: Suggestions based on equation\n3: Suggestions bulk (CSV)\n0: exit\n\n option: "
	fmt.Print(menu)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		opt, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Printf(" Error, Please re enter your option: ")
			continue
		}
		if opt == 0 {
			break
		}
		switch opt {
		case 1:
			for {
				CPM:= initCpmVariables()
				score := suggestedControlByTable(CPM)
				printResult(score,1)
				if confirmation("\n======================\nDone\n======================\nWould you like to back to the menu?[y/n]: ") {
					break
				}
			}
		case 2:
			CPM:= initCpmVariables()
			score := suggestedControlByEquation(CPM)
			printResult(score,1)
			if confirmation("\n======================\nDone\n======================\nWould you like to back to the menu?[y/n]: ") {
				break
			}
		case 3:
			suggestionsFromCSV()
			if confirmation("\n======================\nDone\n======================\nWould you like to back to the menu?[y/n]: ") {
				break
			}
		default:
			fmt.Printf("\n======================\nOption not Found\n======================\nPlease enter one of the following options:\n")
			fmt.Print(menu)
		}
		fmt.Printf("\n======================\nMenu\n======================\nPlease enter one of the following options:\n")
		fmt.Print(menu)
	}
	if err := scanner.Err(); err != nil {
		fmt.Print("Error, Please re enter a integer")
		os.Exit(1)
	}
	fmt.Printf("\n======================\nHASTA LA VISTA!\n======================\nThanks for using the program!\nHave a lovely day!\n")
}
func printResult(score []string,i int) {
	fmt.Printf("Suggested Traffic Control for intersection %d CPM(%s) is %s with a efficiency score of: %s\n",i,score[2], score[1], score[0])
}
func confirmation(msg string) bool {
	fmt.Print(msg)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if strings.EqualFold(scanner.Text(), "y") || strings.EqualFold(scanner.Text(), "yes") {
			return true
		} else if strings.EqualFold(scanner.Text(), "n") || strings.EqualFold(scanner.Text(), "no") {
			return false
		} else {
			fmt.Print("Error, Please re enter : ")
		}
	}
	return false
}
func initCpmVariables() map[string]float64 {
	fmt.Printf("\n======================\nPlease enter the following Cars-Per-Minute (CPM) rates as integers > 0\n\n")
	// Ncpm, Ecpm, Scpm, Wcpm := 11, 5, 21, 24
	Ncpm := readCPM("North")
	Ecpm := readCPM("East")
	Scpm := readCPM("South")
	Wcpm := readCPM("West")

	N, E, S, W := float64(Ncpm), float64(Ecpm), float64(Scpm), float64(Wcpm)
	return map[string]float64{
			"N":N,
			"E":E,
			"S":S,
			"W":W,
		}
}
func readCPM(road string) int {
	fmt.Printf("%s CPM:", road)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		cpm, err := strconv.Atoi(scanner.Text())
		if err == nil {
			fmt.Printf(" %s CPM is %d\n\n", road, cpm)
			return cpm
		} else {
			fmt.Printf(" Error, Please re enter %s as integer > 0\n %s CPM:", road, road)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Print("Error, Please re enter a integer")
	}
	return 0
}

func controlPerformanceTable(cpm float64) (map[string]float64) {
	HT,MT,LT := initPerformanceControlConst()

	var controlPerformance map[string]float64

	switch {
		case cpm < 10:
			controlPerformance = LT
		case cpm < 20:
			controlPerformance = MT
		default:
			controlPerformance = HT
	}
	return controlPerformance
}
func efficiencyScore(CPM map[string]float64) float64{
	var cpmScore float64 = 0.0
	for _, value := range CPM {
		weight := 0.8/getCpmState(value)
		cpmScore += (value*weight)
		fmt.Printf("CPM %f Weight %f Score %f CPMTotal %f \n\n",value,weight,(value*weight),cpmScore)
	}
	return cpmScore
}
func suggestedControlByTable(CPM map[string]float64) []string {
	var scoreMax float64 = 0.0
	var controlMethod string = ""
	var cpmScore float64 = efficiencyScore(CPM)
	controlPerformance := controlPerformanceTable(cpmScore)

	bonus := false
	if((getCpmState(CPM["N"]) > getCpmState(CPM["E"])) && (getCpmState(CPM["N"]) > getCpmState(CPM["W"])) && (getCpmState(CPM["S"]) > getCpmState(CPM["W"])) && (getCpmState(CPM["S"]) > getCpmState(CPM["E"]))) {
		controlPerformance["Roundabout"] = controlPerformance["Roundabout"] + 0.1
		bonus = true
	}
	if(getCpmState(cpmScore) == 2){
		if cpmScore >= 14.5{
			controlPerformance["Traffic Lights"] = controlPerformance["Traffic Lights"] + 0.1
		}
		if cpmScore < 14.5 {
			controlPerformance["Roundabout"] = controlPerformance["Roundabout"] + 0.1

		}
		controlPerformance["Roundabout"] = controlPerformance["Roundabout"] + 0.1
	}

	for key, value := range controlPerformance {
		if value > scoreMax {
			scoreMax = value
			controlMethod = key
		}
	}
	if scoreMax > 1.0 {
		scoreMax = 1.0
	}
	if(controlMethod == "Roundabout" && bonus){
		controlMethod = controlMethod + "(Bonus Applied)"
	}
	scoreString := []string{fmt.Sprintf("%.2f", (scoreMax*100)) + "%", controlMethod,fmt.Sprintf("%.2f", cpmScore)}
	return scoreString
}
func initPerformanceControlConst()(map[string]float64,map[string]float64,map[string]float64){
	HT := map[string]float64{
		"Roundabout":     0.5,
		"Stop Signs":     0.2,
		"Traffic Lights": 0.9,
	}
	MT := map[string]float64{
		"Roundabout":     0.75,
		"Stop Signs":     0.3,
		"Traffic Lights": 0.75,
	}
	LT := map[string]float64{
		"Roundabout":     0.9,
		"Stop Signs":     0.4,
		"Traffic Lights": 0.3,
	}
	return HT,MT,LT
}
func getCpmState(cpm float64) float64{
	state := 0.0
	switch {
		case cpm < 10:
			state = 1
		case cpm < 20:
			state = 2
		default:
			state = 3
	}
	return state
}
func suggestedControlByEquation(CPM map[string]float64) []string {
	var cpmScore float64 = efficiencyScore(CPM)
	var scoreMax float64 = 0.0
	var controlMethod string = ""

	performanceControl := map[string]float64{
		"Roundabout":    getRoundEquation(cpmScore),
		"Stop Signs":     getStopEquation(cpmScore),
		"Traffic Lights": getLightEquation(cpmScore),
	}

	bonus := false
	if((getCpmState(CPM["N"]) > getCpmState(CPM["E"])) && (getCpmState(CPM["N"]) > getCpmState(CPM["W"])) && (getCpmState(CPM["S"]) > getCpmState(CPM["W"])) && (getCpmState(CPM["S"]) > getCpmState(CPM["E"]))) {
		performanceControl["Roundabout"] = performanceControl["Roundabout"] + 0.1
		bonus = true
	}

	for key, value := range performanceControl {
		// fmt.Printf("\ncpm: %f key:%s score:%.2f \n",cpmScore, key, value)
		if value > scoreMax {
			scoreMax = value
			controlMethod = key
		}
	}

	if(controlMethod == "Roundabout" && bonus){
		controlMethod = controlMethod + "(Bonus Applied)"
	}
	scoreString := []string{fmt.Sprintf("%.2f", (scoreMax*100)) + "%", controlMethod,fmt.Sprintf("%.2f", cpmScore)}
	return scoreString
}
func getRoundEquation(cpm float64) float64{
	return (-0.002 * math.Pow(cpm,2))+ (0.02*cpm) + 0.9
}
func getStopEquation(cpm float64) float64{
	return (-0.02*cpm) + 0.6
}
func getLightEquation(cpm float64) float64{
	return (-0.006 * math.Pow(cpm,2))+ (0.24*cpm) - 1.5
}
func suggestionsFromCSV() bool{
	csvFile, err := os.Open("intersections.csv")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\n======================\nSuccessfully Opened CSV file\n======================\n\n")
	defer csvFile.Close()
	csvLines, err := csv.NewReader(csvFile).ReadAll()
    if err != nil {
        fmt.Println(err)
    }
	for i, line := range csvLines {
		N, E, S, W := string2Float(line[0]), string2Float(line[1]), string2Float(line[2]) , string2Float(line[3])
		CPM := map[string]float64{
				"N":N,
				"E":E,
				"S":S,
				"W":W,
		}
		score := suggestedControlByEquation(CPM)
		printResult(score,i+1)
    }
	return false
}
func string2Float(cpm string) float64 {
	number, err := strconv.ParseFloat(cpm, 64)
	if err != nil {
		fmt.Printf("Error in value %s. It occur while reading the csv file: %v",cpm, err)
		os.Exit(1)
	}
	return number
}