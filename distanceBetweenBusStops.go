package main

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	clockwise, counterClockwise := 0, 0
	for i := start; i != destination; i = (i + 1) % len(distance) {
		clockwise += distance[i]
	}
	for i := destination; i != start; i = (i + 1) % len(distance) {
		counterClockwise += distance[i]
	}
	if counterClockwise < clockwise {
		return counterClockwise
	}
	return clockwise
}
