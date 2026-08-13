import { useState, useEffect } from 'react'
import { getSession, clearSession } from '../lib/session'
import { getFlags } from '../lib/api'
import LogVisit from './LogVisit'

export default function Dashboard({ onLogout }) {
  const user = getSession()
  const [showLogVisit, setShowLogVisit] = useState(false)
  const [flags, setFlags] = useState([])
  const [loadingFlags, setLoadingFlags] = useState(true)

  useEffect(() => {
    loadFlags()
  }, [])

  function loadFlags() {
    setLoadingFlags(true)
    getFlags()
      .then(setFlags)
      .catch(() => {})
      .finally(() => setLoadingFlags(false))
  }

  function handleLogout() {
    clearSession()
    onLogout()
  }

  if (showLogVisit) {
    return (
      <LogVisit
        onDone={() => {
          setShowLogVisit(false)
          loadFlags()
        }}
        onCancel={() => setShowLogVisit(false)}
      />
    )
  }

  const reasonLabels = {
    malnutrition_risk: 'Possible malnutrition',
    overdue_vaccination: 'Overdue vaccination',
  }

  const statusColors = {
    flagged: 'bg-red-50 text-red-700 border-red-200',
    under_review: 'bg-amber-50 text-amber-700 border-amber-200',
    action_taken: 'bg-blue-50 text-blue-700 border-blue-200',
    resolved: 'bg-green-50 text-green-700 border-green-200',
  }

  return (
    <div className="min-h-screen bg-cream px-6 py-8">
      <div className="max-w-sm mx-auto">
        <div className="flex items-center justify-between mb-8">
          <div className="flex items-center gap-3">
            <div className="w-10 h-10 rounded-full bg-teal flex items-center justify-center">
              <span className="text-cream font-heading font-extrabold text-sm">A+</span>
            </div>
            <h1 className="font-heading font-extrabold text-xl text-ink">Afya Plus</h1>
          </div>
          <button
            onClick={handleLogout}
            className="text-sm text-ink/50 font-medium"
          >
            Log out
          </button>
        </div>

        <div className="bg-white rounded-2xl border border-ink/10 p-6 mb-4">
          <p className="text-ink/50 text-sm mb-1">Logged in as</p>
          <p className="font-heading font-bold text-lg text-ink">{user?.name}</p>
          <p className="text-ink/60 text-sm capitalize">{user?.role} · {user?.zone}</p>
        </div>

        <button
          onClick={() => setShowLogVisit(true)}
          className="w-full h-14 rounded-xl bg-teal hover:bg-teal-dark text-white font-heading font-bold mb-4 transition-colors"
        >
          + Log a visit
        </button>

        <h2 className="font-heading font-bold text-ink mb-3">Flagged cases</h2>

        {loadingFlags ? (
          <div className="bg-white rounded-2xl border border-ink/10 p-6 text-center">
            <p className="text-ink/50 text-sm">Loading...</p>
          </div>
        ) : flags.length === 0 ? (
          <div className="bg-white rounded-2xl border border-ink/10 p-6 text-center">
            <p className="text-ink/50 text-sm">No flagged cases right now.</p>
          </div>
        ) : (
          <div className="space-y-3">
            {flags.map((f) => (
              <div key={f.id} className={`rounded-2xl border p-4 ${statusColors[f.status] || 'bg-white border-ink/10'}`}>
                <div className="flex items-center justify-between mb-1">
                  <p className="font-heading font-bold">{reasonLabels[f.reason] || f.reason}</p>
                  <span className="text-xs font-medium capitalize px-2 py-1 rounded-full bg-white/60">
                    {f.status.replace('_', ' ')}
                  </span>
                </div>
                <p className="text-xs opacity-70">Household #{f.household_id} · Visit #{f.visit_id}</p>
              </div>
            ))}
          </div>
        )}
      </div>
    </div>
  )
}
