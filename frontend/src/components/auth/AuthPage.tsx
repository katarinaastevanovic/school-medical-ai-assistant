import { useState } from 'react'
import LoginForm from './LoginForm'
import RegisterForm from './RegisterForm'

export default function AuthPage() {
  const [isLogin, setIsLogin] = useState(true)

  return (
    <div className="auth-page">
      <div className="auth-card">

        <div className="auth-switch">
          <button
            className={isLogin ? 'active' : ''}
            onClick={() => setIsLogin(true)}
          >
            Log In
          </button>

          <button
            className={!isLogin ? 'active' : ''}
            onClick={() => setIsLogin(false)}
          >
            Sign Up
          </button>
        </div>

        {isLogin ? <LoginForm /> : <RegisterForm />}

      </div>
    </div>
  )
}