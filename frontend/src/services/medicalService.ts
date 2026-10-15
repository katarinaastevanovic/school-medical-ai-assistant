function authHeaders(): HeadersInit {
  return {
    'Content-Type': 'application/json',
    'Authorization': `Bearer ${localStorage.getItem('token') ?? ''}`,
  }
}

export interface ExaminationRequest {
  id: number
  patientID: number
  patientUserID: number
  doctorID: number
  type: string
  status: string
  medicalRecordID: number
}

export interface MedicalCertificate {
  ID: number
  PatientID: number
  DoctorID: number
  Date: string
  Type: string
  Note: string
}

export async function getExaminationRequests(): Promise<ExaminationRequest[]> {
  const response = await fetch('http://localhost:8080/examination-requests', {
    headers: authHeaders(),
  })
  if (!response.ok) throw new Error(await response.text())
  return response.json()
}

export async function updateRequestStatus(requestID: number, status: string): Promise<void> {
  const response = await fetch(`http://localhost:8080/examination-requests/${requestID}/status`, {
    method: 'PUT',
    headers: authHeaders(),
    body: JSON.stringify({ status }),
  })
  if (!response.ok) throw new Error(await response.text())
}

export async function createExamination(requestID: number, diagnosis: string, therapy: string, note: string): Promise<void> {
  const response = await fetch(`http://localhost:8080/examination-requests/${requestID}/examination`, {
    method: 'POST',
    headers: authHeaders(),
    body: JSON.stringify({ diagnosis, therapy, note }),
  })
  if (!response.ok) throw new Error(await response.text())
}

export async function createCertificate(patientUserID: number, type: string, note: string, date: string): Promise<void> {
  const response = await fetch('http://localhost:8080/certificates', {
    method: 'POST',
    headers: authHeaders(),
    body: JSON.stringify({ patientUserID, type, note, date }),
  })
  if (!response.ok) throw new Error(await response.text())
}

export interface Patient {
  ID: number
  UserID: number
}

export async function getAllPatients(): Promise<Patient[]> {
  const response = await fetch('http://localhost:8080/patients', {
    headers: authHeaders(),
  })
  if (!response.ok) throw new Error(await response.text())
  return response.json()
}

export interface MyExaminationRequest {
  ID: number
  PatientID: number
  DoctorID: number
  Type: string
  Status: string
}

export interface MyCertificate {
  ID: number
  PatientID: number
  DoctorID: number
  Date: string
  Type: string
  Note: string
}

export async function getMyExaminationRequests(): Promise<MyExaminationRequest[]> {
  const response = await fetch('http://localhost:8080/examination-requests/me', {
    headers: authHeaders(),
  })
  if (!response.ok) throw new Error(await response.text())
  return response.json()
}

export async function createExaminationRequest(doctorID: number, type: string): Promise<void> {
  const response = await fetch('http://localhost:8080/examination-requests', {
    method: 'POST',
    headers: authHeaders(),
    body: JSON.stringify({ doctorID, type }),
  })
  if (!response.ok) throw new Error(await response.text())
}

export async function getMyCertificates(): Promise<MyCertificate[]> {
  const response = await fetch('http://localhost:8080/certificates/me', {
    headers: authHeaders(),
  })
  if (!response.ok) throw new Error(await response.text())
  return response.json()
}

export async function getCertificatesByUserID(userID: number): Promise<MedicalCertificate[]> {
  const response = await fetch(`http://localhost:8080/certificates/${userID}`, {
    headers: authHeaders(),
  })
  if (!response.ok) throw new Error(await response.text())
  return response.json()
}

export interface Examination {
  ID: number
  RequestID: number
  MedicalRecordID: number
  Diagnosis: string
  Therapy: string
  Note: string
  CreatedAt: string
}

export interface MedicalRecord {
  ID: number
  PatientID: number
  Examinations: Examination[]
}

export async function getMedicalRecordByUserID(userID: number): Promise<MedicalRecord> {
  const response = await fetch(`http://localhost:8080/medical-record/${userID}`, {
    headers: authHeaders(),
  })
  if (!response.ok) throw new Error(await response.text())
  return response.json()
}

export async function getMyMedicalRecord(): Promise<MedicalRecord> {
  const response = await fetch('http://localhost:8080/medical-record/me', {
    headers: authHeaders(),
  })
  if (!response.ok) throw new Error(await response.text())
  return response.json()
}