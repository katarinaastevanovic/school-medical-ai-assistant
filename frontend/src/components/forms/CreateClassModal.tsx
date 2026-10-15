import { useState } from 'react'
import { createClass } from '../../services/schoolService'

interface Props {
  onClose: () => void
  onSuccess: () => void
}

export default function CreateClassModal({ onClose, onSuccess }: Props) {
  const [name, setName] = useState('')
  const [year, setYear] = useState('')
  const [error, setError] = useState('')
  const [loading, setLoading] = useState(false)

  async function handleSubmit() {
    setError('')

    if (!name.trim()) {
      setError('Naziv razreda je obavezan.')
      return
    }

    const parsedYear = parseInt(year)
    if (!year || isNaN(parsedYear) || parsedYear <= 0) {
      setError('Unesite ispravan razred (npr. 1–8).')
      return
    }

    setLoading(true)
    try {
      await createClass(name.trim(), parsedYear)
      onSuccess()
      onClose()
    } catch (err: any) {
      setError(err.message || 'Greška pri kreiranju razreda.')
    } finally {
      setLoading(false)
    }
  }

  return (
    <div className="modal-overlay" onClick={onClose}>
      <div className="modal" onClick={(e) => e.stopPropagation()}>
        <h2>Create class</h2>

        <div className="form-group">
          <label htmlFor="class-name">Name</label>
          <input
            id="class-name"
            type="text"
            placeholder="III-4"
            value={name}
            onChange={(e) => setName(e.target.value)}
          />
        </div>

        <div className="form-group">
          <label htmlFor="class-year">Year</label>
          <input
            id="class-year"
            type="number"
            placeholder="3"
            value={year}
            onChange={(e) => setYear(e.target.value)}
          />
        </div>

        {error && <p className="error">{error}</p>}

        <div className="modal-actions">
          <button className="btn-secondary" onClick={onClose} disabled={loading}>
            Cancel
          </button>
          <button className="btn-primary" onClick={handleSubmit} disabled={loading}>
            {loading ? 'Creating...' : 'Create'}
          </button>
        </div>
      </div>
    </div>
  )
}
