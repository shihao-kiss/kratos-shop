import React, { useEffect, useState } from 'react'
import { getTodos, addTodo, updateTodo, deleteTodo } from './api'

export default function App() {
  const [todos, setTodos] = useState([])
  const [input, setInput] = useState('')

  const load = async () => setTodos(await getTodos())

  useEffect(() => { load() }, [])

  const handleAdd = async () => {
    if (!input.trim()) return
    await addTodo(input)
    setInput('')
    load()
  }

  const handleToggle = async (todo) => {
    await updateTodo(todo.id, { ...todo, completed: !todo.completed })
    load()
  }

  const handleDelete = async (id) => {
    await deleteTodo(id)
    load()
  }

  return (
    <div style={{ maxWidth: 400, margin: '40px auto', fontFamily: 'sans-serif' }}>
      <h2>Todo List</h2>
      <div>
        <input
          value={input}
          onChange={e => setInput(e.target.value)}
          placeholder="新任务"
        />
        <button onClick={handleAdd}>添加</button>
      </div>
      <ul>
        {todos.map(todo =>
          <li key={todo.id} style={{ margin: '8px 0' }}>
            <input
              type="checkbox"
              checked={todo.completed}
              onChange={() => handleToggle(todo)}
            />
            <span style={{
              textDecoration: todo.completed ? 'line-through' : 'none',
              marginLeft: 8
            }}>{todo.title}</span>
            <button style={{ marginLeft: 8 }} onClick={() => handleDelete(todo.id)}>删除</button>
          </li>
        )}
      </ul>
    </div>
  )
}