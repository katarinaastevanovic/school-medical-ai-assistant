import { useState } from 'react'
import { createExamination } from '../../services/medicalService'

interface Props {
  requestID: number
  patientName: string
  onClose: () => void
  onSuccess: () => void
}

export default function CreateExaminationModal({ requestID, patientName, onClose, onSuccess }: Props) {
  const [diagnosis, setDiagnosis] = useState('')
  const [therapy, setTherapy] = useState('')
  const [note, setNote] = useState('')
  const [loading, setLoading] = useState(false)

  async function handleSave() {
    setLoading(true)
    try {
      await createExamination(requestID, diagnosis, therapy, note)
      onSuccess()
    } catch (err) {
      console.error(err)
    }
    setLoading(false)
  }

  return (
    <div className="modal-overlay" onClick={onClose}>
      <div className="modal" onClick={e => e.stopPropagation()}>
        <h2>Examination — {patientName}</h2>

        <div className="form-group">
          <label>Diagnosis</label>
          <input type="text" value={diagnosis} onChange={e => setDiagnosis(e.target.value)} />
        </div>

        <div className="form-group">
          <label>Therapy</label>
          <input type="text" value={therapy} onChange={e => setTherapy(e.target.value)} />
        </div>

        <div className="form-group">
          <label>Note</label>
          <input type="text" value={note} onChange={e => setNote(e.target.value)} />
        </div>

        <div className="modal-actions">
          <button className="btn-secondary" onClick={onClose}>Cancel</button>
          <button className="btn-primary" onClick={handleSave} disabled={loading || !diagnosis}>
            {loading ? 'Saving...' : 'Save'}
          </button>
        </div>
      </div>
    </div>
  )
}