import { useEffect, useState } from "react"
import { getMyClass, type StudentSubject } from "../../services/schoolService"
import { useNavigate } from "react-router-dom"

export default function StudentSubjects() {

    const [subjects, setSubjects] = useState<StudentSubject[]>([])
    const navigate = useNavigate()

    useEffect(() => {

        getMyClass()
            .then(data => setSubjects(data.subjects))
            .catch(console.error)

    }, [])

    return (
        <div className="crud-page">

            <h1>My Subjects</h1>

            <div className="class-grid">

    {subjects.map(subject => (

        <div
            key={subject.subjectID}
            className="student-card"
            style={{ cursor: "pointer" }}
            onClick={() =>navigate(`/student-subjects/${subject.subjectID}`, {state: subject})}>
            <h3>{subject.subjectName}</h3>
            <p>Teacher: {subject.firstName} {subject.lastName}</p>
        </div>
    ))}
</div>
        </div>
    )
}