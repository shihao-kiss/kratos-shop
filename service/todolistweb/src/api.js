import axios from 'axios'

export const getTodos = () =>
  axios.get('/v1/todos').then(res => res.data.data.todos)

export const addTodo = (title) =>
  axios.post('/v1/todo', { title }).then(res => res.data)

export const updateTodo = (id, data) =>
  axios.put(`/v1/todo/${id}`, data).then(res => res.data)

export const deleteTodo = (id) =>
  axios.delete(`/v1/todo/${id}`).then(res => res.data)