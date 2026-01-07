export const useAuth = () => {
  const token = useState<string | null>('token', () => null)
  const user = useState<any | null>('user', () => null)

  const config = useRuntimeConfig()
  const api = config.public.apiBase

  const isLoggedIn = computed(() => !!token.value)

  const login = async (email: string, password: string) => {
    const res = await $fetch<{ token: string }>(`${api}/auth/login`, {
      method: 'POST',
      body: { email, password }
    })

    token.value = res.token
    await fetchProfile()
  }

  const register = async (email: string, password: string) => {
    await $fetch(`${api}/auth/register`, {
      method: 'POST',
      body: { email, password }
    })
  }

  const fetchProfile = async () => {
    if (!token.value) return

    user.value = await $fetch(`${api}/auth/me`, {
      headers: {
        Authorization: `Bearer ${token.value}`
      }
    })
  }

  const logout = () => {
    token.value = null
    user.value = null
    navigateTo('/login')
  }

  return {
    token,
    user,
    isLoggedIn,
    login,
    register,
    fetchProfile,
    logout
  }
}
