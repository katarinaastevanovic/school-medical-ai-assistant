import { useNavigate } from 'react-router-dom'
import { FaHome, FaUserCircle, FaComments } from 'react-icons/fa'
import { logout } from '../../services/authService'


export default function Header() {
  const navigate = useNavigate()

  const isLoggedIn = !!localStorage.getItem('token')
  const user = JSON.parse(localStorage.getItem('user') || '{}')
  const isAdmin = user.role === 'ADMIN'
  const isTeacher = user.role === 'TEACHER'
  const isStudent = user.role === 'STUDENT'
  const isDoctor = user.role === 'DOCTOR'


  const handleLogout = () => {
    logout()
    navigate('/login')
  }

  return (
    <header className="main-header">
      <button className="home-btn" onClick={() => navigate('/')}>
        <FaHome size={28} />
      </button>

      {isAdmin && (
        <nav className="header-nav">
          <button className="nav-link" onClick={() => navigate('/classes')}>
            Classes
          </button>
          <button className="nav-link" onClick={() => navigate('/subjects')}>
            Subjects
          </button>
        </nav>
      )}

      {isTeacher && (
        <nav className="header-nav">
          <button className="nav-link" onClick={() => navigate('/teacher-subjects')}>
            My Subjects
          </button>
        </nav>
      )}

      {isDoctor && (
        <nav className="header-nav">
          <button className="nav-link" onClick={() => navigate('/patients')}>
            Patients
          </button>
          <button className="nav-link" onClick={() => navigate('/examination-requests')}>
            Requests
          </button>
        </nav>
      )}

      {isStudent && (
        <nav className="header-nav">
          <button className="nav-link" onClick={() => navigate('/student-subjects')}>
            My Subjects
          </button>
          <button className="nav-link" onClick={() => navigate('/medical')}>
            Medical
          </button>
        </nav>
      )}

      {isLoggedIn && (
  <button className="chatbot-btn" onClick={() => navigate('/chatbot')} title="Asistent">
    <FaComments size={26} />
  </button>
)}

      {isLoggedIn && (
        <div className="header-actions">
          <button className="profile-btn" onClick={() => navigate('/profile')}>
            <FaUserCircle size={32} />
          </button>
          <button className="logout-btn" onClick={handleLogout}>
            Logout
          </button>
        </div>
      )}
    </header>
  )
}