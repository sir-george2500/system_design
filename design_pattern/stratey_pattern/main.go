package main

func main() {
	// Create a Duck with a FlyWithWings behavior
	mallard := Duck{FlyBehavior: FlyWithWind{}}
	mallard.performFly()

	// Change the fly behavior to FlyNoWay at runtime
	mallard.setFlyBehavior(FlyWithNoWind{})
	mallard.performFly()
}
