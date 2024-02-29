package pattern

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/
/*
Паттерн State относится к поведенческим паттернам уровня объекта.

Паттерн State позволяет объекту изменять свое поведение в зависимости от внутреннего состояния
и является объектно-ориентированной реализацией конечного автомата.
Поведение объекта изменяется настолько, что создается впечатление, будто изменился класс объекта.

Паттерн должен применяться:

- когда поведение объекта зависит от его состояния
- поведение объекта должно изменяться во время выполнения программы
- состояний достаточно много и использовать для этого условные операторы, разбросанные по коду, достаточно затруднительно

Требуется для реализации:

Класс Context, представляет собой объектно-ориентированное представление конечного автомата;
Абстрактный класс State, определяющий интерфейс различных состояний;
Класс ConcreteStateA реализует одно из поведений, ассоциированное с определенным состоянием;
Класс ConcreteStateB реализует одно из поведений, ассоциированное с определенным состоянием.

[!] В описании паттерна применяются общие понятия, такие как Класс, Объект, Абстрактный класс.
Применимо к языку Go, это Пользовательский Тип, Значение этого Типа и Интерфейс.
Также в языке Go вместо наследования используется композиция.
*/

// MobileAlertStater provides a common interface for various states.
type MobileAlertStater interface {
	Alert() string
}

// MobileAlert implements an alert depending on its state.
type MobileAlert struct {
	state MobileAlertStater
}

// Alert returns a alert string
func (a *MobileAlert) Alert() string {
	return a.state.Alert()
}

// SetState changes state
func (a *MobileAlert) SetState(state MobileAlertStater) {
	a.state = state
}

// NewMobileAlert is the MobileAlert constructor.
func NewMobileAlert() *MobileAlert {
	return &MobileAlert{state: &MobileAlertVibration{}}
}

// MobileAlertVibration implements vibration alert
type MobileAlertVibration struct {
}

// Alert returns a alert string
func (a *MobileAlertVibration) Alert() string {
	return "Vrrr... Brrr... Vrrr..."
}

// MobileAlertSong implements beep alert
type MobileAlertSong struct {
}

// Alert returns a alert string
func (a *MobileAlertSong) Alert() string {
	return "Белые розы, Белые розы. Беззащитны шипы..."
}

func RunState() {
	// expect "Vrrr... Brrr... Vrrr..." + "Белые розы, Белые розы. Беззащитны шипы..."

	mobile := NewMobileAlert()

	fmt.Println(mobile.Alert())
	mobile.SetState(&MobileAlertSong{})
	fmt.Println(mobile.Alert())
}
