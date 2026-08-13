export function generateVisitId() {
  return 'visit-' + Date.now() + '-' + Math.random().toString(36).slice(2, 8)
}
