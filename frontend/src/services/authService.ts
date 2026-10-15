function authHeaders(): HeadersInit {
  return {
    'Content-Type': 'application/json',
    'Authorization': `Bearer ${localStorage.getItem('token') ?? ''}`,
  }
}

export async function login(email: string, password: string) {
  const response = await fetch('http://localhost:8080/login', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ email, password }),
  })
  if (!response.ok) throw new Error(await response.text())
  return response.json()
}

export async function register(
  umcn: string,
  firstName: string,
  lastName: string,
  email: string,
  password: string,
  role: string,
) {
  const response = await fetch('http://localhost:8080/register', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ umcn, firstName, lastName, email, password, role }),
  })
  if (!response.ok) throw new Error(await response.text())
  return response.json()
}

export function logout() {
  localStorage.removeItem('token')
  localStorage.removeItem('user')
}

export async function getMemberById(id: number) {
  const response = await fetch(`http://localhost:8080/members/${id}`, {
    headers: authHeaders(),
  })
  if (!response.ok) throw new Error('Failed to load profile')
  return response.json()
}

export async function updateMember(
  id: number,
  data: {
    firstName?: string
    lastName?: string
    email?: string
    password?: string
  }
) {
  const response = await fetch(`http://localhost:8080/members/${id}`, {
    method: 'PUT',
    headers: authHeaders(),
    body: JSON.stringify(data),
  })

  if (!response.ok) {
    throw new Error(await response.text())
  }

  return response.json()
}

export function getAvatarByUmcn(umcn: string): string {
  const digits = umcn.replace(/\D/g, '') // za svaki slučaj ukloni ne-cifre

  if (digits.length < 12) {
    return 'https://bootdey.com/img/Content/avatar/avatar7.png' // fallback
  }

  const genderDigits = Number(digits.substring(9, 12)) // cifre 10-12

  return genderDigits >= 500
    ? 'https://bootdey.com/img/Content/avatar/avatar3.png' // žensko
    : 'https://bootdey.com/img/Content/avatar/avatar7.png' // muško
}