package pattern

/*
Команда - Команда — это поведенческий паттерн проектирования, который превращает запросы в объекты,
позволяя передавать их как аргументы при вызове методов, ставить запросы в очередь, логировать их,
а также поддерживать отмену операций.

Плюсы:
* Убирает прямую зависимость между объектами, вызывающими операции, и объектами,
которые их непосредственно выполняют.
* Позволяет реализовать простую отмену и повтор операций.
* Позволяет реализовать отложенный запуск операций.
* Позволяет собирать сложные команды из простых.
* Реализует принцип открытости/закрытости.
Минусы:
* Усложняет код программы из-за введения множества дополнительных классов.
*/

import (
	"encoding/json"
	"log"
	"os"
)

// JSONSaver команда, выполняемая классами при необходимости сохранения моделей в файл
type JSONSaver struct {
	val     interface{}
	encoder json.Encoder
}

// New Создаем экземпляр команды
func (j *JSONSaver) New(filepath string) error {
	file, err := os.Create(filepath)
	if err != nil {
		file = os.Stdout
	}
	(*j).encoder.SetEscapeHTML(true)
	(*j).encoder = *json.NewEncoder(file)
	return err
}

// SetVal устанавливает кодируемое и сохраняемое значение
func (j *JSONSaver) SetVal(val interface{}) {
	(*j).val = val
}

// Save сохраняет значение
func (j *JSONSaver) Save() error {
	err := j.encoder.Encode(j.val)
	return err
}

// Сохраняемые модели Author и Pencil

type Author struct {
	Name  string `json:"name"`
	Bio   string `json:"bio"`
	Books []Book `json:"books"`
	saver JSONSaver
}

func (a Author) Save() {
	a.saver.SetVal(a)
	err := a.saver.Save()
	if err != nil {
		log.Fatal(err)
	}
}

type Book struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Cost        int8   `json:"cost"`
}

type Pencil struct {
	Type  string `json:"type"`
	Mark  string `json:"mark"`
	Cost  int8   `json:"cost"`
	saver JSONSaver
}

func (p Pencil) Save() {
	p.saver.SetVal(p)
	err := p.saver.Save()
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	saver := JSONSaver{}
	_ = saver.New("foo")
	book := Book{Cost: 100, Description: "foo bar", Title: "bar foo"}
	author := Author{Name: "foo", Bio: "bar", Books: []Book{book}, saver: saver}
	author.Save()
	pencil := Pencil{Cost: 100, Mark: "foo", Type: "bar"}
	pencil.Save()
}
