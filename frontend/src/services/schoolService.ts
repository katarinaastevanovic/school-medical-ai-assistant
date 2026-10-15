function authHeaders(): HeadersInit {
  return {
    'Content-Type': 'application/json',
    'Authorization': `Bearer ${localStorage.getItem('token') ?? ''}`,
  }
}

export interface SchoolClass {
  ID: number
  Name: string
  Year: number
}

export async function createClass(name: string, year: number): Promise<SchoolClass> {
  const response = await fetch('http://localhost:8080/class', {
    method: 'POST',
    headers: authHeaders(),
    body: JSON.stringify({ name, year }),
  })
  if (!response.ok) throw new Error(await response.text())
  return response.json()
}

export async function getAllClasses(): Promise<SchoolClass[]> {
  const response = await fetch('http://localhost:8080/classes', {
    headers: authHeaders(),
  })
  if (!response.ok) throw new Error(await response.text())
  return response.json()
}

export interface Student {
  ID: number
  UserID: number
  ClassID: number
}

export async function getStudentsFromClass(classID: number): Promise<Student[]> {
  const response = await fetch(`http://localhost:8080/classes/${classID}/students`, {
    headers: authHeaders(),
  })
  if (!response.ok) throw new Error(await response.text())
  return response.json()
}

export async function removeStudentFromClass(classID: number, studentID: number): Promise<void> {
  const response = await fetch(`http://localhost:8080/classes/${classID}/students/${studentID}`, {
    method: 'DELETE',
    headers: authHeaders(),
  })
  if (!response.ok) throw new Error(await response.text())
}

export interface FreeStudent {
  id: number
  firstName: string
  lastName: string
}

export async function getFreeStudents(): Promise<FreeStudent[]> {
  const response = await fetch('http://localhost:8080/students/free', {
    headers: authHeaders(),
  })
  if (!response.ok) throw new Error(await response.text())
  return response.json()
}

export interface Subject {
  ID: number
  Name: string
}

export async function getAllSubjects(): Promise<Subject[]> {
  const response = await fetch('http://localhost:8080/subjects', {
    headers: authHeaders(),
  })
  if (!response.ok) throw new Error(await response.text())
  return response.json()
}

export async function createSubject(name: string): Promise<Subject> {
  const response = await fetch('http://localhost:8080/subjects', {
    method: 'POST',
    headers: authHeaders(),
    body: JSON.stringify({ Name: name }),
  })
  if (!response.ok) throw new Error(await response.text())
  return response.json()
}

export async function updateSubject(id: number, name: string): Promise<void> {
  const response = await fetch(`http://localhost:8080/subjects/${id}`, {
    method: 'PUT',
    headers: authHeaders(),
    body: JSON.stringify({ name }),
  })
  if (!response.ok) throw new Error(await response.text())
}

export async function deleteSubject(id: number): Promise<void> {
  const response = await fetch(`http://localhost:8080/subjects/${id}`, {
    method: 'DELETE',
    headers: authHeaders(),
  })
  if (!response.ok) throw new Error(await response.text())
}

export interface TeacherEnriched {
  id?: number
  userID: number
  firstName: string
  lastName: string
}

export async function getTeachersBySubject(subjectID: number): Promise<TeacherEnriched[]> {
  const response = await fetch(`http://localhost:8080/subjects/${subjectID}/teachers`, {
    headers: authHeaders(),
  })
  if (!response.ok) throw new Error(await response.text())
  return response.json()
}

export async function getAllTeachers(): Promise<TeacherEnriched[]> {
  const response = await fetch('http://localhost:8080/teachers', {
    headers: authHeaders(),
  })
  if (!response.ok) throw new Error(await response.text())
  return response.json()
}

export async function addTeacherToSubject(subjectID: number, userID: number): Promise<void> {
  const response = await fetch(`http://localhost:8080/subjects/${subjectID}/teachers`, {
    method: 'POST',
    headers: authHeaders(),
    body: JSON.stringify({ userID }),
  })
  if (!response.ok) throw new Error(await response.text())
}

export interface ClassSubject {
  id: number
  subjectID: number
  subjectName: string
  teacherID: number
  userID: number
  firstName: string
  lastName: string
}

export async function getSubjectsByClass(classID: number): Promise<ClassSubject[]> {
  const response = await fetch(`http://localhost:8080/classes/${classID}/subjects`, {
    headers: authHeaders(),
  })
  if (!response.ok) throw new Error(await response.text())
  return response.json()
}

export async function assignTeacher(userID: number, classID: number, subjectID: number): Promise<void> {
  const response = await fetch('http://localhost:8080/teachers/assign', {
    method: 'POST',
    headers: authHeaders(),
    body: JSON.stringify({ userID, classID, subjectID }),
  })
  if (!response.ok) throw new Error(await response.text())
}

export interface TeacherSubjectClass {
  id: number
  subjectID: number
  subjectName: string
  classID: number
  className: string
}

export async function getMySubjects(): Promise<TeacherSubjectClass[]> {
  const response = await fetch('http://localhost:8080/teachers/me/subjects', {
    headers: authHeaders(),
  })
  if (!response.ok) throw new Error(await response.text())
  return response.json()
}

export interface Grade {
  id: number
  value: number
  date: string
  studentId: number
  teacherId: number
  subjectId: number
}

export interface Absence {
  id: number
  status: string
  date: string
  studentId: number
  subjectId: number
}

export interface GradesAndAbsences {
  grades: Grade[]
  absences: Absence[]
}

