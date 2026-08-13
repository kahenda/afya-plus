import { useState, useEffect } from 'react'
import { getHouseholds, createHousehold, createVisit } from '../lib/api'
import { generateVisitId } from '../lib/idgen'
import { getSession } from '../lib/session'

export default function LogVisit({ onDone, onCancel }) {
  const user = getSession()
  const [households, setHouseholds] = useState([])
  const [selectedHousehold, setSelectedHousehold] = useState('')
  const [showNewHousehold, setShowNewHousehold] = useState(false)
  const [newHeadName, setNewHeadName] = useState('')
  const [newLocation, setNewLocation] = useState('')

  const [weight, setWeight] = useState('')
  const [ageMonths, setAgeMonths] = useState('')
  const [vaccinationDone, setVaccinationDone] = useState(false)
  const [notes, setNotes] = useState('')

  const [loading, setLoading] = useState(false)
  const [error, setError] = useState('')
  const [success, setSuccess] = useState('')

  useEffect(() => {
    getHouseholds().then(setHouseholds).catch(() => {})
  }, [])

  async function handleAddHousehold() {
    if (!newHeadName.trim()) {
      setError('Enter the household head\'s name.')
      return
    }
    setError('')
    try {
      const household = await createHousehold({
        head_name: newHeadName,
        location: newLocation,
        member_count: 1,
        created_by: user.id,
      })
      setHouseholds([household, ...households])
      setSelectedHousehold(String(household.id))
      setShowNewHousehold(false)
      setNewHeadName('')
      setNewLocation('')
    } catch (err) {
      setError('Could not add household: ' + err.message)
    }
  }

  async function handleSubmitVisit(e) {
    e.preventDefault()
    setError('')

    if (!selectedHousehold) {
      setError('Select or add a household first.')
      return
    }

    setLoading(true)
    try {
      await createVisit({
        household_id: parseInt(selectedHousehold),
        chw_id: user.id,
        client_visit_id: generateVisitId(),
        visit_type: 'child_checkup',
        child_weight_kg: weight ? parseFloat(weight) : null,
        child_age_months: ageMonths ? parseInt(ageMonths) : null,
        vaccination_done: vaccinationDone,
        notes,
        visited_at: new Date().toISOString(),
      })
      setSuccess('Visit logged successfully.')
      setTimeout(() => onDone(), 1200)
    } catch (err) {
      setError('Could not log visit: ' + err.message)
    }
    setLoading(false)
  }

  return (
    <div className="min-h-screen bg-cream px-6 py-8">
      <div className="max-w-sm mx-auto">
        <div className="flex items-center gap-3 mb-6">
          <button onClick={onCancel} className="text-ink/50 text-xl">←</button>
          <h1 className="font-heading font-extrabold text-xl text-ink">Log a Visit</h1>
        </div>

        <div className="bg-white rounded-2xl border border-ink/10 p-6 mb-4">
          <label className="block text-sm font-medium text-ink mb-2">Household</label>

          {!showNewHousehold ? (
            <>
              <select
                value={selectedHousehold}
                onChange={(e) => setSelectedHousehold(e.target.value)}
                className="w-full h-12 px-3 rounded-xl border border-ink/15 text-ink mb-3"
              >
                <option value="">Select a household</option>
                {households.map((h) => (
                  <option key={h.id} value={h.id}>{h.head_name} · {h.location}</option>
                ))}
              </select>
              <button
                type="button"
                onClick={() => setShowNewHousehold(true)}
                className="text-teal text-sm font-medium"
              >
                + Add new household
              </button>
            </>
          ) : (
            <div>
              <input
                type="text"
                placeholder="Head of household name"
                value={newHeadName}
                onChange={(e) => setNewHeadName(e.target.value)}
                className="w-full h-12 px-3 rounded-xl border border-ink/15 text-ink mb-3"
              />
              <input
                type="text"
                placeholder="Location (e.g. Manyatta B)"
                value={newLocation}
                onChange={(e) => setNewLocation(e.target.value)}
                className="w-full h-12 px-3 rounded-xl border border-ink/15 text-ink mb-3"
              />
              <div className="flex gap-2">
                <button
                  type="button"
                  onClick={handleAddHousehold}
                  className="flex-1 h-11 rounded-xl bg-teal text-white font-medium text-sm"
                >
                  Save household
                </button>
                <button
                  type="button"
                  onClick={() => setShowNewHousehold(false)}
                  className="h-11 px-4 rounded-xl border border-ink/15 text-ink/60 text-sm"
                >
                  Cancel
                </button>
              </div>
            </div>
          )}
        </div>

        <form onSubmit={handleSubmitVisit} className="bg-white rounded-2xl border border-ink/10 p-6">
          <h2 className="font-heading font-bold text-ink mb-4">Visit details</h2>

          <label className="block text-sm font-medium text-ink mb-2">Child's weight (kg)</label>
          <input
            type="number"
            step="0.1"
            value={weight}
            onChange={(e) => setWeight(e.target.value)}
            className="w-full h-12 px-3 rounded-xl border border-ink/15 text-ink mb-4"
            placeholder="e.g. 11.5"
          />

          <label className="block text-sm font-medium text-ink mb-2">Child's age (months)</label>
          <input
            type="number"
            value={ageMonths}
            onChange={(e) => setAgeMonths(e.target.value)}
            className="w-full h-12 px-3 rounded-xl border border-ink/15 text-ink mb-4"
            placeholder="e.g. 18"
          />

          <label className="flex items-center gap-2 mb-4">
            <input
              type="checkbox"
              checked={vaccinationDone}
              onChange={(e) => setVaccinationDone(e.target.checked)}
              className="w-5 h-5"
            />
            <span className="text-sm text-ink">Vaccination up to date</span>
          </label>

          <label className="block text-sm font-medium text-ink mb-2">Notes</label>
          <textarea
            value={notes}
            onChange={(e) => setNotes(e.target.value)}
            className="w-full h-20 px-3 py-2 rounded-xl border border-ink/15 text-ink mb-4"
            placeholder="Anything worth noting..."
          />

          {error && <p className="text-red-600 text-sm mb-3">{error}</p>}
          {success && <p className="text-teal text-sm mb-3">{success}</p>}

          <button
            type="submit"
            disabled={loading}
            className="w-full h-14 rounded-xl bg-teal hover:bg-teal-dark text-white font-heading font-bold disabled:opacity-60"
          >
            {loading ? 'Saving...' : 'Log visit'}
          </button>
        </form>
      </div>
    </div>
  )
}
