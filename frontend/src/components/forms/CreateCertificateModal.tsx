import { useState } from 'react'
import { createCertificate } from '../../services/medicalService'

interface Props {
  patientUserID: number
  patientName: string
  onClose: () => void
  onSuccess: () => void
}

export default function CreateCertificateModal({ patientUserID, patientName, onClose, onSuccess }: Props) {
  const [type, setType] = useState('REGULAR')
  const [note, setNote] = useState('')
  const [date, setDate] = useState('')
  const [loading, setLoading] = useState(false)

  async function handleSave() {
    setLoading(true)
    try {
      await createCertificate(patientUserID, type, note, new Date(date).toISOString())
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
        <h2>Certificate — {patientName}</h2>

        <div className="form-group">
          <label>Type</label>
          <select className="role-select" value={type} onChange={e => setType(e.target.value)}>
            <option value="REGULAR">Regular</option>
            <option value="PE">PE</option>
            <option value="EXCURSION">Excursion</option>
            <option value="SICKNESS">Sickness</option>
          </select>
        </div>

        <div className="form-group">
          <label>Date</label>
          <input type="date" value={date} onChange={e => setDate(e.target.value)} />
        </div>

        <div className="form-group">
          <label>Note</label>
          <input type="text" value={note} onChange={e => setNote(e.target.value)} placeholder="Optional" />
        </div>

        <div className="modal-actions">
          <button className="btn-secondary" onClick={onClose}>Cancel</button>
          <button className="btn-primary" onClick={handleSave} disabled={loading || !date}>
            {loading ? 'Saving...' : 'Issue'}
          </button>
        </div>
      </div>
    </div>
  )
}