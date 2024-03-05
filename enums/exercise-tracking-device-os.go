package enums

type ExerciseTrackingDeviceOs int

const (
	Android ExerciseTrackingDeviceOs = iota
	IOS
)

type HealthTrackingPoint int

const (
	Walking HealthTrackingPoint = iota
	Sleep
	EnergyBurne
	ExerciseTime
	DistanceWalking
	HeartRate
)
