package d2animdata

// AnimationDataRecord represents a single record from the AnimData.d2 file
type AnimationDataRecord struct {
	name               string
	framesPerDirection uint32
	speed              uint16
	events             map[int]AnimationEvent
}

// FramesPerDirection returns frames per direction value
func (r *AnimationDataRecord) FramesPerDirection() int {
	return int(r.framesPerDirection)
}

// Speed returns animation's speed
func (r *AnimationDataRecord) Speed() int {
	return int(r.speed)
}

// FPS returns the frames per second for this animation record
func (r *AnimationDataRecord) FPS() float64 {
	speedf := float64(r.speed)
	divisorf := float64(speedDivisor)
	basef := float64(speedBaseFPS)

	return basef * speedf / divisorf
}

// FrameDurationMS returns the duration in milliseconds that a frame is displayed
func (r *AnimationDataRecord) FrameDurationMS() float64 {
	return milliseconds / r.FPS()
}
