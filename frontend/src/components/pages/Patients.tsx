import { useEffect, useState } from 'react'
import { getAllPatients, type Patient } from '../../services/medicalService'
import { getMemberById } from '../../services/authService'
import CreateCertificateModal from '../forms/CreateCertificateModal'
import { useNavigate } from 'react-router-dom'

interface EnrichedPatient extends Patient {
  firstName: string
  lastName: string
  email: string
}

export default function Patients() {
  const [patients, setPatients] = useState<EnrichedPatient[]>([])
  const [selectedPatient, setSelectedPatient] = useState<EnrichedPatient | null>(null)
  const [search, setSearch] = useState('')
  const [sort, setSort] = useState<'asc' | 'desc'>('asc')
  const navigate = useNavigate()


  useEffect(() => {
    getAllPatients()
      .then(async (data) => {
        const enriched = await Promise.all(
          data.map(async (p) => {
            try {
              const member = await getMemberById(p.UserID)
              return { ...p, firstName: member.firstName, lastName: member.lastName, email: member.email }
            } catch {
              return { ...p, firstName: '?', lastName: '?', email: '?' }
            }
          })
        )
        setPatients(enriched)
      })
      .catch(console.error)
  }, [])

  const filtered = patients
    .filter(p =>
      `${p.firstName} ${p.lastName} ${p.email}`
        .toLowerCase()
        .includes(search.toLowerCase())
    )
    .sort((a, b) => {
      const nameA = `${a.firstName} ${a.lastName}`
      const nameB = `${b.firstName} ${b.lastName}`
      return sort === 'asc'
        ? nameA.localeCompare(nameB)
        : nameB.localeCompare(nameA)
    })

  return (
    <div className="crud-page" style={{ maxWidth: '900px' }}>
      <h1>Patients</h1>

      <div style={{ display: 'flex', gap: '12px' }}>
        <input
          type="text"
          placeholder="Search by name or email..."
          value={search}
          onChange={e => setSearch(e.target.value)}
          style={{
            flex: 1,
            padding: '12px 15px',
            borderRadius: '12px',
            border: '1px solid #d6dce5',
            background: '#f9fafb',
            fontSize: '15px',
            color: '#1f2937',
          }}
        />
        <select
          className="role-select"
          style={{ width: 'fit-content' }}
          value={sort}
          onChange={e => setSort(e.target.value as 'asc' | 'desc')}
        >
          <option value="asc">A → Z</option>
          <option value="desc">Z → A</option>
        </select>
      </div>

      <div className="class-grid">
        {filtered.map(p => (
          <div
            key={p.ID}
            className="student-card"
            style={{ cursor: 'pointer', minWidth: '200px' }}
            onClick={() => navigate(`/patients/${p.UserID}`)}
          >
            <img
              src="https://bootdey.com/img/Content/avatar/avatar7.png"
              alt="avatar"
              className="student-avatar"
            />
            <span>{p.firstName} {p.lastName}</span>
            <span style={{ fontSize: '12px', color: '#6b7280' }}>{p.email}</span>
            <button
              className="btn-primary"
              style={{ marginTop: '8px', fontSize: '13px', padding: '8px 16px' }}
              onClick={e => { e.stopPropagation(); setSelectedPatient(p) }}
            >
              Issue Certificate
            </button>
          </div>
        ))}
      </div>

      {selectedPatient && (
        <CreateCertificateModal
          patientUserID={selectedPatient.UserID}
          patientName={`${selectedPatient.firstName} ${selectedPatient.lastName}`}
          onClose={() => setSelectedPatient(null)}
          onSuccess={() => setSelectedPatient(null)}
        />
      )}
    </div>
  )
}