import { BrowserRouter, Routes, Route, Navigate, useLocation } from 'react-router-dom'
import { type PropsWithChildren } from 'react'


import Header from './components/global/Header'
import Footer from './components/global/Footer'
import AuthPage from './components/auth/AuthPage'
import Home from './components/pages/Home'
import Profile from './components/pages/Profile'
import './App.css'
import NotFound from './components/pages/NotFound'
import ClassDetails from './components/pages/ClassDetails'
import Classes from './components/pages/Classes'
import Subjects from './components/pages/Subjects'
import TeacherSubjects from './components/pages/TeacherSubjects'
import TeacherSubjectClassDetails from './components/pages/TeacherSubjectClassDetails'
import StudentGrades from './components/pages/StudentGrades'
import StudentSubjects from './components/pages/StudentSubjects'
import StudentSubjectDetails from './components/pages/StudentSubjectDetails'
import ExaminationRequests from './components/pages/ExaminationRequests'
import Patients from './components/pages/Patients'
import Medical from './components/pages/Medical'
import PatientDetails from './components/pages/PatientDetails'
import Chatbot from './components/pages/Chatbot'



function App() {
  return (
    <BrowserRouter>
      <AppContent />
    </BrowserRouter>
  )
}
function RootRedirect() {
  const token = localStorage.getItem("token")

  return token
    ? <Navigate to="/home" replace />
    : <Navigate to="/login" replace />
}

function ProtectedRoute({ children }: PropsWithChildren) {
  const token = localStorage.getItem('token')

  return token
    ? children
    : <Navigate to="/login" replace />
}

function AppContent() {
  const location = useLocation()

  const hideLayout = location.pathname === '/login'

  return (
    <div className="app-layout">
      {!hideLayout && <Header />}

      <main className="page-content">
        <Routes>
          <Route path="/login" element={<AuthPage />} />
          <Route path="/home" element={<ProtectedRoute><Home /></ProtectedRoute>} />
          <Route path="/profile" element={<ProtectedRoute><Profile /></ProtectedRoute>} />
          <Route path="/classes/:id" element={<ClassDetails />} />
          <Route path="/classes" element={<Classes />} />
          <Route path="/subjects" element={<ProtectedRoute><Subjects /></ProtectedRoute>} />
          <Route path="/teacher-subjects" element={<ProtectedRoute><TeacherSubjects /></ProtectedRoute>} />
          <Route path="/teacher-subjects/:classID/:subjectID" element={<ProtectedRoute><TeacherSubjectClassDetails /></ProtectedRoute>} />
          <Route path="/teacher-subjects/:classID/:subjectID/students/:userID" element={ <ProtectedRoute> <StudentGrades /> </ProtectedRoute>}/>
          <Route path="/student-subjects" element={ <ProtectedRoute> <StudentSubjects /> </ProtectedRoute>}/>
          <Route path="/student-subjects/:subjectID" element={<ProtectedRoute><StudentSubjectDetails /></ProtectedRoute>} />
          <Route path="/examination-requests" element={<ProtectedRoute><ExaminationRequests /></ProtectedRoute>} />
          <Route path="/patients" element={<ProtectedRoute><Patients /></ProtectedRoute>} />
          <Route path="/medical" element={<ProtectedRoute><Medical /></ProtectedRoute>} />
          <Route path="/patients/:userID" element={<ProtectedRoute><PatientDetails /></ProtectedRoute>} />
          <Route path="/chatbot" element={<ProtectedRoute><Chatbot /></ProtectedRoute>} />
          <Route path="/" element={<RootRedirect />} />
          <Route path="*" element={<NotFound />} />
        </Routes>
      </main>

      {!hideLayout && <Footer />}
    </div>
  )
}
export default App