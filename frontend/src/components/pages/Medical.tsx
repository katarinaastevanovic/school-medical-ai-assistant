import { useEffect, useState } from 'react'
import {
  getMyExaminationRequests,
  getMyCertificates,
  getMyMedicalRecord,
  type MyExaminationRequest,
  type MyCertificate,
  type MedicalRecord,
} from '../../services/medicalService'
import CreateExaminationRequestModal from '../forms/CreateExaminationRequestModal'

export default function Medical() {
  const [requests, setRequests] = useState<MyExaminationRequest[]>([])
  const [certificates, setCertificates] = useState<MyCertificate[]>([])
  const [medicalRecord, setMedicalRecord] = useState<MedicalRecord | null>(null)
  const [showModal, setShowModal] = useState(false)
  const [activeTab, setActiveTab] = useState<'requests' | 'certificates' | 'record'>('requests')

  async function fetchData() {
    try {
      const [reqs, certs] = await Promise.all([
        getMyExaminationRequests(),
        getMyCertificates(),
      ])
      setRequests(reqs)
      setCertificates(certs)
    } catch (err) {
      console.error(err)
    }

    try {
      const record = await getMyMedicalRecord()
      setMedicalRecord(record)
    } catch (err) {
      console.error(err)
      setMedicalRecord(null)
    }
  }

  useEffect(() => {
    fetchData()
  }, [])

  const statusStyle = (status: string) => {
    switch (status) {
      case 'APPROVED': return { background: '#dcfce7', color: '#166534' }
      case 'REJECTED': return { background: '#fee2e2', color: '#991b1b' }
      default: return { background: '#fef9c3', color: '#854d0e' }
    }
  }

  const certTypeLabel = (type: string) => {
    switch (type) {
      case 'REGULAR': return 'Regular'
      case 'PE': return 'Physical Education'
      case 'EXCURSION': return 'Excursion'
      case 'SICKNESS': return 'Sickness'
      default: return type
    }
  }

  return (
    <div className="crud-page" style={{ maxWidth: '800px' }}>
      <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
        <h1>Medical</h1>
        {activeTab === 'requests' && (
          <button className="btn-primary" onClick={() => setShowModal(true)}>
            New Request
          </button>
        )}
      </div>

      <div className="tabs">
        <button
          className={`tab-btn ${activeTab === 'requests' ? 'active' : ''}`}
          onClick={() => setActiveTab('requests')}
        >
          Requests
        </button>
        <button
          className={`tab-btn ${activeTab === 'certificates' ? 'active' : ''}`}
          onClick={() => setActiveTab('certificates')}
        >
          Certificates
        </button>
        <button
          className={`tab-btn ${activeTab === 'record' ? 'active' : ''}`}
          onClick={() => setActiveTab('record')}
        >
          Medical Record
        </button>
      </div>

      {activeTab === 'requests' && (
        <div className="crud-list">
          {requests.length === 0 ? (
            <p style={{ color: '#6b7280', fontSize: '14px' }}>No requests yet.</p>
          ) : (
            requests.map(req => (
              <div key={req.ID} className="crud-item">
                <span style={{ fontWeight: 600, color: '#1f2937' }}>{req.Type}</span>
                <span style={{
                  ...statusStyle(req.Status),
                  borderRadius: '20px',
                  padding: '4px 12px',
                  fontSize: '13px',
                  fontWeight: 600,
                }}>
                  {req.Status}
                </span>
              </div>
            ))
          )}
        </div>
      )}

      {activeTab === 'certificates' && (
        <div className="crud-list">
          {certificates.length === 0 ? (
            <p style={{ color: '#6b7280', fontSize: '14px' }}>No certificates yet.</p>
          ) : (
            certificates.map(cert => (
              <div key={cert.ID} className="crud-item" style={{ flexDirection: 'column', alignItems: 'flex-start', gap: '8px' }}>
                <div style={{ display: 'flex', justifyContent: 'space-between', width: '100%' }}>
                  <span style={{ fontWeight: 700, color: '#1f2937' }}>{certTypeLabel(cert.Type)}</span>
                  <span style={{ color: '#6b7280', fontSize: '13px' }}>
                    {new Date(cert.Date).toLocaleDateString()}
                  </span>
                </div>
                {cert.Note && (
                  <span style={{ color: '#6b7280', fontSize: '13px' }}>{cert.Note}</span>
                )}
              </div>
            ))
          )}
        </div>
      )}

      {activeTab === 'record' && (
        <div className="crud-list">
          {!medicalRecord || !medicalRecord.Examinations || medicalRecord.Examinations.length === 0 ? (
            <p style={{ color: '#6b7280', fontSize: '14px' }}>No examinations recorded.</p>
          ) : (
            medicalRecord.Examinations.map(exam => (
              <div key={exam.ID} className="crud-item" style={{ flexDirection: 'column', alignItems: 'flex-start', gap: '8px' }}>
                <div style={{ display: 'flex', justifyContent: 'space-between', width: '100%' }}>
                  <span style={{ fontWeight: 700, color: '#1f2937' }}>{exam.Diagnosis}</span>
                  <span style={{ color: '#6b7280', fontSize: '13px' }}>
                    {new Date(exam.CreatedAt).toLocaleDateString()}
                  </span>
                </div>
                <span style={{ color: '#6b7280', fontSize: '13px' }}>Therapy: {exam.Therapy}</span>
                {exam.Note && (
                  <span style={{ color: '#6b7280', fontSize: '13px' }}>{exam.Note}</span>
                )}
              </div>
            ))
          )}
        </div>
      )}

      {showModal && (
        <CreateExaminationRequestModal
          onClose={() => setShowModal(false)}
          onSuccess={() => { setShowModal(false); fetchData() }}
        />
      )}
    </div>
  )
}