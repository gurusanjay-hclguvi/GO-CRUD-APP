import type { Todo } from '~/types/todo'

export const useTodos = () => {
  const config = useRuntimeConfig()
  const api = config.public.apiBase

  const { token } = useAuth()

  const authHeaders = () => ({
    Authorization: `Bearer ${token.value}`
  })

  const getTodos = (): Promise<Todo[]> =>
    $fetch(`${api}/todos`, {
      headers: authHeaders()
    })

  const createTodo = (title: string) =>
    $fetch(`${api}/todos`, {
      method: 'POST',
      headers: authHeaders(),
      body: {
        Title: title,
        Completed: false
      }
    })

  const updateTodo = (id: number, data: Partial<Todo>) =>
    $fetch(`${api}/todos/${id}`, {
      method: 'PUT',
      headers: authHeaders(),
      body: data
    })

  const deleteTodo = (id: number) =>
    $fetch(`${api}/todos/${id}`, {
      method: 'DELETE',
      headers: authHeaders()
    })

  return {
    getTodos,
    createTodo,
    updateTodo,
    deleteTodo
  }
}
