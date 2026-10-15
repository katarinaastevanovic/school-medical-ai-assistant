function authHeaders(): HeadersInit {
  return {
    'Content-Type': 'application/json',
    'Authorization': `Bearer ${localStorage.getItem('token') ?? ''}`,
  }
}

export async function askChatbot(message: string): Promise<string> {
  const response = await fetch('http://localhost:8080/chatbot/ask', {
    method: 'POST',
    headers: authHeaders(),
    body: JSON.stringify({ message }),
  })
  if (!response.ok) throw new Error(await response.text())
  const data = await response.json()
  return data.answer
}

export interface BotInstruction {
  ID: number
  Content: string
  Role: string
  Order: number
  Enabled: boolean
}

export interface BotSettings {
  ID: number
  temperature: number
  maxTokens: number
  groqModel: string
}

export async function getInstructions(): Promise<BotInstruction[]> {
  const response = await fetch('http://localhost:8080/chatbot/instructions', {
    headers: authHeaders(),
  })
  if (!response.ok) throw new Error(await response.text())
  return response.json()
}

export async function createInstruction(content: string, role: string, order: number): Promise<void> {
  const response = await fetch('http://localhost:8080/chatbot/instructions', {
    method: 'POST',
    headers: authHeaders(),
    body: JSON.stringify({ content, role, order, enabled: true }),
  })
  if (!response.ok) throw new Error(await response.text())
}

export async function updateInstruction(id: number, content: string, role: string, order: number, enabled: boolean): Promise<void> {
  const response = await fetch(`http://localhost:8080/chatbot/instructions/${id}`, {
    method: 'PUT',
    headers: authHeaders(),
    body: JSON.stringify({ content, role, order, enabled }),
  })
  if (!response.ok) throw new Error(await response.text())
}

export async function deleteInstruction(id: number): Promise<void> {
  const response = await fetch(`http://localhost:8080/chatbot/instructions/${id}`, {
    method: 'DELETE',
    headers: authHeaders(),
  })
  if (!response.ok) throw new Error(await response.text())
}

export async function getSettings(): Promise<BotSettings> {
  const response = await fetch('http://localhost:8080/chatbot/settings', {
    headers: authHeaders(),
  })
  if (!response.ok) throw new Error(await response.text())
  return response.json()
}

export async function updateSettings(temperature: number, maxTokens: number, groqModel: string): Promise<void> {
  const response = await fetch('http://localhost:8080/chatbot/settings', {
    method: 'PUT',
    headers: authHeaders(),
    body: JSON.stringify({ temperature, maxTokens, groqModel }),
  })
  if (!response.ok) throw new Error(await response.text())
}