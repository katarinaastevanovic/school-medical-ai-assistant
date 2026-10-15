import { useEffect, useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { getMySubjects, type TeacherSubjectClass } from '../../services/schoolService'

export default function MySubjects() {
  const [items, setItems] = useState<TeacherSubjectClass[]>([])
  const navigate = useNavigate()

  useEffect(() => {
    getMySubjects().then(setItems).catch(console.error)
  }, [])

  return (
    <div className="crud-page">
      <h1>My Subjects</h1>

      <div className="class-grid">
        {items.map(item => (
          <div
            key={item.id}
            className="subject-card"
            style={{ cursor: 'pointer' }}
            onClick={() => navigate(`/teacher-subjects/${item.classID}/${item.subjectID}`)}
          >
            <span className="subject-card-name">{item.subjectName}</span>
            <span className="subject-card-teacher">{item.className}</span>
          </div>
        ))}
      </div>
    </div>
  )
}