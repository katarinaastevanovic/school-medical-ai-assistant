import { useEffect, useState } from 'react'
import { useParams, useNavigate } from 'react-router-dom'
import { getStudentsFromClass } from '../../services/schoolService'
import { getMemberById } from '../../services/authService'

interface StudentWithInfo {
  ID: number
  UserID: number
  firstName: string
  lastName: string
}

export default function MySubjectClassDetails() {
  const { classID, subjectID } = useParams()
  const navigate = useNavigate()
  const [students, setStudents] = useState<StudentWithInfo[]>([])

  useEffect(() => {
    if (!classID) return

    getStudentsFromClass(Number(classID))
      .then(async (data) => {
        const enriched = await Promise.all(
          data.map(async (s) => {
            const member = await getMemberById(s.UserID)
            return {
              ID: s.ID,
              UserID: s.UserID,
              firstName: member.firstName,
              lastName: member.lastName,
            }
          })
        )
        setStudents(enriched)
      })
      .catch(err => console.error(err))
  }, [classID])

  return (
    <div className="crud-page">
      <h1>Students</h1>

      <div className="class-grid">
        {students.map((s) => (
          <div
            key={s.ID}
            className="student-card"
            style={{ cursor: 'pointer' }}
            onClick={() => navigate(`/teacher-subjects/${classID}/${subjectID}/students/${s.UserID}`)}
          >
            <img
              src="https://bootdey.com/img/Content/avatar/avatar7.png"
              alt="avatar"
              className="student-avatar"
            />
            <span>{s.firstName} {s.lastName}</span>
          </div>
        ))}
      </div>
    </div>
  )
}