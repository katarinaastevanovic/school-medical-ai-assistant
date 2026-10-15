import { useEffect, useState } from 'react'
import { getTeachersBySubject, getAllTeachers, addTeacherToSubject, type TeacherEnriched } from '../../services/schoolService'

interface Props {
  subjectID: number
  subjectName: string
  onClose: () => void
}

export default function SubjectTeachersModal({ subjectID, subjectName, onClose }: Props) {
  const [assigned, setAssigned] = useState<TeacherEnriched[]>([])
  const [available, setAvailable] = useState<TeacherEnriched[]>([])
  const [selected, setSelected] = useState<number[]>([])
  const [loading, setLoading] = useState(false)

  async function fetchData() {
    try {
      const [all, current] = await Promise.all([
        getAllTeachers(),
        getTeachersBySubject(subjectID),
      ])
      setAssigned(current)
      const assignedIDs = new Set(current.map(t => t.userID))
      setAvailable(all.filter(t => !assignedIDs.has(t.userID)))
    } catch (err) {
      console.error(err)
    }
  }

  useEffect(() => {
    fetchData()
  }, [subjectID])

  function toggleTeacher(userID: number) {
    setSelected(prev =>
      prev.includes(userID) ? prev.filter(id => id !== userID) : [...prev, userID]
    )
  }

  async function handleAdd() {
    if (selected.length === 0) return
    setLoading(true)
    try {
      await Promise.all(selected.map(userID => addTeacherToSubject(subjectID, userID)))
      setSelected([])
      fetchData()
    } catch (err) {
      console.error(err)
    } finally {
      setLoading(false)
    }
  }

  return (
    <div className="modal-overlay" onClick={onClose}>
      <div className="modal" onClick={e => e.stopPropagation()}>
        <h2>Teachers — {subjectName}</h2>

        {assigned.length > 0 && (
          <div>
            <p className="modal-section-label">Assigned</p>
            <div className="assigned-list">
              {assigned.map(t => (
                <span key={t.userID} className="assigned-tag">
                  {t.firstName} {t.lastName}
                </span>
              ))}
            </div>
          </div>
        )}

        {available.length > 0 ? (
          <div>
            <p className="modal-section-label">Add teachers</p>
            <div className="checkbox-list">
              {available.map(t => (
                <label key={t.userID} className="checkbox-item">
                  <input
                    type="checkbox"
                    checked={selected.includes(t.userID)}
                    onChange={() => toggleTeacher(t.userID)}
                  />
                  {t.firstName} {t.lastName}
                </label>
              ))}
            </div>
          </div>
        ) : (
          <p className="no-students">No available teachers to add.</p>
        )}

        <div className="modal-actions">
          <button className="btn-secondary" onClick={onClose}>Close</button>
          {available.length > 0 && (
            <button className="btn-primary" onClick={handleAdd} disabled={loading || selected.length === 0}>
              {loading ? 'Adding...' : `Add (${selected.length})`}
            </button>
          )}
        </div>
      </div>
    </div>
  )
}