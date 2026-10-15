import { useEffect, useState } from 'react'
import { getExaminationRequests } from '../../services/medicalService'
import { getMemberById } from '../../services/authService'
import CreateExaminationModal from '../forms/CreateExaminationModal'
import { updateRequestStatus } from '../../services/medicalService'

interface ExaminationRequest {
  id: number
  patientID: number
  patientUserID: number
  doctorID: number
  type: string
  status: string
  medicalRecordID: number
}

interface EnrichedRequest extends ExaminationRequest {
  firstName: string
  lastName: string
}

export default function ExaminationRequests() {
  const [requests, setRequests] = useState<EnrichedRequest[]>([])
  const [selectedRequest, setSelectedRequest] = useState<EnrichedRequest | null>(null)

  async function fetchRequests() {
    try {
      const data = await getExaminationRequests()
      const enriched = await Promise.all(
        data.map(async (req: ExaminationRequest) => {
          try {
            const member = await getMemberById(req.patientUserID)
            return { ...req, firstName: member.firstName, lastName: member.lastName }
          } catch {
            return { ...req, firstName: '?', lastName: '?' }
          }
        })
      )
      setRequests(enriched)
    } catch (err) {
      console.error(err)
    }
  }

  useEffect(() => {
    fetchRequests()
  }, [])

  async function handleStatus(requestID: number, status: string) {
    try {
      await updateRequestStatus(requestID, status)
      fetchRequests()
    } catch (err) {
      console.error(err)
    }
  }

  const statusStyle = (status: string) => {
    switch (status) {
      case 'APPROVED': return { background: '#dcfce7', color: '#166534' }
      case 'REJECTED': return { background: '#fee2e2', color: '#991b1b' }
      default: return { background: '#fef9c3', color: '#854d0e' }
    }
  }

  return (
    <div className="crud-page" style={{ maxWidth: '900px' }}>
      <h1>Examination Requests</h1>

      <div className="crud-list">
        {requests.map(req => (
          <div key={req.id} className="crud-item" style={{ flexDirection: 'row', alignItems: 'center', gap: '12px' }}>
            <div style={{ flex: 1, display: 'flex', alignItems: 'center', gap: '12px' }}>
              <span style={{ fontWeight: 700, color: '#1f2937', fontSize: '15px' }}>
                {req.firstName} {req.lastName}
              </span>
              <span style={{ color: '#6b7280', fontSize: '13px' }}>{req.type}</span>
              <span style={{
                ...statusStyle(req.status),
                borderRadius: '20px',
                padding: '4px 12px',
                fontSize: '13px',
                fontWeight: 600,
              }}>
                {req.status}
              </span>
            </div>

            <div style={{ display: 'flex', gap: '8px' }}>
              {req.status === 'REQUESTED' && (
                <>
                  <button
                    style={{
                      border: 'none',
                      borderRadius: '10px',
                      padding: '8px 18px',
                      background: '#dcfce7',
                      color: '#166534',
                      fontSize: '14px',
                      fontWeight: 600,
                      cursor: 'pointer',
                    }}
                    onClick={() => handleStatus(req.id, 'APPROVED')}
                  >
                    Approve
                  </button>
                  <button className="btn-danger" onClick={() => handleStatus(req.id, 'REJECTED')}>
                    Reject
                  </button>
                </>
              )}

              {req.status === 'APPROVED' && (
                <button className="btn-secondary" onClick={() => setSelectedRequest(req)}>
                  Add Examination
                </button>
              )}

              {req.status === 'REJECTED' && (
                <div style={{ width: '160px' }} />
              )}
            </div>
          </div>
        ))}
      </div>

      {selectedRequest && (
        <CreateExaminationModal
          requestID={selectedRequest.id}
          patientName={`${selectedRequest.firstName} ${selectedRequest.lastName}`}
          onClose={() => setSelectedRequest(null)}
          onSuccess={() => { setSelectedRequest(null); fetchRequests() }}
        />
      )}
    </div>
  )
}