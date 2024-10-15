package main

import "fmt"

const (
	TerroristDressType        = "tDress"
	CounterTerroristDressType = "ctDress"
)

var dressFactorySingleInstance = &DressFactory{
	dressMap: make(map[string]Dress),
}

// 享元工厂
type DressFactory struct {
	dressMap map[string]Dress
}

func (d *DressFactory) getDressByType(dressType string) (Dress, error) {
	if d.dressMap[dressType] != nil {
		return d.dressMap[dressType], nil
	}
	if dressType == TerroristDressType {
		d.dressMap[dressType] = newTerroristDress()
		return d.dressMap[dressType], nil
	}
	if dressType == CounterTerroristDressType {
		d.dressMap[dressType] = newCounterTerroristDress()
		return d.dressMap[dressType], nil
	}
	return nil, fmt.Errorf("Wrong dress type passed")
}

func getDressFactorySingleInstance() *DressFactory {
	return dressFactorySingleInstance
}

// 享元接口
type Dress interface {
	getColor() string
}

// 具体享元对象
type TerroristDress struct {
	color string
}

func (t *TerroristDress) getColor() string {
	return t.color
}

func newTerroristDress() *TerroristDress {
	return &TerroristDress{color: "red"}
}

// 具体享元对象
type CounterTerroristDress struct {
	color string
}

func (c *CounterTerroristDress) getColor() string {
	return c.color
}

func newCounterTerroristDress() *CounterTerroristDress {
	return &CounterTerroristDress{color: "blue"}
}

type Player struct {
	dress      Dress
	playerType string
	lat        int
	long       int
}

// 背景
func newPlayer(playerType, dressType string) *Player {
	dress, _ := getDressFactorySingleInstance().getDressByType(dressType)
	return &Player{playerType: playerType, dress: dress}
}

func (p *Player) newLocation(lat, long int) {
	p.lat = lat
	p.long = long
}

type game struct {
	terrorists        []*Player
	CounterTerrorists []*Player
}

// 开始一盘游戏
func newGame() *game {
	return &game{
		terrorists:        make([]*Player, 1),
		CounterTerrorists: make([]*Player, 1),
	}
}

// 添加玩家
func (c *game) addTerrorist(dressType string) {
	// 选取阵营
	player := newPlayer("T", dressType)
	c.terrorists = append(c.terrorists, player)
	return
}

// 添加玩家
func (c *game) addCounterTerrorist(dressType string) {
	player := newPlayer("CT", dressType)
	c.CounterTerrorists = append(c.CounterTerrorists, player)
	return
}

func main() {
	game := newGame()
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)

	game.addCounterTerrorist(CounterTerroristDressType)
	game.addCounterTerrorist(CounterTerroristDressType)
	game.addCounterTerrorist(CounterTerroristDressType)

	dressFactoryInstance := getDressFactorySingleInstance()
	for dressType, dress := range dressFactoryInstance.dressMap {
		fmt.Printf("DressColorType: %s\nDressColor: %s\n", dressType, dress.getColor())
	}
}
