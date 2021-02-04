const parseArgs = require('minimist')
import { Todo } from './model/Todo'
import { save, getAll, get } from './persist'

const main = () => {
  const args = parseArgs(process.argv.slice(2))
  // https://github.com/substack/minimist

  const type = args['_'][0]

  if (type == 'add') {
    add(args['_'][1])
  } else if (type == 'done') {
    done(args['_'][1])
  } else if (type == 'list') {
    list(args['_'][1])
  }
}

const add = (content) => {
  const todo = new Todo(content)
  // persist
  save(todo)
  console.log(`\n${todo.toString()}\n\n\nItem ${todo.id} added.`);
}

const done = (id) => {
  const todo = get(id)
  if (todo == null) {
    console.log('todo not found.');
    return
  }

  todo.done()
  save(todo)
  console.log(`\nItem ${todo.id} done.`);
}

const list = (all) => {

}

main()
