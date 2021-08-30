package bowling

type Game struct {
	score  int
	throw  int
	throws [21]int
}

func (this *Game) RecordRoll(pins int) {
	this.throws[this.throw] = pins
	this.throw++
}

func (this *Game) CalculateScore() int {
	this.throw = 0
	this.score = 0
	for frame := 0; frame < 10; frame++ {
		this.score += this.scoreCurrentFrame()
		this.throw += this.advanceFrame()
	}
	return this.score
}
func (this *Game) scoreCurrentFrame() int {
	if this.currentFrameIsStrike() {
		return this.scoreStrikeFrame()
	} else if this.currentFrameIsSpare() {
		return this.scoreSpareFrame()
	} else {
		return this.scoreRegularFrame()
	}
}
func (this *Game) currentFrameIsStrike() bool {
	return this.pins(0) == 10
}
func (this *Game) currentFrameIsSpare() bool {
	return this.frameScore() == 10
}
func (this *Game) scoreStrikeFrame() int {
	return 10 + this.pins(1) + this.pins(2)
}
func (this *Game) scoreSpareFrame() int {
	return 10 + this.pins(2)
}
func (this *Game) scoreRegularFrame() int {
	return this.frameScore()
}
func (this *Game) frameScore() int {
	return this.pins(0) + this.pins(1)
}
func (this *Game) pins(offset int) int {
	return this.throws[this.throw+offset]
}
func (this *Game) advanceFrame() int {
	if this.currentFrameIsStrike() {
		return 1
	} else {
		return 2
	}
}
