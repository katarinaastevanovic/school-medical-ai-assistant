import { useEffect, useState } from 'react'
import { getFreeStudents, addStudentToClass, type FreeStudent } from '../../services/schoolService'

interface Props {
  classID: number
  onClose: () => void
  onSuccess: () => void
}

export default function AddStudentsModal({ classID, onClose, onSuccess }: Props) {
  const [freeStudents, setFreeStudents] = useState<FreeStudent[]>([])
  const [selected, setSelected] = useState<number[]>([])
  const [loading, setLoading] = useState(false)

  useEffect(() => {
    getFreeStudents()
      .then(setFreeStudents)
      .catch(err => console.error(err))
  }, [])

  function toggleStudent(id: number) {
    setSelected(prev =>
      prev.includes(id) ? prev.filter(s => s !== id) : [...prev, id]
    )
  }

 async function handleSubmit() {
  if (selected.length === 0) return
  setLoading(true)

  try {
    await Promise.all(
      selected.map(userID => addStudentToClass(classID, userID))
    )
    onSuccess()
    onClose()
  } catch (err) {
    console.error(err)
  } finally {
    setLoading(false)
  }
}

  return (
    <div className="modal-overlay" onClick={onClose}>
      <div className="modal" onClick={e => e.stopPropagation()}>
        <h2>Add Students</h2>

        {freeStudents.length === 0 ? (
          <p className="no-students">No available students.</p>
        ) : (
          <div className="checkbox-list">
            {freeStudents.map(s => (
              <label key={s.id} className="checkbox-item">
                <input
                  type="checkbox"
                  checked={selected.includes(s.id)}
                  onChange={() => toggleStudent(s.id)}
                />
                {s.firstName} {s.lastName}
              </label>
            ))}
          </div>
        )}

        <div className="modal-actions">
          <button className="btn-secondary" onClick={onClose} disabled={loading}>
            Cancel
          </button>
          <button className="btn-primary" onClick={handleSubmit} disabled={loading || selected.length === 0}>
            {loading ? 'Adding...' : `Add (${selected.length})`}
          </button>
        </div>
      </div>
    </div>
  )
}