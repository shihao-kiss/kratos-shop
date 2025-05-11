import React, { useEffect, useState } from 'react'
import { getTodos, addTodo, updateTodo, deleteTodo } from './api'

export default function App() {
  const [todos, setTodos] = useState([])
  const [input, setInput] = useState('')

  const load = async () => {
    const result = await getTodos()
    setTodos(Array.isArray(result) ? result : [])
  }

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
    <div style={{ maxWidth: 420, margin: '40px auto', fontFamily: 'Segoe UI, sans-serif', background: '#fff', borderRadius: 12, boxShadow: '0 4px 24px #0001', padding: 32 }}>
      <h2 style={{ textAlign: 'center', color: '#3b82f6', letterSpacing: 2, marginBottom: 24 }}>📝 Todo List</h2>
      <div style={{ display: 'flex', gap: 12, marginBottom: 24 }}>
        <input
          value={input}
          onChange={e => setInput(e.target.value)}
          placeholder="新任务"
          style={{
            flex: 1,
            padding: '10px 14px',
            border: '1px solid #e5e7eb',
            borderRadius: 6,
            fontSize: 16,
            outline: 'none',
            transition: 'border 0.2s',
            boxShadow: '0 1px 2px #0001',
          }}
        />
        <button
          onClick={handleAdd}
          style={{
            background: '#3b82f6',
            color: '#fff',
            border: 'none',
            borderRadius: 6,
            padding: '10px 20px',
            fontSize: 16,
            cursor: 'pointer',
            boxShadow: '0 1px 2px #0001',
            transition: 'background 0.2s',
          }}
        >添加</button>
      </div>
      <ul style={{ listStyle: 'none', padding: 0, margin: 0 }}>
        {todos.map(todo =>
          <li
            key={todo.id}
            style={{
              display: 'flex',
              alignItems: 'center',
              background: todo.completed ? '#f0f9ff' : '#f9fafb',
              borderRadius: 8,
              marginBottom: 12,
              padding: '10px 14px',
              boxShadow: '0 1px 4px #0001',
              transition: 'background 0.2s, box-shadow 0.2s',
              borderLeft: todo.completed ? '4px solid #3b82f6' : '4px solid transparent',
              opacity: todo.completed ? 0.7 : 1,
            }}
            onMouseOver={e => e.currentTarget.style.boxShadow = '0 4px 12px #3b82f622'}
            onMouseOut={e => e.currentTarget.style.boxShadow = '0 1px 4px #0001'}
          >
            <input
              type="checkbox"
              checked={todo.completed}
              onChange={() => handleToggle(todo)}
              style={{ width: 18, height: 18 }}
            />
            <span style={{
              textDecoration: todo.completed ? 'line-through' : 'none',
              marginLeft: 12,
              fontSize: 16,
              color: todo.completed ? '#94a3b8' : '#222',
              flex: 1,
              fontWeight: 500,
            }}>{todo.title}</span>
            <button
              style={{
                marginLeft: 8,
                background: '#f87171',
                color: '#fff',
                border: 'none',
                borderRadius: 6,
                padding: '6px 14px',
                fontSize: 14,
                cursor: 'pointer',
                transition: 'background 0.2s',
              }}
              onClick={() => handleDelete(todo.id)}
            >删除</button>
          </li>
        )}
      </ul>
    </div>
  )
}