import { useEffect, useState } from 'react'
import { getMemberById, updateMember, getAvatarByUmcn } from '../../services/authService'
import { getStudentOverview, getStudentAbsenceOverview, type StudentOverview, type AbsenceOverview, getMyClass, type StudentSummaryResponse } from '../../services/schoolService'
import { useNavigate } from 'react-router-dom'
import { getStudentSummary } from '../../services/schoolService'

interface Member {
  id: number
  umcn: string
  firstName: string
  lastName: string
  email: string
  role: string
}

export default function Profile() {
  const [user, setUser] = useState<Member | null>(null)
  const [isEditing, setIsEditing] = useState(false)
  const [editData, setEditData] = useState({
    firstName: '',
    lastName: '',
    email: '',
    password: '',
  })
  const [saving, setSaving] = useState(false)
  const [overview, setOverview] = useState<StudentOverview[]>([])
  const [absenceOverview, setAbsenceOverview] = useState<AbsenceOverview[]>([])
  const navigate = useNavigate()
  const [myClass, setMyClass] = useState<{ className: number } | null>(null)
  const [summary, setSummary] = useState<StudentSummaryResponse | null>(null)

  const startEditing = () => {
    if (!user) return
    setEditData({
      firstName: user.firstName,
      lastName: user.lastName,
      email: user.email,
      password: '',
    })
    setIsEditing(true)
  }

  const cancelEditing = () => {
    setIsEditing(false)
  }

  const handleSave = async () => {
    if (!user) return
    try {
      setSaving(true)
      const dataToUpdate: {
        firstName: string
        lastName: string
        email: string
        password?: string
      } = {
        firstName: editData.firstName,
        lastName: editData.lastName,
        email: editData.email,
      }
      if (editData.password.trim() !== '') {
        dataToUpdate.password = editData.password
      }
      const updatedUser = await updateMember(user.id, dataToUpdate)
      setUser(updatedUser)
      localStorage.setItem('user', JSON.stringify(updatedUser))
      setIsEditing(false)
    } catch (error) {
      console.error(error)
      alert(error instanceof Error ? error.message : 'Failed to update profile')
    } finally {
      setSaving(false)
    }
  }

  useEffect(() => {
    const storedUser = localStorage.getItem('user')
    if (!storedUser) return
    const currentUser = JSON.parse(storedUser)

    getStudentSummary()
      .then(setSummary)
      .catch(console.error)

    getMemberById(currentUser.id)
      .then(data => {
        setUser(data)
        if (data.role === 'STUDENT') {
          getStudentOverview().then(setOverview).catch(console.error)
          getStudentAbsenceOverview().then(setAbsenceOverview).catch(console.error)
          getMyClass().then(data => setMyClass({ className: data.className })).catch(console.error)
        }
      })
      .catch(err => console.error(err))
  }, [])

  if (!user) return <p>Loading...</p>

  return (
    <div className="profile-page" style={{ fontSize: '14px' }}>
      <div className="profile-container" style={{ gap: '16px' }}>

        <div className="profile-left" style={{ padding: '12px' }}>
          <img
            src={getAvatarByUmcn(user.umcn)}
            alt="avatar"
            className="profile-avatar"
            style={{ width: '64px', height: '64px' }}
          />
          <h2 style={{ margin: '4px 0 0', fontSize: '16px' }}>{user.firstName} {user.lastName}</h2>
          <p className="profile-role" style={{ margin: '2px 0', fontSize: '12px' }}>{user.role}</p>

          {myClass && (
            <div style={{
              marginTop: '10px',
              background: '#f1f4f8',
              borderRadius: '8px',
              padding: '6px 12px',
              fontSize: '12px',
              fontWeight: 600,
              color: '#1e3a5f'
            }}>
              Class: {myClass.className}
            </div>
          )}
          {user.role === 'STUDENT' && summary && (
            <div style={{
              marginTop: '12px',
              borderTop: '1px solid #e5e7eb',
              paddingTop: '10px',
              width: '100%'
            }}>
              <h3 style={{
                marginBottom: '8px',
                color: '#1f2937',
                textAlign: 'center',
                fontSize: '13px'
              }}>
                Statistics
              </h3>

              <div style={{
                background: '#f9fafb',
                borderRadius: '8px',
                padding: '6px',
                marginBottom: '6px',
                textAlign: 'center'
              }}>
                <div style={{ fontSize: '10px', color: '#6b7280', marginBottom: '2px' }}>
                  Average grade
                </div>
                <div style={{ fontSize: '20px', fontWeight: 700, color: '#2563eb' }}>
                  {summary.averageGrade.toFixed(2)}
                </div>
              </div>

              <div style={{ display: 'flex', gap: '6px' }}>

                <div style={{ background: '#dcfce7', borderRadius: '8px', padding: '6px', textAlign: 'center', flex: 1 }}>
                  <div style={{ fontSize: '10px', color: '#166534' }}>Excused</div>
                  <div style={{ fontSize: '16px', fontWeight: 700, color: '#166534' }}>
                    {summary.excused}
                  </div>
                </div>

                <div style={{ background: '#fee2e2', borderRadius: '8px', padding: '6px', textAlign: 'center', flex: 1 }}>
                  <div style={{ fontSize: '10px', color: '#991b1b' }}>Unexcused</div>
                  <div style={{ fontSize: '16px', fontWeight: 700, color: '#991b1b' }}>
                    {summary.unexcused}
                  </div>
                </div>

                <div style={{ background: '#fef9c3', borderRadius: '8px', padding: '6px', textAlign: 'center', flex: 1 }}>
                  <div style={{ fontSize: '10px', color: '#854d0e' }}>Pending</div>
                  <div style={{ fontSize: '16px', fontWeight: 700, color: '#854d0e' }}>
                    {summary.pending}
                  </div>
                </div>

              </div>
            </div>
          )}
        </div>

        <div style={{ display: 'flex', flexDirection: 'column', gap: '16px' }}>

          <div className="profile-right" style={{ padding: '12px' }}>
            <div
              style={{
                display: 'flex',
                justifyContent: 'space-between',
                alignItems: 'center',
                marginBottom: '6px',
              }}
            >
              <h3 style={{ color: '#1f2937', margin: 0, fontSize: '14px' }}>Profile</h3>

              {!isEditing && (
                <button
                  onClick={startEditing}
                  title="Edit profile"
                  style={{
                    border: 'none',
                    background: 'transparent',
                    cursor: 'pointer',
                    fontSize: '16px',
                    padding: '2px 6px',
                  }}
                >
                  ✎
                </button>
              )}
            </div>

            <div className="profile-row" style={{ padding: '4px 0', fontSize: '13px' }}>
              <span>First Name</span>
              {isEditing ? (
                <input
                  type="text"
                  className="profile-edit-input"
                  value={editData.firstName}
                  onChange={e => setEditData({ ...editData, firstName: e.target.value })}
                />
              ) : (
                <span>{user.firstName}</span>
              )}
            </div>

            <div className="profile-row" style={{ padding: '4px 0', fontSize: '13px' }}>
              <span>Last Name</span>
              {isEditing ? (
                <input
                  type="text"
                  className="profile-edit-input"
                  value={editData.lastName}
                  onChange={e => setEditData({ ...editData, lastName: e.target.value })}
                />
              ) : (
                <span>{user.lastName}</span>
              )}
            </div>

            <div className="profile-row" style={{ padding: '4px 0', fontSize: '13px' }}>
              <span>Email</span>
              {isEditing ? (
                <input
                  type="email"
                  className="profile-edit-input"
                  value={editData.email}
                  onChange={e => setEditData({ ...editData, email: e.target.value })}
                />
              ) : (
                <span>{user.email}</span>
              )}
            </div>

            <div className="profile-row" style={{ padding: '4px 0', fontSize: '13px' }}>
              <span>JMBG</span>
              <span>{user.umcn}</span>
            </div>

            <div className="profile-row" style={{ padding: '4px 0', fontSize: '13px' }}>
              <span>Role</span>
              <span>{user.role}</span>
            </div>

            {isEditing && (
              <>
                <div className="profile-row" style={{ padding: '4px 0', fontSize: '13px' }}>
                  <span>New Password</span>
                  <input
                    type="password"
                    className="profile-edit-input"
                    placeholder="Leave empty to keep current"
                    value={editData.password}
                    onChange={e => setEditData({ ...editData, password: e.target.value })}
                  />
                </div>

                <div className="profile-edit-actions" style={{ marginTop: '8px' }}>
                  <button
                    className="profile-cancel-btn"
                    onClick={cancelEditing}
                    disabled={saving}
                  >
                    Cancel
                  </button>
                  <button
                    className="profile-save-btn"
                    onClick={handleSave}
                    disabled={saving}
                  >
                    {saving ? 'Saving...' : 'Save'}
                  </button>
                </div>
              </>
            )}
          </div>

          {user.role === 'STUDENT' && (
            <div style={{ display: 'flex', gap: '16px' }}>

              {overview.length > 0 && (
                <div className="profile-right" style={{ flex: 1, padding: '12px', maxHeight: '40vh', overflowY: 'auto' }}>
                  <h3 style={{ color: '#1f2937', marginBottom: '8px', fontSize: '14px' }}>Grades</h3>
                  <div style={{ display: 'flex', flexDirection: 'column', gap: '6px' }}>
                    {overview.map(s => (
                      <div
                        key={s.subjectID}
                        style={{
                          display: 'flex',
                          justifyContent: 'space-between',
                          alignItems: 'center',
                          background: '#f9fafb',
                          borderRadius: '10px',
                          padding: '6px 12px',
                          transition: 'background 0.15s',
                        }}
                        onMouseEnter={e => (e.currentTarget.style.background = '#f1f4f8')}
                        onMouseLeave={e => (e.currentTarget.style.background = '#f9fafb')}
                      >
                        <span style={{ fontWeight: 600, color: '#1f2937', fontSize: '13px' }}>{s.subjectName}</span>
                        <span style={{
                          fontWeight: 700,
                          fontSize: '12px',
                          padding: '2px 10px',
                          borderRadius: '20px',
                          background: s.average >= 4 ? '#dcfce7' : s.average >= 3 ? '#fef9c3' : s.average > 0 ? '#fee2e2' : '#f3f4f6',
                          color: s.average >= 4 ? '#166534' : s.average >= 3 ? '#854d0e' : s.average > 0 ? '#991b1b' : '#9ca3af'
                        }}>
                          {s.average > 0 ? s.average.toFixed(2) : '—'}
                        </span>
                      </div>
                    ))}
                  </div>
                </div>
              )}

              {absenceOverview.length > 0 && (
                <div className="profile-right" style={{ flex: 1, padding: '12px', maxHeight: '40vh', overflowY: 'auto' }}>
                  <h3 style={{ color: '#1f2937', marginBottom: '8px', fontSize: '14px' }}>Absences</h3>
                  <div style={{ display: 'flex', flexDirection: 'column', gap: '6px' }}>
                    {absenceOverview.map(s => (
                      <div
                        key={s.subjectID}
                        style={{
                          background: '#f9fafb',
                          borderRadius: '10px',
                          padding: '6px 12px',
                        }}
                      >
                        <span style={{ fontWeight: 600, color: '#1f2937', fontSize: '13px', display: 'block', marginBottom: '4px' }}>
                          {s.subjectName}
                        </span>
                        <div style={{ display: 'flex', gap: '6px' }}>
                          <span style={{ background: '#dcfce7', color: '#166534', borderRadius: '20px', padding: '2px 8px', fontSize: '11px', fontWeight: 600 }}>
                            {s.excused} excused
                          </span>
                          <span style={{ background: '#fee2e2', color: '#991b1b', borderRadius: '20px', padding: '2px 8px', fontSize: '11px', fontWeight: 600 }}>
                            {s.unexcused} unexcused
                          </span>
                          <span style={{ background: '#fef9c3', color: '#854d0e', borderRadius: '20px', padding: '2px 8px', fontSize: '11px', fontWeight: 600 }}>
                            {s.pending} pending
                          </span>
                        </div>
                      </div>
                    ))}
                  </div>
                </div>
              )}

            </div>
          )}

        </div>

      </div>
    </div>
  )
}