package gobject

min_emotion = 0
max_emotion = 10
neutral_emotion = max_emotion / 2

type Emotion struct {
	level float64
	stability float64
}

// Constructor (useful for default settings)
func NewEmotion() Emotion {
	return Emotion{neutral_emotion, 0.005}
}

// Change level by increment
func (e *Emotion) Update(increment float64) {
	if ((e.level + increment) <= min_emotion) {
		e.level = min_emotion
		return
	}
	if ((e.level + increment) >= max_emotion) {
		e.level = max_emotion
	}
	else {
		e.level = e.level + increment
	}
}

// Set emotion level
func (e *Emotion) SetLevel(level float64) {
	if (level >= min_emotion && level <= max_emotion) {
		e.level = level
	}
}

// Set stability (amount between 0 and 1)
// Stability affects rate of neutralization
func (e *Emotion) SetStability(stability float64) {
	if (stability >= 0 && stability <= 1) {
		e.stability = stability
	}
}

// Get if emotion is low
func (e Emotion) IsLow() bool {
	if (e.level < neutral_emotion) {
		return true
	}
	return false
}

// Get if emotion is high
func (e Emotion) IsHigh() bool {
	if (e.level > neutral_emotion) {
		return true
	}
	return false
}

// Get ratio of emotion to max level
func (e Emotion) Ratio() float64 {
	return e.level / max_emotion
}

// Neutralize emotion by an increment of its rate
func (e *Emotion) Neutralize() {
	delta = neutral_emotion - e.level
	e.level = e.level + (delta * (1 - e.stability))
}
