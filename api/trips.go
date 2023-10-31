package api

import (
	"github.com/teleyos/go_ex3_routerank/types"
	"math"
)

const MIN_WAGE_EUR float64 = 9.12

func getScore(trip types.Trip, oneWay bool, green bool) float64 {
	score := trip.PriceEUR
	score += (trip.DurationOutSec - trip.WorkTimeSec) / 3600.0 * MIN_WAGE_EUR
	if green { score += trip.CO2Kg }
	if !(oneWay && !trip.BooReturn) { score += math.Pow(2,20) }

	return math.Round(score*100)/100
}

func ScoreTrips(trips *[]types.Trip) {

	for i:=0; i<len(*trips);i++{
		(*trips)[i].Score = getScore((*trips)[i],true,true)
	}
}
