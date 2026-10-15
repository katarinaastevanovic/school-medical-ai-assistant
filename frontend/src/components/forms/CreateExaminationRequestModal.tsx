import { useEffect, useState } from 'react'
import { createExaminationRequest } from '../../services/medicalService'
import { getMemberById } from '../../services/authService'

interface Doctor {
  id: number
  firstName: string
  lastName: string
}

interface Props {
  onClose: () => void
  onSuccess: () => void
}

export default function CreateExaminationRequestModal({ onClose, onSuccess }: Props) {
  const [doctors, setDoctors] = useState<Doctor[]>([])
  const [selectedDoctor, setSelectedDoctor] = useState<number | null>(null)
  const [type, setType] = useState('REGULAR')
  const [loading, setLoading] = useState(false)

  useEffect(() => {
    fetch('http://localhost:8080/members/doctors', {
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${localStorage.getItem('token') ?? ''}`,
      }
    })
      .then(r => r.json())
      .then(setDoctors)
      .catch(console.error)
  }, [])

  async function handleSave() {
    if (!selectedDoctor) return
    setLoading(true)
    try {
      await createExaminationRequest(selectedDoctor, type)
      onSuccess()
    } catch (err) {
      console.error(err)
    }
    setLoading(false)
  }

  return (
    <div className="modal-overlay" onClick={onClose}>
      <div className="modal" onClick={e => e.stopPropagation()}>
        <h2>New Examination Request</h2>

        <div className="form-group">
          <label>Doctor</label>
          <select
            className="role-select"
            value={selectedDoctor ?? ''}
            onChange={e => setSelectedDoctor(Number(e.target.value))}
          >
            <option value="">Select doctor...</option>
            {doctors.map(d => (
              <option key={d.id} value={d.id}>
                {d.firstName} {d.lastName}
              </option>
            ))}
          </select>
        </div>

        <div className="form-group">
          <label>Type</label>
          <select className="role-select" value={type} onChange={e => setType(e.target.value)}>
            <option value="REGULAR">Regular</option>
            <option value="SPECIALIST">Specialist</option>
            <option value="URGENT">Urgent</option>
          </select>
        </div>

        <div className="modal-actions">
          <button className="btn-secondary" onClick={onClose}>Cancel</button>
          <button className="btn-primary" onClick={handleSave} disabled={loading || !selectedDoctor}>
            {loading ? 'Sending...' : 'Send'}
          </button>
        </div>
      </div>
    </div>
  )
}