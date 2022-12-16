package main

import (
	"fmt"
	"sort"
)

func main() {

	tariffs := []Tariff{
		Tariff{
			Range:      0000,
			DropRate:   500,
			MinPayment: 18000,
			FlagFall:   7000,
			FixedFare:  0,
			StartTime:  "0000",
			EndTime:    "0000",
		},
		Tariff{
			Range:      5000,
			DropRate:   480,
			MinPayment: 18000,
			FlagFall:   7000,
			FixedFare:  0,
			StartTime:  "0000",
			EndTime:    "0000",
		},
		// Tariff{
		// 	Range:      2000,
		// 	DropRate:   600,
		// 	MinPayment: 18000,
		// 	FlagFall:   7000,
		// 	FixedFare:  0,
		// 	StartTime:  "0000",
		// 	EndTime:    "3000",
		// },
		// Tariff{
		// 	Range:      10000,
		// 	DropRate:   450,
		// 	MinPayment: 18000,
		// 	FlagFall:   7000,
		// 	FixedFare:  0,
		// 	StartTime:  "0000",
		// 	EndTime:    "0000",
		// },
		// Tariff{
		// 	Range:      15000,
		// 	DropRate:   430,
		// 	MinPayment: 18000,
		// 	FlagFall:   7000,
		// 	FixedFare:  0,
		// 	StartTime:  "0000",
		// 	EndTime:    "0000",
		// },
		// Tariff{
		// 	Range:      20000,
		// 	DropRate:   420,
		// 	MinPayment: 18000,
		// 	FlagFall:   7000,
		// 	FixedFare:  0,
		// 	StartTime:  "0000",
		// 	EndTime:    "0000",
		// },
		// Tariff{
		// 	Range:      25000,
		// 	DropRate:   400,
		// 	MinPayment: 18000,
		// 	FlagFall:   7000,
		// 	FixedFare:  0,
		// 	StartTime:  "0000",
		// 	EndTime:    "0000",
		// },
		Tariff{
			Range:      30000,
			DropRate:   480,
			MinPayment: 18000,
			FlagFall:   7000,
			FixedFare:  0,
			StartTime:  "0000",
			EndTime:    "0000",
		},
	}

	fmt.Println("QQQQQ")
	fmt.Println("GEGEGE")

	distance := 11.0
	distanceInMeter := distance * 1000
	tempMapRange := make(map[int32]Tariff)
	tempSliceRange := []int{}
	var actualFare, remainingDistance float64
	isTimedTariff := false

	for _, v := range tariffs {
		if v.StartTime != "0000" || v.EndTime != "0000" {
			isTimedTariff = true
			actualFare = calculateOld(v, distance)
		}

		if !isTimedTariff {
			// range type by distance
			if float64(v.Range) <= distanceInMeter {
				// 5. append v.Range to slice and make v.Range into key for tempMapRange with DropDistance value
				if _, ok := tempMapRange[v.Range]; !ok {
					tempSliceRange = append(tempSliceRange, int(v.Range))
				}
				tempMapRange[v.Range] = Tariff{
					DropRate:   v.DropRate,
					MinPayment: v.MinPayment,
					FixedFare:  v.FixedFare,
					FlagFall:   v.FlagFall,
					StartTime:  v.StartTime,
					EndTime:    v.EndTime,
				}
				// else {
				// 	// check if map with key range has StartTime or EndTime
				// 	// if doesnt have StartTime or EndTime it will be replace with the new one
				// 	if tempMapRange[v.Range].StartTime == "0000" {
				// 		if tempMapRange[v.Range].EndTime == "0000" {
				// tempMapRange[v.Range] = Tariff{
				// 	DropRate:   v.DropRate,
				// 	MinPayment: v.MinPayment,
				// 	FixedFare:  v.FixedFare,
				// 	FlagFall:   v.FlagFall,
				// 	StartTime:  v.StartTime,
				// 	EndTime:    v.EndTime,
				// }
				// }
				// 	}
				// }
			}
		}
	}

	if !isTimedTariff && len(tempSliceRange) > 1 {

		sort.Ints(tempSliceRange)
		fmt.Println(tempSliceRange)
		fmt.Println(tempMapRange)

		for i := 0; i < len(tempSliceRange); i++ {
			// 7. if FixedFare not zero, actualFare = FixedFare
			if tempMapRange[int32(tempSliceRange[i])].FixedFare > 0 {
				actualFare = float64(tempMapRange[int32(tempSliceRange[i])].FixedFare)
				break
			} else {
				// 7. if FixedFare is zero, begin calculation
				if distanceInMeter > 1000 {
					if i+1 == len(tempSliceRange) {
						// When it reach to the last value of tempSliceRange, the drop rate will be multiply by remainingDistance
						// After finish calculate actualFare, it will be compared to min Payment in the last tarif
						fmt.Println(actualFare, "3333")
						fmt.Println(remainingDistance)
						actualFare += float64(float64(tempMapRange[int32(tempSliceRange[i])].DropRate) * (remainingDistance) / 100)

						if actualFare < float64(tempMapRange[int32(tempSliceRange[i])].MinPayment) {
							actualFare = float64(tempMapRange[int32(tempSliceRange[i])].MinPayment)
						}
					} else {
						if remainingDistance == 0 {
							// When customer open the taxi door use flagfall
							// Calculate the first remainingDistance using distanceInMeter
							actualFare += float64(tempMapRange[int32(tempSliceRange[i])].FlagFall)
							if tempSliceRange[i] == 0 {
								remainingDistance = distanceInMeter - float64(tempSliceRange[i+1])
								actualFare += float64(float64(tempMapRange[int32(tempSliceRange[i])].DropRate) * float64(tempSliceRange[i+1]-1000) / 100)
							} else {
								remainingDistance = distanceInMeter - float64(tempSliceRange[i+1])
								actualFare += float64(float64(tempMapRange[int32(tempSliceRange[i])].DropRate) * float64(tempSliceRange[i+1]-tempSliceRange[i]) / 100)
							}

						} else {
							/*
								Value of tempSliceRange is Meter
								DropRate value every 100 meter, so to make the value by km, the result of substraction tempSliceRange divide with 100
								ex : 480 * (2000-1000)/100 = 480 * 10 = 4800 per km
								Calculate the other remaining distance using remainingDistance
							*/

							fmt.Println(actualFare, "1111")
							actualFare += float64(float64(tempMapRange[int32(tempSliceRange[i])].DropRate) * float64(tempSliceRange[i+1]-tempSliceRange[i]) / 100)
							// calculate the remaining distance with current DropRate
							// actualFare += float64(float64(tempMapRange[int32(tempSliceRange[i-1])].DropRate) * 10.0)
							fmt.Println(actualFare, "2222")
							fmt.Println(remainingDistance, " REMAIN")

							remainingDistance = remainingDistance - float64(tempSliceRange[i+1]-tempSliceRange[i])
						}
					}
				} else {
					// if distanceInMeter <=1000, use FlagFall.
					// compare again the actualFare to MinPayment
					actualFare = float64(tempMapRange[int32(tempSliceRange[i])].FlagFall)
					if actualFare < float64(tempMapRange[int32(tempSliceRange[i])].MinPayment) {
						actualFare = float64(tempMapRange[int32(tempSliceRange[i])].MinPayment)
					}
				}
			}
		}
	} else if !isTimedTariff {
		fmt.Println("DEfault")
		actualFare = calculateOld(tempMapRange[0], distance)
	}

	fmt.Println(actualFare, "TOTAL ACTUAL Fare")

}

