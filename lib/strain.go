package lib

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type Strain struct {
	Url       string
	Streams   int
	Itt       int
	Finishing int
	Method    string
	Listing   []string
}

func (strain *Strain) Get(w http.ResponseWriter, streamsNumber int, ittr int) bool {

	timeStart := time.Now()
	averageTime := 0.0
	maxMin := make([]float64, 2)
	maxMin[1] = 100.12
	countOfHttpErrors := 0
	listOfOll := make([]float64, ittr)

	for i := 0; i < strain.Itt; i++ {
		timeStart = time.Now()
		_, httpErr := http.Get(strain.Url)
		if httpErr != nil {
			countOfHttpErrors += 1
		}
		timeStop := time.Since(timeStart).Seconds()
		averageTime += timeStop
		if maxMin[0] < timeStop {
			maxMin[0] = timeStop
		}
		if maxMin[1] > timeStop {
			maxMin[1] = timeStop
		}
		listOfOll[i] = timeStop
	}
	averageTime = averageTime / float64(strain.Itt)
	listMinMax := make([]float64, 2)
	for _, element := range listOfOll {
		if element >= averageTime {
			listMinMax[0] += 1.0
		}
		if element < averageTime {
			listMinMax[1] += 1.0
		}
	}
	httpMessage := "Iteration: " + strconv.Itoa(strain.Itt) + " | Average Time: " + strconv.FormatFloat(averageTime, 'f', 6, 64)
	httpMessage += " | Max Time: " + strconv.FormatFloat(maxMin[0], 'f', 6, 64)
	httpMessage += " | Min Time: " + strconv.FormatFloat(maxMin[1], 'f', 6, 64)
	httpMessage += " | Min Count: " + strconv.FormatFloat(listMinMax[1], 'f', 6, 64)
	httpMessage += " | Max Count: " + strconv.FormatFloat(listMinMax[0], 'f', 6, 64)
	httpMessage += " | Count of http errors: " + strconv.Itoa(countOfHttpErrors)
	strain.Listing[streamsNumber] = httpMessage
	strain.Finishing += 1

	return true
}

func (strain *Strain) Requests(w http.ResponseWriter) bool {

	if strain.Streams < 1 {
		return false
	}

	streamsInt := strconv.Itoa(strain.Streams)
	strain.Listing = make([]string, strain.Streams)

	for i := 0; i < strain.Streams; i++ {
		switch strain.Method {
		case "GET":
			go strain.Get(w, i, strain.Itt)
			break
		default:
			fmt.Fprintf(w, streamsInt)
		}
	}

	return true
}
