import { useEffect, useState } from "react"
import { useParams } from "react-router-dom"
import { getAbsenceStats, getAverageGrade, getGradesAndAbsences, updateAbsenceStatus, type AbsenceStatsResponse, type AverageGradeResponse, type GradesAndAbsences } from "../../services/schoolService"
import { getCertificatesByUserID, type MedicalCertificate } from '../../services/medicalService'
import { getMemberById } from "../../services/authService"
import AddAbsenceModal from "../forms/AddAbsenceModal"
import AddGradeModal from "../forms/AddGradeModal"

interface Member {
    firstName: string
    lastName: string
    email: string
    role: string
}

export default function StudentGrades() {
    const { userID, subjectID } = useParams()
    const [data, setData] = useState<GradesAndAbsences | null>(null)
    const [student, setStudent] = useState<Member | null>(null)
    const [showGradeModal, setShowGradeModal] = useState(false)
    const [showAbsenceModal, setShowAbsenceModal] = useState(false)
    const [average, setAverage] = useState<AverageGradeResponse | null>(null)
    const [absenceStats, setAbsenceStats] = useState<AbsenceStatsResponse | null>(null)
    const [certificates, setCertificates] = useState<MedicalCertificate[]>([])


    async function loadData() {
        if (!userID || !subjectID) return

        const [grades, averageGrade, stats] = await Promise.all([
            getGradesAndAbsences(Number(userID), Number(subjectID)),
            getAverageGrade(Number(userID), Number(subjectID)),
            getAbsenceStats(Number(userID), Number(subjectID))
        ])

        setData(grades)
        setAverage(averageGrade)
        setAbsenceStats(stats)
    }

    async function handleStatusChange(absenceID: number, status: string) {
        try {
            await updateAbsenceStatus(absenceID, status)
            loadData()
        } catch (err) {
            console.error(err)
        }
    }

    useEffect(() => {
        if (!userID) return
        getMemberById(Number(userID)).then(setStudent).catch(console.error)
        getCertificatesByUserID(Number(userID)).then(setCertificates).catch(console.error)
        loadData()
    }, [userID, subjectID])

    if (!data || !student) return <p>Loading...</p>


    return (
        <div className="profile-page-student">
            <div className="profile-container">

                <div className="profile-left">
                    <img
                        src="https://bootdey.com/img/Content/avatar/avatar7.png"
                        alt="avatar"
                        className="profile-avatar"
                    />
                    <h2>{student.firstName} {student.lastName}</h2>
                    <p className="profile-role">{student.role}</p>

                    <div style={{ display: 'flex', flexDirection: 'column', gap: '10px', marginTop: '16px' }}>
                        <button className="btn-primary" onClick={() => setShowGradeModal(true)}>
                            Add Grade
                        </button>
                        <button className="btn-secondary" onClick={() => setShowAbsenceModal(true)}>
                            Add Absence
                        </button>
                    </div>

                    <div style={{ marginTop: "24px", borderTop: "1px solid #e5e7eb", paddingTop: "16px", width: "100%" }}>
                        <h3 style={{ marginBottom: "12px", color: "#1f2937", textAlign: "center", fontSize: "15px" }}>
                            Statistics
                        </h3>

                        <div style={{ background: "#f9fafb", borderRadius: "10px", padding: "10px", marginBottom: "10px", textAlign: "center" }}>
                            <div style={{ fontSize: "12px", color: "#6b7280", marginBottom: "4px" }}>Average grade</div>
                            <div style={{ fontSize: "26px", fontWeight: 700, color: "#2563eb" }}>
                                {average?.average.toFixed(2) ?? "0.00"}
                            </div>
                        </div>

                        <div style={{ display: "flex", gap: "8px" }}>
                            <div style={{ background: "#dcfce7", borderRadius: "10px", padding: "8px", textAlign: "center", flex: 1 }}>
                                <div style={{ fontSize: "11px", color: "#166534" }}>Excused</div>
                                <div style={{ fontSize: "20px", fontWeight: 700, color: "#166534" }}>{absenceStats?.excused ?? "-"}</div>
                            </div>
                            <div style={{ background: "#fee2e2", borderRadius: "10px", padding: "8px", textAlign: "center", flex: 1 }}>
                                <div style={{ fontSize: "11px", color: "#991b1b" }}>Unexcused</div>
                                <div style={{ fontSize: "20px", fontWeight: 700, color: "#991b1b" }}>{absenceStats?.unexcused ?? "-"}</div>
                            </div>
                            <div style={{ background: "#fef9c3", borderRadius: "10px", padding: "8px", textAlign: "center", flex: 1 }}>
                                <div style={{ fontSize: "11px", color: "#854d0e" }}>Pending</div>
                                <div style={{ fontSize: "20px", fontWeight: 700, color: "#854d0e" }}>{absenceStats?.pending ?? "-"}</div>
                            </div>
                        </div>
                    </div>
                </div>

                <div style={{ display: "flex", gap: "24px", alignItems: "flex-start" }}>

                    {/* GRADES */}
                    <div className="profile-right" style={{ flex: 1 }}>

                        <h3 style={{ color: "#1f2937", marginBottom: "18px" }}>
                            Grades
                        </h3>

                        {data.grades.length === 0 ? (
                            <p style={{ color: "#6b7280" }}>No grades yet.</p>
                        ) : (
                            <div style={{ display: "flex", flexDirection: "column", gap: "7px" }}>
                                {data.grades.map(g => {

                                    const colors = {
                                        1: { bg: "#fee2e2", text: "#991b1b" },
                                        2: { bg: "#ffedd5", text: "#c2410c" },
                                        3: { bg: "#fef9c3", text: "#854d0e" },
                                        4: { bg: "#dcfce7", text: "#166534" },
                                        5: { bg: "#dbeafe", text: "#1d4ed8" },
                                    }[g.value] ?? { bg: "#f3f4f6", text: "#374151" }

                                    return (
                                        <div
                                            key={g.id}
                                            style={{
                                                background: colors.bg,
                                                borderRadius: "12px",
                                                padding: "9px 12px",
                                                display: "flex",
                                                justifyContent: "space-between",
                                                alignItems: "center",
                                            }}
                                        >
                                            <span
                                                style={{
                                                    fontSize: 14,
                                                    color: "#374151",
                                                    fontWeight: 500
                                                }}
                                            >
                                                {new Date(g.date).toLocaleDateString()}
                                            </span>

                                            <span
                                                style={{
                                                    width: 38,
                                                    height: 38,
                                                    borderRadius: "999px",
                                                    background: colors.text,
                                                    color: "white",
                                                    display: "flex",
                                                    justifyContent: "center",
                                                    alignItems: "center",
                                                    fontSize: 18,
                                                    fontWeight: 700
                                                }}
                                            >
                                                {g.value}
                                            </span>
                                        </div>
                                    )
                                })}
                            </div>
                        )}

                    </div>

                    {/* ABSENCES */}
                    <div className="profile-right" style={{ flex: 1 }}>

                        <h3 style={{ color: "#1f2937", marginBottom: "18px" }}>
                            Absences
                        </h3>

                        {data.absences.length === 0 ? (
                            <p style={{ color: "#6b7280" }}>No absences.</p>
                        ) : (
                            <div style={{ display: "flex", flexDirection: "column", gap: "7px" }}>
                                {data.absences.map(a => {

                                    const style =
                                        a.status === "EXCUSED"
                                            ? { bg: "#dcfce7", text: "#166534" }
                                            : a.status === "UNEXCUSED"
                                                ? { bg: "#fee2e2", text: "#991b1b" }
                                                : { bg: "#fef9c3", text: "#854d0e" }

                                    return (
                                        <div
                                            key={a.id}
                                            style={{
                                                background: style.bg,
                                                borderRadius: "12px",
                                                padding: "9px 12px",
                                                display: "flex",
                                                justifyContent: "space-between",
                                                alignItems: "center",
                                            }}
                                        >
                                            <span
                                                style={{
                                                    fontSize: 14,
                                                    color: "#374151",
                                                    fontWeight: 500
                                                }}
                                            >
                                                {new Date(a.date).toLocaleDateString()}
                                            </span>

                                            <select
                                                value={a.status}
                                                onChange={(e) => handleStatusChange(a.id, e.target.value)}
                                                style={{
                                                    background: "transparent",
                                                    border: "none",
                                                    color: style.text,
                                                    fontWeight: 700,
                                                    fontSize: 14,
                                                    cursor: "pointer",
                                                    outline: "none",
                                                }}
                                            >
                                                <option value="PENDING">Pending</option>
                                                <option value="EXCUSED">Excused</option>
                                                <option value="UNEXCUSED">Unexcused</option>
                                            </select>

                                        </div>
                                    )
                                })}
                            </div>
                        )}

                    </div>

                    <div className="profile-right" style={{ flex: 1 }}>
                        <h3 style={{ color: "#1f2937", marginBottom: "18px" }}>Certificates</h3>
                        {certificates.length === 0 ? (
                            <p style={{ color: "#6b7280" }}>No certificates.</p>
                        ) : (
                            <div style={{ display: "flex", flexDirection: "column", gap: "7px" }}>
                                {certificates.map(c => (
                                    <div key={c.ID} style={{
                                        background: "#f0f9ff",
                                        borderRadius: "12px",
                                        padding: "9px 12px",
                                        display: "flex",
                                        flexDirection: "column",
                                        gap: "4px"
                                    }}>
                                        <div style={{ display: "flex", justifyContent: "space-between", alignItems: "center" }}>
                                            <span style={{ fontWeight: 700, color: "#1f2937", fontSize: "14px" }}>{c.Type}</span>
                                            <span style={{ fontSize: "13px", color: "#6b7280" }}>{new Date(c.Date).toLocaleDateString()}</span>
                                        </div>
                                        {c.Note && <span style={{ fontSize: "13px", color: "#6b7280" }}>{c.Note}</span>}
                                    </div>
                                ))}
                            </div>
                        )}
                    </div>

                </div>

            </div>

            {showGradeModal && (
                <AddGradeModal
                    studentUserID={Number(userID)}
                    subjectID={Number(subjectID)}
                    onClose={() => setShowGradeModal(false)}
                    onSuccess={loadData}
                />
            )}

            {showAbsenceModal && (
                <AddAbsenceModal
                    studentUserID={Number(userID)}
                    subjectID={Number(subjectID)}
                    onClose={() => setShowAbsenceModal(false)}
                    onSuccess={loadData}
                />
            )}
        </div>
    )
}