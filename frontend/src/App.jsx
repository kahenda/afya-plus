import { useState, useEffect } from 'react'
import Login from './pages/Login'
import Dashboard from './pages/Dashboard'
import { getSession } from './lib/session'

function App() {
  const [loggedIn, setLoggedIn] = useState(false)

  useEffect(() => {
    setLoggedIn(!!getSession())
  }, [])

  if (loggedIn) {
    return <Dashboard onLogout={() => setLoggedIn(false)} />
  }

  return <Login onSuccess={() => setLoggedIn(true)} />
}

export default App
