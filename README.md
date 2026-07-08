
Afya Plus
Offline-first household visit tracking for Community Health Workers.

Afya Plus is a lightweight tool that lets Community Health Workers (CHWs) log household visits without needing an internet connection, automatically flags at-risk cases (like malnutrition or overdue vaccinations), and syncs everything to a central dashboard the moment connectivity returns — so supervisors can follow up quickly instead of waiting on paper reports.

The Problem
Community Health Workers are often the first point of contact for maternal and child health in underserved areas — but most still rely on paper forms. This means at-risk cases can take days or weeks to reach a clinic, by which point early intervention windows may already be closed. Afya Plus exists to close that gap.

How It Works
A CHW logs a household visit on their phone — fully offline.
The app checks the data against simple risk rules (e.g. malnutrition thresholds, overdue vaccinations) and flags the household if needed — no internet required.
Once the phone regains signal, all queued visits sync automatically in the background.
A supervisor sees flagged cases on a dashboard, sorted by urgency, and assigns follow-up.
Each flagged case moves through a status trail — Flagged → Under Review → Action Taken → Resolved — so nothing gets logged and forgotten.
Tech Stack
Layer	Tech
Backend	Go + Gin
Database	PostgreSQL (Neon.tech)
Frontend	React + Tailwind CSS, built as a PWA
Offline storage	IndexedDB + background sync queue
Project Structure
afya-plus/
├── backend/     # Go API — auth, visits, households, flags, sync
├── frontend/    # React PWA — CHW and supervisor interfaces
└── docs/        # Design notes and decisions
Status
🚧 Actively in development. Currently building the core data model and API.

 Repo, structure, and backend scaffold
 Connected to Neon Postgres
 Health check endpoint (GET /health)
 Full schema — users, households, visits, flags
 CHW authentication (phone + PIN)
 Visit logging with offline support
 Automatic risk-flagging rules
 Sync queue and sync API
 Supervisor dashboard
Why Open Source
This project is public and MIT licensed because tools like this are more useful the more people can adapt them to their own context — different health programs, different regions, different risk criteria. Contributions, issues, and ideas are welcome.

Author
Built by Philip Damwanza Ahenda — apprentice developer at Zone01 Kisumu. More on Dev.to · Portfolio

License
MIT — see LICENSE for details.