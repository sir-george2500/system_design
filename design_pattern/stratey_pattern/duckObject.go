package main

type Duck struct {
	FlyBehavior FlyBehavior
}

func (d *Duck) performFly() {
	d.FlyBehavior.Fly()
}

// setter of the FlyBehavior
func (d *Duck) setFlyBehavior(fb FlyBehavior) {
	d.FlyBehavior = fb
}
