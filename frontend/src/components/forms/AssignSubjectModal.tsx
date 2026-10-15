import { useEffect, useState } from 'react'
import { getAllSubjects, getTeachersBySubject, assignTeacher, type Subject, type TeacherEnriched } from '../../services/schoolService'

interface Props {
  classID: number
  onClose: () => void
  onSuccess: () => void
}

export default function AssignSubjectModal({ classID, onClose, onSuccess }: Props) {
  const [subjects, setSubjects] = useState<Subject[]>([])
  const [teachers, setTeachers] = useState<TeacherEnriched[]>([])
  const [selectedSubject, setSelectedSubject] = useState<number | null>(null)
  const [selectedTeacher, setSelectedTeacher] = useState<number | null>(null)
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState('')

  useEffect(() => {
    getAllSubjects().then(setSubjects).catch(console.error)
  }, [])

  useEffect(() => {
    if (!selectedSubject) {
      setTeachers([])
      setSelectedTeacher(null)
      return
    }
    getTeachersBySubject(selectedSubject)
      .then(setTeachers)
      .catch(console.error)
    setSelectedTeacher(null)
  }, [selectedSubject])

  async function handleAssign() {
    if (!selectedSubject || !selectedTeacher) {
      setError('Please select both a subject and a teacher.')
      return
    }
    setLoading(true)
    setError('')
    try {
      await assignTeacher(selectedTeacher, classID, selectedSubject)
      onSuccess()
      onClose()
    } catch (err: any) {
      setError(err.message)
    } finally {
      setLoading(false)
    }
  }

  return (
    <div className="modal-overlay" onClick={onClose}>
      <div className="modal" onClick={e => e.stopPropagation()}>
        <h2>Assign Subject</h2>

        <div className="form-group">
          <label>Subject</label>
          <select
            className="role-select"
            value={selectedSubject ?? ''}
            onChange={e => setSelectedSubject(Number(e.target.value))}
          >
            <option value="">Select subject...</option>
            {subjects.map(s => (
              <option key={s.ID} value={s.ID}>{s.Name}</option>
            ))}
          </select>
        </div>

        {selectedSubject && (
          <div className="form-group">
            <label>Teacher</label>
            {teachers.length === 0 ? (
              <p className="no-students">No teachers assigned to this subject yet.</p>
            ) : (
              <select
                className="role-select"
                value={selectedTeacher ?? ''}
                onChange={e => setSelectedTeacher(Number(e.target.value))}
              >
                <option value="">Select teacher...</option>
                {teachers.map(t => (
                  <option key={t.userID} value={t.userID}>
                    {t.firstName} {t.lastName}
                  </option>
                ))}
              </select>
            )}
          </div>
        )}

        {error && <p className="error">{error}</p>}

        <div className="modal-actions">
          <button className="btn-secondary" onClick={onClose}>Cancel</button>
          <button
            className="btn-primary"
            onClick={handleAssign}
            disabled={loading || !selectedSubject || !selectedTeacher}
          >
            {loading ? 'Assigning...' : 'Assign'}
          </button>
        </div>
      </div>
    </div>
  )
}