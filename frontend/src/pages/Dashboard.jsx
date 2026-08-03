import { getSession, clearSession } from '../lib/session'

export default function Dashboard({ onLogout }) {
  const user = getSession()

  function handleLogout() {
    clearSession()
    onLogout()
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

        <div className="bg-white rounded-2xl border border-ink/10 p-6 text-center">
          <p className="text-ink/50 text-sm">Household visits and flags will show up here.</p>
        </div>
      </div>
    </div>
  )
}
