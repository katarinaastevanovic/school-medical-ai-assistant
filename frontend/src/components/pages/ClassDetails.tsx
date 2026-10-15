import { useEffect, useState } from 'react'
import { useParams } from 'react-router-dom'
import { getStudentsFromClass, removeStudentFromClass, type Student } from '../../services/schoolService'
import { getMemberById } from '../../services/authService'
import AddStudentsModal from '../forms/AddStudentsModal'
import AssignSubjectModal from '../forms/AssignSubjectModal'
import { getSubjectsByClass, type ClassSubject } from '../../services/schoolService'


interface StudentWithInfo {
    ID: number
    UserID: number
    firstName: string
    lastName: string
}

export default function ClassDetails() {
    const { id } = useParams()
    const [students, setStudents] = useState<StudentWithInfo[]>([])
    const [showModal, setShowModal] = useState(false)
    const [showAssignModal, setShowAssignModal] = useState(false)
    const [classSubjects, setClassSubjects] = useState<ClassSubject[]>([])
    const [activeTab, setActiveTab] = useState<'students' | 'subjects'>('students')

    async function fetchStudents() {
        if (!id) return

        getStudentsFromClass(Number(id))
            .then(async (data) => {
                const enriched = await Promise.all(
                    data.map(async (s) => {
                        const member = await getMemberById(s.UserID)
                        return {
                            ID: s.ID,
                            UserID: s.UserID,
                            firstName: member.firstName,
                            lastName: member.lastName,
                        }
                    })
                )
                setStudents(enriched)
            })
            .catch(err => console.error(err))
    }
    async function fetchSubjectsForClass() {
        if (!id) return
        try {
            const data = await getSubjectsByClass(Number(id))
            setClassSubjects(data)
        } catch (err) {
            console.error(err)
        }
    }

    useEffect(() => {
        fetchStudents()
        fetchSubjectsForClass()
    }, [id])

    async function handleRemove(studentID: number) {
        if (!id) return
        try {
            await removeStudentFromClass(Number(id), studentID)
            fetchStudents()
        } catch (err) {
            console.error(err)
        }
    }

    return (
    <div className="crud-page">
        <div className="tabs">
            <button
                className={`tab-btn ${activeTab === 'students' ? 'active' : ''}`}
                onClick={() => setActiveTab('students')}
            >
                Students
            </button>
            <button
                className={`tab-btn ${activeTab === 'subjects' ? 'active' : ''}`}
                onClick={() => setActiveTab('subjects')}
            >
                Subjects
            </button>
        </div>

        {activeTab === 'students' && (
            <>
                <button className="btn-primary" style={{ width: 'fit-content' }} onClick={() => setShowModal(true)}>
                    + Add Students
                </button>

                <div className="class-grid">
                    {students.map((s) => (
                        <div key={s.ID} className="student-card">
                            <button className="student-remove-btn" onClick={() => handleRemove(s.ID)}>✕</button>
                            <img
                                src="https://bootdey.com/img/Content/avatar/avatar7.png"
                                alt="avatar"
                                className="student-avatar"
                            />
                            <span>{s.firstName} {s.lastName}</span>
                        </div>
                    ))}
                </div>
            </>
        )}

        {activeTab === 'subjects' && (
            <>
                <button className="btn-primary" style={{ width: 'fit-content' }} onClick={() => setShowAssignModal(true)}>
                    Assign Subject
                </button>

                <div className="class-grid">
                    {classSubjects.map(cs => (
                        <div key={cs.id} className="subject-card">
                            <span className="subject-card-name">{cs.subjectName}</span>
                            <span className="subject-card-teacher">{cs.firstName} {cs.lastName}</span>
                        </div>
                    ))}
                </div>
            </>
        )}

        {showAssignModal && (
            <AssignSubjectModal
                classID={Number(id)}
                onClose={() => setShowAssignModal(false)}
                onSuccess={fetchSubjectsForClass}
            />
        )}

        {showModal && (
            <AddStudentsModal
                classID={Number(id)}
                onClose={() => setShowModal(false)}
                onSuccess={fetchStudents}
            />
        )}
    </div>
)
}