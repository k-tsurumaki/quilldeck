import React, { useState } from 'react'

function App() {
  const [status, setStatus] = useState<string>('')

  const checkHealth = async () => {
    try {
      const response = await fetch('/api/health')
      const data = await response.json()
      setStatus(`Backend Status: ${data.status}`)
    } catch (error) {
      setStatus('Backend connection failed')
    }
  }

  return (
    <div style={{ padding: '20px' }}>
      <h1>QuillDeck</h1>
      <button onClick={checkHealth}>Check Backend Status</button>
      <p>{status}</p>
    </div>
  )
}

export default App