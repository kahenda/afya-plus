const BASE_URL = 'http://localhost:8080'

export async function createHousehold(data) {
  const res = await fetch(`${BASE_URL}/households`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(data),
  })
  if (!res.ok) throw new Error('Failed to create household')
  return res.json()
}

export async function getHouseholds() {
  const res = await fetch(`${BASE_URL}/households`)
  if (!res.ok) throw new Error('Failed to fetch households')
  return res.json()
}

export async function createVisit(data) {
  const res = await fetch(`${BASE_URL}/visits`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(data),
  })
  if (!res.ok) throw new Error('Failed to log visit')
  return res.json()
}

export async function getFlags() {
  const res = await fetch(`${BASE_URL}/flags`)
  if (!res.ok) throw new Error('Failed to fetch flags')
  return res.json()
}
