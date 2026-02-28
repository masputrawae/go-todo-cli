package todo

import (
	"encoding/json"
	"errors"
	"github.com/masputrawae/go-todo/model"
	"github.com/masputrawae/go-todo/utils"
)

var Priorities = map[rune]model.Priority{
	'A': {Name: "Highest", Emoji: "🔴"},
	'B': {Name: "High", Emoji: "🟠"},
	'C': {Name: "Medium", Emoji: "🟢"},
	'D': {Name: "Low", Emoji: "🔵"},
	'E': {Name: "Lowest", Emoji: "🟣"},
}

type Todo struct {
	FilePath string
	Data     []model.Todo
}

type Manage interface {
	Load() error
	Save() error
	GenID() int

	Add(nt model.Todo)
	Edit(index int, nt model.Todo)
	Delete(index int)

	FindIndexByID(id int) (int, error)
	FindByIndex(i int) model.Todo
	FindAll() []model.Todo

	IsValidPriority(s string) bool
}

func New(d []model.Todo, fp string) Manage {
	return &Todo{Data: d, FilePath: fp}
}

// save data todos
func (t *Todo) Save() error {
	file, err := utils.ResolveFile(t.FilePath)
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewEncoder(file).Encode(t.Data)
}

// load data todos
func (t *Todo) Load() error {
	file, err := utils.ResolveFile(t.FilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&t.Data)
	if err != nil {
		if err.Error() == "EOF" {
			return json.NewEncoder(file).Encode(t.Data)
		}
		return err
	}
	return nil
}

// add data todos
func (t *Todo) Add(nt model.Todo) {
	t.Data = append(t.Data, nt)

}

// edit data todo by index
func (t *Todo) Edit(index int, nt model.Todo) {
	t.Data[index] = nt
}

// remove todo
func (t *Todo) Delete(index int) {
	t.Data = append(t.Data[:index], t.Data[index+1:]...)
}

// find index by id
func (t *Todo) FindIndexByID(id int) (int, error) {
	for i := range t.Data {
		if t.Data[i].ID == id {
			return i, nil
		}
	}
	return -1, errors.New("id tidak ditemukan")
}

// find todo by index
func (t *Todo) FindByIndex(i int) model.Todo {
	return t.Data[i]
}

// find all
func (t *Todo) FindAll() []model.Todo {
	var results = make([]model.Todo, len(t.Data))
	copy(results, t.Data)
	return results
}

// generate id todo
func (t *Todo) GenID() int {
	id := 1
	for i := range t.Data {
		if t.Data[i].ID == id {
			id++
		}
	}
	return id
}

// check valid priority
func (t *Todo) IsValidPriority(s string) bool {
	if !utils.IsLetter(s) {
		return false
	}
	for k := range Priorities {
		if k == rune(s[0]) {
			return true
		}
	}
	return false
}
