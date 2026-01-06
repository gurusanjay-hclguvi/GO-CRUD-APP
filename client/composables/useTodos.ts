import type { Todo } from '../types/todo'

export const useTodos = () => {
  const config = useRuntimeConfig()
  const api = config.public.apiBase

  const getTodos = (): Promise<Todo[]> =>
    $fetch<Todo[]>(`${api}/todos`)

  const createTodo = (data: Pick<Todo, 'title'>) =>
    $fetch(`${api}/todos`, {
      method: 'POST',
      body: data
    })

  const updateTodo = (id: number, data: Partial<Todo>) =>
    $fetch(`${api}/todos/${id}`, {
      method: 'PUT',
      body: data
    })

  const deleteTodo = (id: number) =>
    $fetch(`${api}/todos/${id}`, {
      method: 'DELETE'
    })

  return {
    getTodos,
    createTodo,
    updateTodo,
    deleteTodo
  }
}
