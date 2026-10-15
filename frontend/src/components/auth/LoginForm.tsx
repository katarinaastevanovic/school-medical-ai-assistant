import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { login } from '../../services/authService'

type LoginResponse = {
  token: string
  user: {
    id: number
    email: string
    role: string
  }
}

export default function LoginForm() {
  const navigate = useNavigate()

  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [error, setError] = useState('')

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()

    setError('')

    try {
      const data: LoginResponse = await login(email, password)

      localStorage.setItem('token', data.token)
      localStorage.setItem('user', JSON.stringify(data.user))

      navigate('/home')
    } catch (error) {
      if (error instanceof Error) {
        setError(error.message)
      } else {
        setError('Login failed')
      }
    }
  }

  return (
    <>
      <h1>Welcome Back</h1>

      <form onSubmit={handleSubmit}>

        <input
          type="email"
          placeholder="Enter Your Email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
        />

        <input
          type="password"
          placeholder="Password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />

        {error && <p className="error-message">{error}</p>}

        <button className="submit-btn" type="submit">
          Log In
        </button>

      </form>
    </>
  )
}