type Tariff struct {
	Range      int32
	DropRate   int32
	MinPayment int64
	FixedFare  int32
	FlagFall   int32
	StartTime  string
	EndTime    string
}

func calculateOld(v Tariff, distance float64) (actualFare float64) {
	if v.FixedFare > 0 {
		actualFare = float64(v.FixedFare)
	} else {
		fareByDistance := (distance - 1) * float64(v.DropRate*10)
		actualFare = float64(v.FlagFall) + fareByDistance
		if actualFare < float64(v.MinPayment) {
			actualFare = float64(v.MinPayment)
		}
	}
	return
}

// SORT OBJECT
// sort.Slice(tariffData.Tariffs.Tariffs, func(i, j int) bool {
// 	return tariffData.Tariffs.Tariffs[i].Ranges > tariffData.Tariffs.Tariffs[j].Ranges
// })

// input:
// 0 dr 500 ff 7000
// 5 dr 480
// 10 dr 450

// 0 - 5000 = 5000
// 5 - 10000 = 4000
// 10 - 15000 = 3000

// 0 = flagfall 7000  0 7000

// 1 = 7000 	10-5 * droprate / 100
// 2 = 5000 12	5 -> sisa jarak * droprate / 100
// 3 = 5000 17
// 4 = 5000 22
// 5 = 5000 27
// 6 = 4500 31.5
// 7 = 4500 36
// 8 = 4500 40.5
// 9 = 4500 45
// 10 = 4500 49.5
// 11 = 4000 53.5
// 12 = 4000 57.5
// 13 = 4000 61.5
// 14 = 4000 65.5
// 15 = 4000 69.5
