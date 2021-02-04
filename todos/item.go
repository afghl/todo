package todos

type Todo struct {
  Id int
  Content string
  isDone  bool
}

func (t *Todo) Done() {
  t.isDone = true
}

func (t Todo) IsDone() bool {
  return t.isDone
}

func FromBytes([]byte) *Todos {
  t := &Todos{
  }
  return t
}

type Todos struct {
  t []Todo
}

func (t Todos) Count() int {
  return len(t.t)
}

func (t Todos) Append(todo Todo) {
  todo.Id = len(t.t) + 1
  t.t = append(t.t, todo)
}

func (t Todos) Remove(todo Todo) {
  index := -1
  for i, t := range t.t {
    if t.Id == todo.Id {
      index = i
    }
  }

  t.t = remove(t.t, index)
}

func (t Todos) Iterate(f func(int, Todo)) {
  for i, todo := range t.t {
    f(i, todo)
  }
}

func remove(s []Todo, index int) []Todo {
  return append(s[:index], s[index+1:]...)
}

func (t Todos) Bytes() []byte {
  return nil
}

func (t Todos) Find(id int) *Todo {
  for _, todo := range t.t {
    if todo.Id == id {
      return &todo
    }
  }
  return nil
}

func NewTodo(content string) *Todo {
  return &Todo{Content: content}
}
