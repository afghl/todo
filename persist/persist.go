package persist

import (
	"io/ioutil"
	"os"
	"todo/todos"
)

const (
  fileName = "todo.json"
)

type Storable interface {
	Bytes() []byte
}

func Load() (*todos.Todos, error) {
  if !fileExists() {
		return nil, nil
	}
	serializedItems, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	if len(serializedItems) == 0 {
		return nil, nil
	}

	todoItems := todos.FromBytes(serializedItems)

	return todoItems, nil
}

func Save(storable Storable) error {
	if err := ioutil.WriteFile(fileName, storable.Bytes(), 0644); err != nil {
		return err
	}
	return nil
}

func fileExists() bool {
	_, err := os.Stat(fileName)
	return err == nil
}
