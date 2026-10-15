import { useEffect, useState } from 'react'
import { useParams } from 'react-router-dom'
import { getMemberById } from '../../services/authService'
import {
  getCertificatesByUserID,
  getMedicalRecordByUserID,
  type MedicalCertificate,
  type MedicalRecord,
} from '../../services/medicalService'

interface Member {
  firstName: string
  lastName: string
  email: string
  umcn: string
  role: string
}

export default function PatientDetails() {
  const { userID } = useParams()
  const [patient, setPatient] = useState<Member | null>(null)
  const [certificates, setCertificates] = useState<MedicalCertificate[]>([])
  const [medicalRecord, setMedicalRecord] = useState<MedicalRecord | null>(null)
  const [recordError, setRecordError] = useState<string | null>(null)

  useEffect(() => {
    if (!userID) return
    getMemberById(Number(userID)).then(setPatient).catch(console.error)
    getCertificatesByUserID(Number(userID)).then(setCertificates).catch(console.error)
    getMedicalRecordByUserID(Number(userID))
      .then(setMedicalRecord)
      .catch((err) => setRecordError(err.message))
  }, [userID])

  if (!patient) return <p>Loading...</p>

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
    <div className="profile-page">
      <div className="profile-container">

        <div className="profile-left">
          <img
            src="https://bootdey.com/img/Content/avatar/avatar7.png"
            alt="avatar"
            className="profile-avatar"
          />
          <h2>{patient.firstName} {patient.lastName}</h2>
          <p className="profile-role">{patient.role}</p>
        </div>

        <div style={{ display: 'flex', flexDirection: 'column', gap: '24px' }}>

          <div className="profile-right">
            <div className="profile-row"><span>First Name</span><span>{patient.firstName}</span></div>
            <div className="profile-row"><span>Last Name</span><span>{patient.lastName}</span></div>
            <div className="profile-row"><span>Email</span><span>{patient.email}</span></div>
            <div className="profile-row"><span>JMBG</span><span>{patient.umcn}</span></div>
          </div>

          <div className="profile-right">
            <h3 style={{ color: '#1f2937', marginBottom: '16px' }}>Certificates</h3>
            {certificates.length === 0 ? (
              <p style={{ color: '#6b7280', fontSize: '14px' }}>No certificates.</p>
            ) : (
              <div style={{ display: 'flex', flexDirection: 'column', gap: '10px' }}>
                {certificates.map(c => (
                  <div
                    key={c.ID}
                    style={{
                      background: '#f0f9ff',
                      borderRadius: '12px',
                      padding: '14px 18px',
                      display: 'flex',
                      flexDirection: 'column',
                      gap: '4px',
                    }}
                  >
                    <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
                      <span style={{ fontWeight: 700, color: '#1f2937', fontSize: '14px' }}>
                        {certTypeLabel(c.Type)}
                      </span>
                      <span style={{ fontSize: '13px', color: '#6b7280' }}>
                        {new Date(c.Date).toLocaleDateString()}
                      </span>
                    </div>
                    {c.Note && (
                      <span style={{ fontSize: '13px', color: '#6b7280' }}>{c.Note}</span>
                    )}
                  </div>
                ))}
              </div>
            )}
          </div>

          <div className="profile-right">
            <h3 style={{ color: '#1f2937', marginBottom: '16px' }}>Medical Record</h3>
            {recordError ? (
              <p style={{ color: '#6b7280', fontSize: '14px' }}>No medical record yet.</p>
            ) : !medicalRecord ? (
              <p style={{ color: '#6b7280', fontSize: '14px' }}>Loading...</p>
            ) : medicalRecord.Examinations.length === 0 ? (
              <p style={{ color: '#6b7280', fontSize: '14px' }}>No examinations recorded.</p>
            ) : (
              <div style={{ display: 'flex', flexDirection: 'column', gap: '10px' }}>
                {medicalRecord.Examinations.map(e => (
                  <div
                    key={e.ID}
                    style={{
                      background: '#f0fdf4',
                      borderRadius: '12px',
                      padding: '14px 18px',
                      display: 'flex',
                      flexDirection: 'column',
                      gap: '4px',
                    }}
                  >
                    <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
                      <span style={{ fontWeight: 700, color: '#1f2937', fontSize: '14px' }}>
                        {e.Diagnosis}
                      </span>
                      <span style={{ fontSize: '13px', color: '#6b7280' }}>
                        {new Date(e.CreatedAt).toLocaleDateString()}
                      </span>
                    </div>
                    <span style={{ fontSize: '13px', color: '#6b7280' }}>Therapy: {e.Therapy}</span>
                    {e.Note && (
                      <span style={{ fontSize: '13px', color: '#6b7280' }}>{e.Note}</span>
                    )}
                  </div>
                ))}
              </div>
            )}
          </div>

        </div>

      </div>
    </div>
  )
}