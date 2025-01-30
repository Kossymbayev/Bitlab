/*package main

import (
	"fmt"
	"time"

	"github.com/rivo/tview"
)


func main() {
	clear()

	var playerName string

	typewriter("Введи свое имя: ")
	fmt.Scan(&playerName)
	pause(1)

	typewriter(fmt.Sprintf("%v проиграл", playerName))

	app := tview.NewApplication()

	menu := tview.NewList().
		AddItem("Кнопка 1", "", '1', nil).
		AddItem("Кнопка 2", "", '2', nil).
		AddItem("Выход", "Закрыть приложение", 'q', func() {
			app.Stop()
		})

	if err := app.SetRoot(menu, true).Run(); err != nil {
		panic(err)
	}
}

*/

package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/rivo/tview"
)

type Character struct {
	Name      string
	HP        int
	Defence   int
	Strength  int
	Weapon    int
	Shield    int
	HasShield bool
	Potions   int
}

var dragon = Character{
	Name:     "Дракон",
	HP:       2000,
	Defence:  120,
	Strength: 150,
	Weapon:   0,
}

var hero = Character{
	Name:      "Герой",
	HP:        1000,
	Defence:   100,
	Strength:  120,
	Weapon:    250,
	Shield:    150,
	HasShield: false,
	Potions:   1,
}

func modifyHealth(target *Character, amount int) {
	target.HP += amount
	if target.HP < 0 {
		target.HP = 0
	}
}

func displayInfo(c Character) {
	typewriter(fmt.Sprintf("%s: HP = %d, Defence = %d\n", c.Name, c.HP, c.Defence))
}

func heroTurn() {
	typewriter("Ход героя! Время выбрать действие")
	pause(3)

	app := tview.NewApplication()
	menu := tview.NewList().
		AddItem("Атаковать дракона", "", '1', func() {
			app.Stop()
			if rand.Intn(100) < 75 {
				damage := hero.Strength + hero.Weapon - dragon.Defence
				if damage < 0 {
					damage = 0
				}
				modifyHealth(&dragon, -damage)
				typewriter(fmt.Sprintf("Герой атакует и наносит %d урона!\n", damage))
			} else {
				typewriter("Герой промахнулся!")
			}
			pause(3)
		}).
		AddItem("Защищаться", "", '2', func() {
			app.Stop()
			if !hero.HasShield {
				hero.Defence += hero.Shield
				hero.HasShield = true
				typewriter("Герой поднял щит!")
			}
			pause(3)
		}).
		AddItem("Пропустить ход", "", '3', func() {
			app.Stop()
			typewriter("Герой пропускает ход.")
			pause(3)
		}).
		AddItem("Выпить зелье", "", '4', func() {
			app.Stop()
			if hero.Potions > 0 {
				modifyHealth(&hero, 500)
				hero.Potions--
				typewriter("Герой выпил зелье и восстановил 500 HP!")
			} else {
				typewriter("Зелья закончились! Герой пропускает ход.")
			}
			pause(3)
		})
	if err := app.SetRoot(menu, true).Run(); err != nil {
		panic(err)
	}

}

func dragonTurn() {
	if dragon.HP == 0 {
		return
	}

	typewriter("Дракон наступает!")
	pause(1)
	if rand.Intn(100) < 50 {
		if rand.Intn(100) < 50 {
			typewriter("Дракон плюётся огненным шаром!")
			if hero.HasShield {
				typewriter("Герой отразил огненный шар щитом!")
			} else {
				damage := dragon.Strength * 2
				modifyHealth(&hero, -damage)
				typewriter(fmt.Sprintf("Герой получил %d урона!\n", damage))
			}
		} else {
			damage := dragon.Strength + dragon.Weapon - hero.Defence
			if damage < 0 {
				damage = 0
			}
			modifyHealth(&hero, -damage)
			typewriter(fmt.Sprintf("Дракон атакует и наносит %d урона!\n", damage))
		}
	} else {
		typewriter("Дракон промахивается.")
	}

	if hero.HasShield {
		hero.Defence -= hero.Shield
		hero.HasShield = false
	}
}

// UI
func pause(seconds time.Duration) { time.Sleep(seconds * time.Second) }

func typewriter(text string) {
	for _, char := range text {
		fmt.Printf("%c", char)
		time.Sleep(40 * time.Millisecond)
	}
	fmt.Println()
}

