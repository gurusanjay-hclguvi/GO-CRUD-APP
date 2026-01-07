export const useAuth = () => {
  const config = useRuntimeConfig()
  const router = useRouter()
  
  const token = useCookie('auth_token', {
    maxAge: 60 * 60 * 24 // 1 day
  })
  const user = useState('auth_user', () => null)

  const fetchUser = async () => {
    if (!token.value) return
    try {
      const { data, error } = await useFetch(`${config.public.apiBase}/auth/me/`, {
        headers: {
          Authorization: `Bearer ${token.value}`
        }
      })
      if (data.value) {
        user.value = data.value
      } else if (error.value) {
        // Token might be invalid
        token.value = null
        user.value = null
      }
    } catch (e) {
      console.error("Failed to fetch user", e)
      token.value = null
      user.value = null
    }
  }

  const login = async (email: string, password: string) => {
    const { data, error } = await useFetch(`${config.public.apiBase}/auth/login`, {
      method: 'POST',
      body: { email, password }
    })

    if (error.value) {
      throw createError({
        statusCode: error.value.statusCode,
        statusMessage: 'Login failed: Invalid credentials'
      })
    }

    if (data.value && data.value.token) {
      token.value = data.value.token
      await fetchUser()
      router.push('/')
    }
  }

  const register = async (name: string, email: string, password: string) => {
    const { data, error } = await useFetch(`${config.public.apiBase}/auth/register`, {
      method: 'POST',
      body: { name, email, password }
    })

    if (error.value) {
      throw createError({
        statusCode: error.value.statusCode,
        statusMessage: 'Registration failed'
      })
    }
    
    // Auto login or redirect to login? Let's redirect to login for simplicity as per plan
    router.push('/login')
  }

  const logout = () => {
    token.value = null
    user.value = null
    router.push('/login')
  }

  return {
    token,
    user,
    login,
    register,
    logout,
    fetchUser
  }
}
