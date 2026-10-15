import { useState } from "react"
import { createGrade } from "../../services/schoolService"

interface Props {
  studentUserID: number
  subjectID: number
  onClose: () => void
  onSuccess: () => void
}

export default function AddGradeModal({ studentUserID, subjectID, onClose, onSuccess }: Props) {
  const [value, setValue] = useState(5)
  const [date, setDate] = useState("")
  const [loading, setLoading] = useState(false)

  async function handleSave() {
    setLoading(true)
    try {
      await createGrade(studentUserID, subjectID, value, new Date(date).toISOString())
      onSuccess()
      onClose()
    } catch (err) {
      console.error(err)
    }
    setLoading(false)
  }

  return (
    <div className="modal-overlay" onClick={onClose}>
      <div className="modal" onClick={e => e.stopPropagation()}>
        <h2>Add Grade</h2>

        <div className="form-group">
          <label>Grade</label>
          <select
            className="role-select"
            value={value}
            onChange={e => setValue(Number(e.target.value))}
          >
            {[1, 2, 3, 4, 5].map(v => (
              <option key={v}>{v}</option>
            ))}
          </select>
        </div>

        <div className="form-group">
          <label>Date</label>
          <input
            type="date"
            value={date}
            onChange={e => setDate(e.target.value)}
          />
        </div>

        <div className="modal-actions">
          <button className="btn-secondary" onClick={onClose}>Cancel</button>
          <button className="btn-primary" disabled={loading || !date} onClick={handleSave}>
            Save
          </button>
        </div>
      </div>
    </div>
  )
}