func loader(text string) {
	clearConsole()
	for _, char := range text {
		fmt.Printf("%c", char)
		time.Sleep(200 * time.Millisecond)
	}
	fmt.Println()
	pause(2)
}

func clearConsole() { fmt.Print("\033[H\033[2J") }

// main
func main() {
	clearConsole()

	// Prologue
	typewriter("Местный житель: Ты… кто ты?")
	typewriter("Введи свое имя:")
	fmt.Scan(&hero.Name)
	clearConsole()
	typewriter("Местный житель: Воин?")
	typewriter(fmt.Sprintf("%v: Охотник. На драконов.", hero.Name))
	typewriter("Местный житель: Если ты правда такой, как говоришь… спаси нас.")
	typewriter(fmt.Sprintf("%v: Расскажи всё, что знаешь.", hero.Name))
	typewriter("Местный житель: Он приходит с севера, из пещеры за скалами. Каждую ночь забирает скот, сжигает дома. Вчера унес человека.")
	pause(4)
	clearConsole()
	typewriter(fmt.Sprintf("%v: Кто-нибудь пытался его остановить?", hero.Name))
	typewriter("Местный житель: Были смельчаки. Но их больше нет.")
	typewriter(fmt.Sprintf("%v: Как он охотится?", hero.Name))
	typewriter("Местный житель: Бьёт с неба, огнём. А если кто бежит — хватает когтями.")
	typewriter(fmt.Sprintf("%v: Тогда он не ждёт, что его встретят лицом к лицу.", hero.Name))
	pause(4)
	clearConsole()
	typewriter("Местный житель: Что ты собираешься делать?")
	typewriter(fmt.Sprintf("%v: Закончить это.", hero.Name))
	typewriter("Местный житель: Если ты победишь…")
	typewriter(fmt.Sprintf("%v: Ждите рассвета. Если дракон не вернётся — значит, я справился.", hero.Name))
	typewriter("Местный житель: А если он прилетит снова?")
	typewriter(fmt.Sprintf("%v: Значит, я не вышел из пещеры.", hero.Name))
	pause(4)
	loader("...")
	loader("...")
	// Road
	clearConsole()
	typewriter(fmt.Sprintf("%v ушёл ещё до рассвета. Деревня провожала его молчанием. Никто не верил в победу, но и слова не сказали — страх перед драконом был сильнее.", hero.Name))
	typewriter(fmt.Sprintf("Тропа к скалам вела через дремучий лес. Когда лес кончился, %v увидел его логово — зияющая дыра в горе, из которой валил дым. Земля вокруг была чёрной, выжженной, словно сама скала боялась зверя, что спал в её недрах.", hero.Name))
	pause(4)
	clearConsole()
	typewriter("Он шагнул вперёд.")
	loader("...")
	loader("...")

	// Duel
	rand.Seed(time.Now().UnixNano())

	for hero.HP > 0 && dragon.HP > 0 {
		clearConsole()
		displayInfo(hero)
		displayInfo(dragon)
		heroTurn()
		if dragon.HP == 0 {
			pause(3)
			clearConsole()
			typewriter("Рассвет встретил деревню тишиной. Впервые за долгие месяцы не было рева крыльев в небе, не было огня, пожирающего дома. Люди ждали, боялись выйти из хижин, боялись увидеть пепел вместо надежды.")
			pause(1)
			typewriter("Но затем, на холме у скал, показалась одинокая фигура.")
			pause(1)
			typewriter(hero.Name)
			pause(1)
			typewriter("Измученный, раненый, но живой. Его плащ был опалён, меч покрыт черной кровью. Он не сказал ни слова, только поднял руку — и бросил к их ногам сломанную кость дракона.")
			pause(1)
			typewriter("— Всё кончено, — только и произнёс он.")
			break
		}
		dragonTurn()
		if hero.HP == 0 {
			pause(3)
			clearConsole()
			typewriter("Рассвет пришёл, но с ним — тьма.")
			pause(1)
			typewriter("Сначала деревня ждала. Они смотрели на север, надеясь увидеть его силуэт на холме, меч, поднятый в победе.")
			pause(1)
			typewriter("Но с горы не пришёл никто.")
			pause(1)
			typewriter("А затем… тень закрыла небо.")
			pause(1)
			typewriter("Дракон.")
			pause(1)
			typewriter("Он вернулся раньше, чем обычно.")
			break
		}
		typewriter("====================================")
		pause(4)
	}
}
