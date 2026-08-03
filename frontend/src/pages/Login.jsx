import { useState, useRef } from 'react'
import { saveSession } from '../lib/session'

export default function Login({ onSuccess }) {
  const [phone, setPhone] = useState('')
  const [pin, setPin] = useState(['', '', '', ''])
  const [error, setError] = useState('')
  const [loading, setLoading] = useState(false)
  const pinRefs = [useRef(), useRef(), useRef(), useRef()]

  function handlePinChange(index, value) {
    if (!/^\d?$/.test(value)) return
    const newPin = [...pin]
    newPin[index] = value
    setPin(newPin)
    if (value && index < 3) {
      pinRefs[index + 1].current?.focus()
    }
  }

  function handlePinKeyDown(index, e) {
    if (e.key === 'Backspace' && !pin[index] && index > 0) {
      pinRefs[index - 1].current?.focus()
    }
  }

  async function handleSubmit(e) {
    e.preventDefault()
    setError('')

    if (phone.length < 10 || pin.some((d) => d === '')) {
      setError('Enter your full phone number and 4-digit PIN.')
      return
    }

    setLoading(true)

    let res
    try {
      res = await fetch('http://localhost:8080/login', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ phone_number: phone, pin: pin.join('') }),
      })
    } catch (networkErr) {
      console.error('NETWORK ERROR:', networkErr)
      setError('Network error: ' + networkErr.message)
      setLoading(false)
      return
    }

    if (!res.ok) {
      setError('Phone number or PIN is incorrect.')
      setLoading(false)
      return
    }

    try {
      const user = await res.json()
      console.log('Parsed user:', user)
      saveSession(user)
      onSuccess()
    } catch (parseErr) {
      console.error('PARSE/SESSION ERROR:', parseErr)
      setError('Error after login: ' + parseErr.message)
    }

    setLoading(false)
  }

  return (
    <div className="min-h-screen bg-cream flex flex-col items-center justify-center px-6">
      <div className="w-full max-w-sm">
        <div className="flex items-center gap-3 mb-10 justify-center">
          <div className="w-11 h-11 rounded-full bg-teal flex items-center justify-center">
            <span className="text-cream font-heading font-extrabold text-lg">A+</span>
          </div>
          <h1 className="font-heading font-extrabold text-2xl text-ink">Afya Plus</h1>
        </div>

        <form onSubmit={handleSubmit} className="bg-white rounded-2xl shadow-sm border border-ink/10 p-6">
          <h2 className="font-heading font-bold text-xl text-ink mb-1">Welcome back</h2>
          <p className="text-ink/60 text-sm mb-6">Log in to log household visits.</p>

          <label className="block text-sm font-medium text-ink mb-2">Phone number</label>
          <input
            type="tel"
            inputMode="numeric"
            value={phone}
            onChange={(e) => setPhone(e.target.value)}
            placeholder="07XX XXX XXX"
            className="w-full h-14 px-4 rounded-xl border border-ink/15 text-lg text-ink placeholder:text-ink/30 focus:outline-none focus:ring-2 focus:ring-teal mb-6"
          />

          <label className="block text-sm font-medium text-ink mb-3">4-digit PIN</label>
          <div className="flex gap-3 mb-2 justify-between">
            {pin.map((digit, i) => (
              <input
                key={i}
                ref={pinRefs[i]}
                type="password"
                inputMode="numeric"
                maxLength={1}
                value={digit}
                onChange={(e) => handlePinChange(i, e.target.value)}
                onKeyDown={(e) => handlePinKeyDown(i, e)}
                className="w-14 h-16 text-center text-2xl rounded-xl border-2 border-ink/15 text-ink focus:outline-none focus:border-teal focus:ring-2 focus:ring-teal/30 transition-colors"
              />
            ))}
          </div>

          {error && (
            <p className="text-red-600 text-sm mt-3">{error}</p>
          )}

          <button
            type="submit"
            disabled={loading}
            className="w-full h-14 mt-6 rounded-xl bg-teal hover:bg-teal-dark text-white font-heading font-bold text-base transition-colors disabled:opacity-60"
          >
            {loading ? 'Logging in...' : 'Log in'}
          </button>
        </form>

        <p className="text-center text-ink/50 text-xs mt-6">
          Don't have an account? Ask your supervisor to register you.
        </p>
      </div>
    </div>
  )
}
