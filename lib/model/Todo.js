class Todo {
  constructor(content) {
    // todo get id
    this.id = 1
    this.content = content
    // todo, enum
    this.state = "Pending"
  }

  isDone() {
    return this.state == "Done"
  }

  done() {
    this.state = "Done"
  }

  toString() {
    return `${this.id}. ${this.isDone() ? '[Done]' : '[Pending]'} ${this.content}.`
  }
}

export { Todo }
