import { useState, useEffect } from 'react'
import { useNavigate } from 'react-router-dom'
import CreateClassModal from '../forms/CreateClassModal'
import { getAllClasses, type SchoolClass } from '../../services/schoolService'

export default function Classes() {
    const [showModal, setShowModal] = useState(false)
    const [successMessage, setSuccessMessage] = useState('')
    const [classes, setClasses] = useState<SchoolClass[]>([])
    const navigate = useNavigate()

    async function fetchClasses() {
        try {
            const data = await getAllClasses()
            setClasses(data)
        } catch (err) {
            console.error(err)
        }
    }

    useEffect(() => {
        fetchClasses()
    }, [])

    function handleSuccess() {
        setSuccessMessage('Class created successfully!')
        setTimeout(() => setSuccessMessage(''), 3000)
        fetchClasses()
    }

    return (
        <div className="crud-page">
            <div style={{ display: 'flex', alignItems: 'center', justifyContent: 'center', gap: '16px' }}>
                <h1>Classes</h1>
                <button className="btn-primary" onClick={() => setShowModal(true)}>
                    Create Class
                </button>
            </div>

            {successMessage && <p className="success-message">{successMessage}</p>}

            <div className="class-grid">
                {classes.map((cls) => (
                    <div key={cls.ID} className="class-card" onClick={() => navigate(`/classes/${cls.ID}`)}>
                        <span>{cls.Name}</span>
                    </div>
                ))}
            </div>

            {showModal && (
                <CreateClassModal
                    onClose={() => setShowModal(false)}
                    onSuccess={handleSuccess}
                />
            )}
        </div>
    )
}