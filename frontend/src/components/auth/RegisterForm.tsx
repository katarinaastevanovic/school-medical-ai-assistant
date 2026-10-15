import { useState } from 'react'
import { register } from '../../services/authService'

type FormErrors = {
  umcn?: string
  firstName?: string
  lastName?: string
  email?: string
  password?: string
  confirmPassword?: string
  general?: string
}

export default function RegisterForm() {

  const [umcn, setUmcn] = useState('')
  const [firstName, setFirstName] = useState('')
  const [lastName, setLastName] = useState('')
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [confirmPassword, setConfirmPassword] = useState('')
  const [role, setRole] = useState('STUDENT')

  const [errors, setErrors] = useState<FormErrors>({})
  const [success, setSuccess] = useState('')

  const validateForm = () => {

    const newErrors: FormErrors = {}

    // JMBG - UMCN
    if (!umcn.trim()) {
      newErrors.umcn = 'JMBG is required'
    } else if (!/^\d{13}$/.test(umcn)) {
      newErrors.umcn = 'JMBG must contain 13 digits'
    } else {

      const day = parseInt(umcn.substring(0, 2))
      const month = parseInt(umcn.substring(2, 4))
      const yearPart = parseInt(umcn.substring(4, 7))

      let year = 0

      if (yearPart >= 800) {
        year = 1000 + yearPart
      } else {
        year = 2000 + yearPart
      }

      const birthDate = new Date(year, month - 1, day)

      const isValidDate =
        birthDate.getFullYear() === year &&
        birthDate.getMonth() === month - 1 &&
        birthDate.getDate() === day

      if (!isValidDate) {

        newErrors.umcn = 'Invalid birth date in JMBG'

      } else if (birthDate > new Date()) {

        newErrors.umcn = 'Birth date cannot be in the future'

      } else if (year < 1900) {

        newErrors.umcn = 'Invalid year in JMBG'

      }
    }
    //First name
    if (!firstName.trim()) {
      newErrors.firstName = 'First name is required'
    }
    //Last name
    if (!lastName.trim()) {
      newErrors.lastName = 'Last name is required'
    }

    // Email
    if (!email.trim()) {
      newErrors.email = 'Email is required'
    } else if (!/\S+@\S+\.\S+/.test(email)) {
      newErrors.email = 'Invalid email format'
    }

    // Password
    if (!password) {
      newErrors.password = 'Password is required'
    } else if (password.length < 6) {
      newErrors.password = 'Password must contain at least 6 characters'
    }

    // Confirm password
    if (!confirmPassword) {
      newErrors.confirmPassword = 'Please confirm password'
    } else if (password !== confirmPassword) {
      newErrors.confirmPassword = 'Passwords do not match'
    }

    setErrors(newErrors)

    return Object.keys(newErrors).length === 0
  }

  const handleSubmit = async (e: React.FormEvent) => {

    e.preventDefault()

    setSuccess('')
    setErrors({})

    const isValid = validateForm()

    if (!isValid) {
      return
    }

    try {

      await register(
        umcn,
        firstName,
        lastName,
        email,
        password,
        role,
      )

      setSuccess('Account created successfully')

      setUmcn('')
      setFirstName('')
      setLastName('')
      setEmail('')
      setPassword('')
      setConfirmPassword('')
      setRole('STUDENT')

    } catch (error) {

      if (error instanceof Error) {
        setErrors({
          general: error.message,
        })
      } else {
        setErrors({
          general: 'Registration failed',
        })
      }

    }
  }

  return (
  <>
    <h1>Create Account</h1>

    <form onSubmit={handleSubmit}>

      <div className="form-grid">

        <div>
          <input
            type="text"
            placeholder="UMCN"
            value={umcn}
            onChange={(e) => setUmcn(e.target.value)}
          />

          {errors.umcn && (
            <p className="field-error">
              {errors.umcn}
            </p>
          )}
        </div>

        <div>
          <select
            value={role}
            onChange={(e) => setRole(e.target.value)}
            className="role-select"
          >
            <option value="STUDENT">Student</option>
            <option value="TEACHER">Teacher</option>
            <option value="DOCTOR">Doctor</option>
            <option value="ADMIN">Admin</option>
          </select>
        </div>

        <div>
          <input
            type="text"
            placeholder="First Name"
            value={firstName}
            onChange={(e) => setFirstName(e.target.value)}
          />

          {errors.firstName && (
            <p className="field-error">
              {errors.firstName}
            </p>
          )}
        </div>

        <div>
          <input
            type="text"
            placeholder="Last Name"
            value={lastName}
            onChange={(e) => setLastName(e.target.value)}
          />

          {errors.lastName && (
            <p className="field-error">
              {errors.lastName}
            </p>
          )}
        </div>

        

        <div className="full-width">
          <input
            type="email"
            placeholder="Enter Your Email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
          />

          {errors.email && (
            <p className="field-error">
              {errors.email}
            </p>
          )}
        </div>

        <div>
          <input
            type="password"
            placeholder="Password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />

          {errors.password && (
            <p className="field-error">
              {errors.password}
            </p>
          )}
        </div>

        <div>
          <input
            type="password"
            placeholder="Confirm Password"
            value={confirmPassword}
            onChange={(e) => setConfirmPassword(e.target.value)}
          />

          {errors.confirmPassword && (
            <p className="field-error">
              {errors.confirmPassword}
            </p>
          )}
        </div>

      </div>

      {errors.general && (
        <p className="error-message">
          {errors.general}
        </p>
      )}

      {success && (
        <p className="success-message">
          {success}
        </p>
      )}

      <button className="submit-btn" type="submit">
        Create Account
      </button>

    </form>
  </>
)
}