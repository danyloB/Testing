package bowling

import (
	"testing"

	"github.com/mdwhatcott/testing/assert"
	"github.com/mdwhatcott/testing/should"
	"github.com/mdwhatcott/testing/suite"
)

func TestGameFixture(t *testing.T) {
	suite.Run(&GameFixture{T: suite.New(t)}, suite.Options.UnitTests())
}

type GameFixture struct {
	*suite.T
	game *Game
}

func (this *GameFixture) Setup() {
	this.game = new(Game)
}
func (this *GameFixture) assertScore(expected int) {
	this.Helper()

	this.So(this.game.CalculateScore(), should.Equal, expected)
	this.So(this.game.CalculateScore(), should.NOT.Equal, expected+1)

	assert.Error(this).So(this.game.CalculateScore(), should.Equal, expected)
	assert.Error(this).So(this.game.CalculateScore(), should.NOT.Equal, expected+1)
}
func (this *GameFixture) rollMany(times, pins int) {
	for x := 0; x < times; x++ {
		this.game.RecordRoll(pins)
	}
}
func (this *GameFixture) rollSeveral(throws ...int) {
	for _, throw := range throws {
		this.game.RecordRoll(throw)
	}
}
func (this *GameFixture) TestGutterGame() {
	this.rollMany(20, 0)
	this.assertScore(0)
}
func (this *GameFixture) TestAllOnes() {
	this.rollMany(20, 1)
	this.assertScore(20)
}
func (this *GameFixture) TestSpare() {
	this.rollSeveral(5, 5, 3)
	this.assertScore(16)
}
func (this *GameFixture) TestStrike() {
	this.rollSeveral(10, 3, 4)
	this.assertScore(24)
}
func (this *GameFixture) TestPerfectGame() {
	this.rollMany(12, 10)
	this.assertScore(300)
}
