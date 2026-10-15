import { useEffect, useState } from "react"
import { useLocation, useParams } from "react-router-dom"
import { getMyAbsenceStats, getMyAverageGrade, getMyGradesAndAbsences, type AbsenceStatsResponse, type AverageGradeResponse, type GradesAndAbsences, type StudentSubject } from "../../services/schoolService"

export default function StudentSubjectDetails() {

    const { subjectID } = useParams()
    const { state } = useLocation()
    const subject = state as StudentSubject
    const [data, setData] = useState<GradesAndAbsences | null>(null)
    const [average, setAverage] = useState<AverageGradeResponse | null>(null)
    const [absenceStats, setAbsenceStats] = useState<AbsenceStatsResponse | null>(null)

    const getGradeStyle = (value: number) => {
        switch (value) {
            case 1: return { backgroundColor: "#fee2e2", color: "#b91c1c" }
            case 2: return { backgroundColor: "#ffedd5", color: "#c2410c" }
            case 3: return { backgroundColor: "#fef9c3", color: "#854d0e" }
            case 4: return { backgroundColor: "#dcfce7", color: "#166534" }
            case 5: return { backgroundColor: "#dbeafe", color: "#1d4ed8" }
            default: return { backgroundColor: "#f3f4f6", color: "#374151" }
        }
    }

    useEffect(() => {
        if (!subjectID) return
        getMyGradesAndAbsences(Number(subjectID)).then(setData).catch(console.error)
        getMyAverageGrade(Number(subjectID)).then(setAverage).catch(console.error)
        getMyAbsenceStats(Number(subjectID)).then(setAbsenceStats).catch(console.error)
    }, [subjectID])

    if (!data) return <p>Loading...</p>

    return (
        <div className="profile-page" style={{ fontSize: "14px" }}>
            <div className="profile-container" style={{ gap: "16px" }}>

                <div className="profile-left" style={{ padding: "12px" }}>

                    <img
                        src="https://bootdey.com/img/Content/avatar/avatar7.png"
                        alt="avatar"
                        className="profile-avatar"
                        style={{ width: "64px", height: "64px" }}
                    />

                    <h3 className="profile-role" style={{ margin: "4px 0 0", fontSize: "12px" }}>
                        Teacher:
                    </h3>

                    <h2 style={{ margin: "2px 0", fontSize: "16px" }}>{subject.firstName} {subject.lastName}</h2>

                    <h3 style={{ marginTop: "8px", fontSize: "14px" }}>
                        {subject.subjectName}
                    </h3>

                    <div
                        style={{
                            marginTop: "12px",
                            padding: "8px",
                            borderRadius: "10px",
                            background: "#f3f4f6",
                            textAlign: "center"
                        }}
                    >
                        <div style={{ fontSize: "11px", color: "#6b7280" }}>
                            Average grade
                        </div>

                        <div style={{ fontSize: "22px", fontWeight: 700, color: "#2563eb" }}>
                            {average ? average.average.toFixed(2) : "-"}
                        </div>

                        <div style={{ marginTop: "10px", display: "flex", flexDirection: "column", gap: "6px" }}>

                            <div style={{ background: "#dcfce7", borderRadius: "10px", padding: "6px 8px", textAlign: "center" }}>
                                <div style={{ fontSize: 11, color: "#166534" }}>Excused</div>
                                <div style={{ fontSize: 18, fontWeight: 700, color: "#166534" }}>
                                    {absenceStats?.excused ?? "-"}
                                </div>
                            </div>

                            <div style={{ background: "#fee2e2", borderRadius: "10px", padding: "6px 8px", textAlign: "center" }}>
                                <div style={{ fontSize: 11, color: "#991b1b" }}>Unexcused</div>
                                <div style={{ fontSize: 18, fontWeight: 700, color: "#991b1b" }}>
                                    {absenceStats?.unexcused ?? "-"}
                                </div>
                            </div>

                            <div style={{ background: "#fef9c3", borderRadius: "10px", padding: "6px 8px", textAlign: "center" }}>
                                <div style={{ fontSize: 11, color: "#854d0e" }}>Pending</div>
                                <div style={{ fontSize: 18, fontWeight: 700, color: "#854d0e" }}>
                                    {absenceStats?.pending ?? "-"}
                                </div>
                            </div>

                        </div>
                    </div>

                </div>

                <div style={{ display: "flex", gap: "16px", alignItems: "flex-start" }}>

                    <div className="profile-right" style={{ flex: 1, maxHeight: "70vh", overflowY: "auto" }}>
                        <h3 style={{ color: '#1f2937', marginBottom: '8px', fontSize: "14px" }}>Grades</h3>
                        {data.grades.length === 0 ? (
                            <p style={{ color: '#6b7280', fontSize: '13px' }}>No grades yet.</p>
                        ) : (
                            <div style={{ display: 'flex', flexDirection: 'column', gap: '6px' }}>
                                {data.grades.map(g => (
                                    <div
                                        key={g.id}
                                        style={{
                                            ...getGradeStyle(g.value),
                                            borderRadius: "10px",
                                            padding: "6px 12px",
                                            display: "flex",
                                            justifyContent: "space-between",
                                            alignItems: "center",
                                            fontWeight: 600
                                        }}
                                    >
                                        <span style={{ fontSize: "12px", color: "#374151" }}>
                                            {new Date(g.date).toLocaleDateString()}
                                        </span>
                                        <span style={{ fontSize: "16px", fontWeight: 700, color: getGradeStyle(g.value).color }}>
                                            {g.value}
                                        </span>
                                    </div>
                                ))}
                            </div>
                        )}
                    </div>

                    <div className="profile-right" style={{ flex: 1, maxHeight: "70vh", overflowY: "auto" }}>
                        <h3 style={{ color: '#1f2937', marginBottom: '8px', fontSize: "14px" }}>Absences</h3>
                        {data.absences.length === 0 ? (
                            <p style={{ color: '#6b7280', fontSize: '13px' }}>No absences.</p>
                        ) : (
                            <div style={{ display: 'flex', flexDirection: 'column', gap: '6px' }}>
                                {data.absences.map(a => {
                                    const absenceStyle = {
                                        PENDING: { background: '#fef9c3', color: '#854d0e' },
                                        EXCUSED: { background: '#dcfce7', color: '#166534' },
                                        UNEXCUSED: { background: '#fee2e2', color: '#991b1b' },
                                    }[a.status] ?? { background: '#f3f4f6', color: '#1f2937' }

                                    return (
                                        <div
                                            key={a.id}
                                            style={{
                                                ...absenceStyle,
                                                borderRadius: '10px',
                                                padding: '6px 12px',
                                                display: 'flex',
                                                justifyContent: 'space-between',
                                                alignItems: 'center',
                                            }}
                                        >
                                            <span style={{ fontSize: '12px', fontWeight: 600 }}>
                                                {new Date(a.date).toLocaleDateString()}
                                            </span>
                                            <span style={{ fontSize: '11px', fontWeight: 700, color: absenceStyle.color }}>
                                                {a.status}
                                            </span>
                                        </div>
                                    )
                                })}
                            </div>
                        )}
                    </div>

                </div>

            </div>
        </div>
    )
}