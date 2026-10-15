import { useState, useEffect } from 'react'
import { getAllSubjects, createSubject, updateSubject, deleteSubject, type Subject } from '../../services/schoolService'
import SubjectTeachersModal from '../forms/SubjectTeachersModal'


export default function Subjects() {
    const [subjects, setSubjects] = useState<Subject[]>([])
    const [newName, setNewName] = useState('')
    const [editingID, setEditingID] = useState<number | null>(null)
    const [editingName, setEditingName] = useState('')
    const [error, setError] = useState('')
    const [teachersModal, setTeachersModal] = useState<{ id: number; name: string } | null>(null)


    async function fetchSubjects() {
        try {
            const data = await getAllSubjects()
            setSubjects(data)
        } catch (err) {
            console.error(err)
        }
    }

    useEffect(() => {
        fetchSubjects()
    }, [])

    async function handleCreate() {
        setError('')
        if (!newName.trim()) {
            setError('Subject name is required.')
            return
        }
        try {
            await createSubject(newName.trim())
            setNewName('')
            fetchSubjects()
        } catch (err: any) {
            setError(err.message)
        }
    }

    async function handleUpdate(id: number) {
        setError('')
        if (!editingName.trim()) {
            setError('Subject name is required.')
            return
        }
        try {
            await updateSubject(id, editingName.trim())
            setEditingID(null)
            setEditingName('')
            fetchSubjects()
        } catch (err: any) {
            setError(err.message)
        }
    }

    async function handleDelete(id: number) {
        try {
            await deleteSubject(id)
            fetchSubjects()
        } catch (err) {
            console.error(err)
        }
    }

    return (
        <div className="crud-page">
            <h1>Subjects</h1>

            <div className="crud-add">
                <input
                    type="text"
                    placeholder="New subject name"
                    value={newName}
                    onChange={(e) => setNewName(e.target.value)}
                />
                <button className="btn-primary" onClick={handleCreate}>Add</button>
            </div>

            {error && <p className="error">{error}</p>}

            <div className="crud-list">
                {subjects.map((s) => (
                    <div key={s.ID} className="crud-item">
                        {editingID === s.ID ? (
                            <>
                                <input
                                    type="text"
                                    value={editingName}
                                    onChange={(e) => setEditingName(e.target.value)}
                                />
                                <div className="crud-item-actions">
                                    <button className="btn-primary" onClick={() => handleUpdate(s.ID)}>Save</button>
                                    <button className="btn-secondary" onClick={() => setEditingID(null)}>Cancel</button>
                                </div>
                            </>
                        ) : (
                            <>
                                <span>{s.Name}</span>
                                <div className="crud-item-actions">
                                    <button className="btn-secondary" onClick={() => { setEditingID(s.ID); setEditingName(s.Name) }}>Edit</button>
                                    <button className="btn-danger" onClick={() => handleDelete(s.ID)}>Delete</button>
                                    <button className="btn-secondary" onClick={() => setTeachersModal({ id: s.ID, name: s.Name })}>
                                        Teachers
                                    </button>
                                </div>
                            </>
                        )}
                    </div>
                ))}
            </div>
            {teachersModal && (
                <SubjectTeachersModal
                    subjectID={teachersModal.id}
                    subjectName={teachersModal.name}
                    onClose={() => setTeachersModal(null)}
                />
            )}
        </div>
    )
}