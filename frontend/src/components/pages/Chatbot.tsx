import { useState, useRef, useEffect } from 'react'
import {
  askChatbot,
  getInstructions,
  createInstruction,
  updateInstruction,
  deleteInstruction,
  getSettings,
  updateSettings,
  type BotInstruction,
  type BotSettings,
} from '../../services/chatbotService'

interface ChatMessage {
  role: 'user' | 'bot'
  text: string
}

export default function Chatbot() {
  const user = JSON.parse(localStorage.getItem('user') || '{}')
  const isAdmin = user.role === 'ADMIN'

  const [activeTab, setActiveTab] = useState<'chat' | 'settings'>('chat')

  const [messages, setMessages] = useState<ChatMessage[]>([])
  const [input, setInput] = useState('')
  const [loading, setLoading] = useState(false)
  const bottomRef = useRef<HTMLDivElement>(null)

  const [instructions, setInstructions] = useState<BotInstruction[]>([])
  const [settings, setSettings] = useState<BotSettings | null>(null)
  const [newContent, setNewContent] = useState('')
  const [newRole, setNewRole] = useState('')
  const [newOrder, setNewOrder] = useState(1)

  useEffect(() => {
    bottomRef.current?.scrollIntoView({ behavior: 'smooth' })
  }, [messages])

  useEffect(() => {
    if (activeTab === 'settings' && isAdmin) {
      fetchSettingsData()
    }
  }, [activeTab])

  async function fetchSettingsData() {
    try {
      const [instr, sett] = await Promise.all([getInstructions(), getSettings()])
      setInstructions(instr)
      setSettings(sett)
    } catch (err) {
      console.error(err)
    }
  }

  async function handleSend() {
    const trimmed = input.trim()
    if (!trimmed || loading) return

    setMessages(prev => [...prev, { role: 'user', text: trimmed }])
    setInput('')
    setLoading(true)

    try {
      const answer = await askChatbot(trimmed)
      setMessages(prev => [...prev, { role: 'bot', text: answer }])
    } catch (err) {
      setMessages(prev => [...prev, { role: 'bot', text: 'Doslo je do greske. Pokusaj ponovo.' }])
      console.error(err)
    } finally {
      setLoading(false)
    }
  }

  function handleKeyDown(e: React.KeyboardEvent<HTMLInputElement>) {
    if (e.key === 'Enter') handleSend()
  }

  async function handleAddInstruction() {
    if (!newContent.trim()) return
    try {
      await createInstruction(newContent.trim(), newRole, newOrder)
      setNewContent('')
      setNewRole('')
      setNewOrder(1)
      fetchSettingsData()
    } catch (err) {
      console.error(err)
    }
  }

  async function handleToggleEnabled(instr: BotInstruction) {
    try {
      await updateInstruction(instr.ID, instr.Content, instr.Role, instr.Order, !instr.Enabled)
      fetchSettingsData()
    } catch (err) {
      console.error(err)
    }
  }

  async function handleDeleteInstruction(id: number) {
    try {
      await deleteInstruction(id)
      fetchSettingsData()
    } catch (err) {
      console.error(err)
    }
  }

  async function handleSaveSettings() {
    if (!settings) return
    try {
      await updateSettings(settings.temperature, settings.maxTokens, settings.groqModel)
    } catch (err) {
      console.error(err)
    }
  }

  return (
    <div className="crud-page" style={{ maxWidth: '700px' }}>
      <h1>Asistent</h1>

      {isAdmin && (
        <div className="tabs">
          <button
            className={`tab-btn ${activeTab === 'chat' ? 'active' : ''}`}
            onClick={() => setActiveTab('chat')}
          >
            Chat
          </button>
          <button
            className={`tab-btn ${activeTab === 'settings' ? 'active' : ''}`}
            onClick={() => setActiveTab('settings')}
          >
            Settings
          </button>
        </div>
      )}

      {activeTab === 'chat' && (
        <>
          <div
            style={{
              display: 'flex',
              flexDirection: 'column',
              gap: '12px',
              background: '#f9fafb',
              borderRadius: '16px',
              padding: '20px',
              minHeight: '400px',
              maxHeight: '500px',
              overflowY: 'auto',
            }}
          >
            {messages.length === 0 && (
              <p style={{ color: '#6b7280', fontSize: '14px' }}>
                Postavi pitanje o predmetima, nastavnicima ili doktorima.
              </p>
            )}

            {messages.map((m, i) => (
              <div
                key={i}
                style={{
                  alignSelf: m.role === 'user' ? 'flex-end' : 'flex-start',
                  background: m.role === 'user' ? '#2563eb' : '#e5e7eb',
                  color: m.role === 'user' ? '#fff' : '#1f2937',
                  borderRadius: '14px',
                  padding: '10px 16px',
                  maxWidth: '75%',
                  fontSize: '14px',
                  whiteSpace: 'pre-wrap',
                }}
              >
                {m.text}
              </div>
            ))}

            {loading && (
              <div
                style={{
                  alignSelf: 'flex-start',
                  background: '#e5e7eb',
                  color: '#6b7280',
                  borderRadius: '14px',
                  padding: '10px 16px',
                  fontSize: '14px',
                }}
              >
                Kucam...
              </div>
            )}

            <div ref={bottomRef} />
          </div>

          <div style={{ display: 'flex', gap: '10px', marginTop: '16px' }}>
            <input
              type="text"
              value={input}
              onChange={e => setInput(e.target.value)}
              onKeyDown={handleKeyDown}
              placeholder="Postavi pitanje..."
              className="profile-edit-input"
              style={{ flex: 1 }}
            />
            <button className="btn-primary" onClick={handleSend} disabled={loading}>
              Posalji
            </button>
          </div>
        </>
      )}

      {activeTab === 'settings' && isAdmin && (
        <div style={{ display: 'flex', flexDirection: 'column', gap: '24px' }}>

          <div className="profile-right">
            <h3 style={{ color: '#1f2937', marginBottom: '16px' }}>Generisanje odgovora</h3>
            {settings && (
              <div style={{ display: 'flex', flexDirection: 'column', gap: '14px' }}>
                <div className="form-group">
                  <label>Temperatura ({settings.temperature})</label>
                  <input
                    type="range"
                    min="0"
                    max="1"
                    step="0.1"
                    value={settings.temperature}
                    onChange={e => setSettings({ ...settings, temperature: Number(e.target.value) })}
                  />
                </div>
                <div className="form-group">
                  <label>Max tokens</label>
                  <input
                    type="number"
                    value={settings.maxTokens}
                    onChange={e => setSettings({ ...settings, maxTokens: Number(e.target.value) })}
                  />
                </div>
                <div className="form-group">
                  <label>Groq model</label>
                  <input
                    type="text"
                    value={settings.groqModel}
                    onChange={e => setSettings({ ...settings, groqModel: e.target.value })}
                  />
                </div>
                <button className="btn-primary" style={{ width: 'fit-content' }} onClick={handleSaveSettings}>
                  Sacuvaj podesavanja
                </button>
              </div>
            )}
          </div>

          <div className="profile-right">
            <h3 style={{ color: '#1f2937', marginBottom: '16px' }}>Instrukcije za bota</h3>

            <div className="crud-list" style={{ marginBottom: '20px' }}>
              {instructions.length === 0 ? (
                <p style={{ color: '#6b7280', fontSize: '14px' }}>Nema instrukcija.</p>
              ) : (
                instructions
                  .sort((a, b) => a.Order - b.Order)
                  .map(instr => (
                    <div key={instr.ID} className="crud-item" style={{ alignItems: 'flex-start' }}>
                      <div style={{ flex: 1 }}>
                        <span style={{ display: 'block' }}>{instr.Content}</span>
                        <span style={{ fontSize: '12px', color: '#9ca3af' }}>
                          {instr.Role || 'sve role'} · redosled {instr.Order}
                        </span>
                      </div>
                      <div className="crud-item-actions">
                        <button
                          className="btn-secondary"
                          onClick={() => handleToggleEnabled(instr)}
                        >
                          {instr.Enabled ? 'Iskljuci' : 'Ukljuci'}
                        </button>
                        <button
                          className="btn-danger"
                          onClick={() => handleDeleteInstruction(instr.ID)}
                        >
                          Obrisi
                        </button>
                      </div>
                    </div>
                  ))
              )}
            </div>

            <div className="modal-section-label">Nova instrukcija</div>
            <div style={{ display: 'flex', flexDirection: 'column', gap: '10px' }}>
              <textarea
                className="profile-edit-input"
                style={{ textAlign: 'left', minHeight: '60px' }}
                placeholder="Sadrzaj instrukcije..."
                value={newContent}
                onChange={e => setNewContent(e.target.value)}
              />
              <div style={{ display: 'flex', gap: '10px' }}>
                <select
                  className="role-select"
                  value={newRole}
                  onChange={e => setNewRole(e.target.value)}
                >
                  <option value="">Sve role</option>
                  <option value="STUDENT">Student</option>
                  <option value="TEACHER">Teacher</option>
                  <option value="DOCTOR">Doctor</option>
                  <option value="ADMIN">Admin</option>
                </select>
                <input
                  type="number"
                  className="profile-edit-input"
                  style={{ maxWidth: '100px' }}
                  value={newOrder}
                  onChange={e => setNewOrder(Number(e.target.value))}
                />
                <button className="btn-primary" onClick={handleAddInstruction}>
                  Dodaj
                </button>
              </div>
            </div>
          </div>

        </div>
      )}
    </div>
  )
}