export async function getGradesAndAbsences(
  studentUserID: number,
  subjectID: number
): Promise<GradesAndAbsences> {

  const response = await fetch(
    `http://localhost:8080/students/${studentUserID}/subjects/${subjectID}/grades-absences`,
    {
      headers: authHeaders(),
    }
  )

  if (!response.ok) throw new Error(await response.text())

  return response.json()
}

export async function createGrade(
  studentUserID: number,
  subjectID: number,
  value: number,
  date: string
): Promise<void> {
  const response = await fetch(
    `http://localhost:8080/students/${studentUserID}/subjects/${subjectID}/grades`,
    {
      method: "POST",
      headers: authHeaders(),
      body: JSON.stringify({
        value,
        date,
      }),
    }
  )

  if (!response.ok) throw new Error(await response.text())
}

export async function createAbsence(
  studentUserID: number,
  subjectID: number,
  date: string
): Promise<void> {
  const response = await fetch(
    `http://localhost:8080/students/${studentUserID}/subjects/${subjectID}/absences`,
    {
      method: "POST",
      headers: authHeaders(),
      body: JSON.stringify({
        date,
      }),
    }
  )

  if (!response.ok) throw new Error(await response.text())
}

export interface StudentSubject {
  id: number
  subjectID: number
  subjectName: string
  teacherID: number
  userID: number
  firstName: string
  lastName: string
}

export interface MyClassResponse {
  className: number
  classID: number
  subjects: StudentSubject[]
}

export async function getMyClass(): Promise<MyClassResponse> {
  const response = await fetch(
    'http://localhost:8080/students/me/class',
    {
      headers: authHeaders(),
    }
  )

  if (!response.ok) throw new Error(await response.text())

  return response.json()
}

export async function getMyGradesAndAbsences(subjectID: number): Promise<GradesAndAbsences> {

  const response = await fetch(
    `http://localhost:8080/students/me/subjects/${subjectID}/grades-absences`,
    {
      headers: authHeaders(),
    }
  )

  if (!response.ok) throw new Error(await response.text())

  return response.json()
}

export interface AverageGradeResponse {
  average: number
}

export async function getMyAverageGrade(subjectID: number): Promise<AverageGradeResponse> {

  const response = await fetch(
    `http://localhost:8080/students/me/subjects/${subjectID}/average-grade`,
    {
      headers: authHeaders(),
    }
  )

  if (!response.ok) throw new Error(await response.text())

  return response.json()
}

export interface AbsenceStatsResponse {
  pending: number
  excused: number
  unexcused: number
}

export async function getMyAbsenceStats(subjectID: number): Promise<AbsenceStatsResponse> {

  const response = await fetch(
    `http://localhost:8080/students/me/subjects/${subjectID}/absence-stats`,
    {
      headers: authHeaders(),
    }
  )

  if (!response.ok) throw new Error(await response.text())

  return response.json()
}

export async function getAverageGrade(
  studentUserID: number,
  subjectID: number
): Promise<AverageGradeResponse> {

  const response = await fetch(
    `http://localhost:8080/students/${studentUserID}/subjects/${subjectID}/average-grade`,
    {
      headers: authHeaders(),
    }
  )

  if (!response.ok) throw new Error(await response.text())

  return response.json()
}

export async function getAbsenceStats(
  studentUserID: number,
  subjectID: number
): Promise<AbsenceStatsResponse> {

  const response = await fetch(
    `http://localhost:8080/students/${studentUserID}/subjects/${subjectID}/absence-stats`,
    {
      headers: authHeaders(),
    }
  )

  if (!response.ok) throw new Error(await response.text())

  return response.json()
}

export async function updateAbsenceStatus(
  absenceID: number,
  status: string
): Promise<void> {

  const response = await fetch(
    `http://localhost:8080/absences/${absenceID}/status`,
    {
      method: "PUT",
      headers: authHeaders(),
      body: JSON.stringify({
        status,
      }),
    }
  )

  if (!response.ok) throw new Error(await response.text())
}

export interface StudentOverview {
  subjectID: number
  subjectName: string
  average: number
}

export async function getStudentOverview(): Promise<StudentOverview[]> {
  const response = await fetch('http://localhost:8080/students/me/overview', {
    headers: authHeaders(),
  })
  if (!response.ok) throw new Error(await response.text())
  return response.json()
}

export interface AbsenceOverview {
  subjectID: number
  subjectName: string
  pending: number
  excused: number
  unexcused: number
}

export async function getStudentAbsenceOverview(): Promise<AbsenceOverview[]> {
  const response = await fetch('http://localhost:8080/students/me/absence-overview', {
    headers: authHeaders(),
  })
  if (!response.ok) throw new Error(await response.text())
  return response.json()
}

export interface StudentSummaryResponse {
    averageGrade: number
    pending: number
    excused: number
    unexcused: number
}

export async function getStudentSummary(): Promise<StudentSummaryResponse> {
   const response = await fetch('http://localhost:8080/students/me/summary', {
    headers: authHeaders(),
  })
  if (!response.ok) throw new Error(await response.text())
  return response.json()
}

export async function addStudentToClass(classID: number, userID: number): Promise<void> {
  const response = await fetch(`http://localhost:8080/classes/${classID}/students/add`, {
    method: 'POST',
    headers: authHeaders(),
    body: JSON.stringify({ UserID: userID, ClassID: classID }),
  })
  if (!response.ok) throw new Error(await response.text())
}

