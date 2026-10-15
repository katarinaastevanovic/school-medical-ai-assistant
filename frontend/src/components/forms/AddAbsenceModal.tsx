import { useState } from "react"
import { createAbsence } from "../../services/schoolService"

interface Props {
  studentUserID: number
  subjectID: number
  onClose: () => void
  onSuccess: () => void
}

export default function AddAbsenceModal({ studentUserID, subjectID, onClose, onSuccess }: Props) {
  const [date, setDate] = useState("")
  const [loading, setLoading] = useState(false)

  async function handleSave() {
    setLoading(true)
    try {
      await createAbsence(studentUserID, subjectID, new Date(date).toISOString())
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
        <h2>Add Absence</h2>

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