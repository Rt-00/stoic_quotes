# 📚 Stoic Quotes

A Fullstack application to search, filter, and explore book quotes, with a **Go (Gin)** backend and **React (Vite + Tailwind + Shadcn)** frontend.

---

## 🛠️ Technologies

### Backend:

- [Go] (https://go.dev/)
- [Gin Web Framework] (https://gin-gonic.com/)
- [SQLite] (https://www.sqlite.org/index.html)
- CORS Middleware
- Clean Architecture (Services, Repositories)

### Frontend:

- [React + Vite](https://vitejs.dev/)
- [TailwindCSS](https://tailwindcss.com/)
- [shadcn/ui](https://ui.shadcn.dev/) (Dialog, Button, Card, Input, etc.)
- Axios for API requests
- Modal for random quotes
- Controlled Pagination
- Scrollbars for large content

---

## 🚀 Getting Started Locally

### Prerequisites:

- Go 1.21+
- Node.js 18+
- SQLite3

### Clone the repository:

```bash
git clone git@github.com:Rt-00/stoic_quotes.git
cd stoic_quotes
```

### Backend (Go):

**1.** Install dependencies:

```bash
go mod tidy
```

**2.** Start the server:

```bash
go run main.go
```

Backend will be available at: `http://localhost:3000`.

### Frontend (React):

**1.** Install dependencies:

```bash
npm install
```

**2.** Start the frontend:

```bash
npm run dev
```

Frontend will be available at: `http://localhost:5173`.

---

## 📝 Features

- ✅ List quotes with filters by Author, Book, and Text.
- ✅ Pagination with layout stability (no shifting).
- ✅ Debounced filtering to avoid excessive API calls.
- ✅ Random Quote modal popup.
- ✅ Scrollbar for long quotes.
- ✅ Backend full-text search index (FTS5 in SQLite).
- ✅ Modular Clean Architecture (Go backend).

---

## 📦 Folder Structure

```
backend
├── db
│   └── db.go
├── go.mod
├── go.sum
├── handlers
│   └── quote_handler.go
├── main.go
├── models
│   └── quote.go
├── quotes.db
├── repository
│   └── quote_repository.go
├── service
│   └── import_service.go
└── stoic_quotes_full.csv
```

```
frontend/src
├── App.tsx
├── assets
│   └── react.svg
├── components
│   ├── modal.tsx
│   ├── pagination.tsx
│   └── ui
│   ├── button.tsx
│   ├── card.tsx
│   ├── dialog.tsx
│   └── input.tsx
├── index.css
├── lib
│   └── utils.ts
├── main.tsx
├── pages
│   └── quote-list-page.tsx
└── vite-env.d.ts
```
---

## 📄 API Endpoints

| Method | Route            | Description                                                  |
| ------ | ---------------- | ------------------------------------------------------------ |
| GET    | `/quotes`        | List quotes (supports `author`, `book`, `q`, `page` filters) |
| GET    | `/quotes/random` | Returns a random quote                                       |

---

## 💡 Future Improvements

- [ ] Sort search results by relevance.
- [ ] Infinite Scroll (optional enhancement).
- [ ] Autocomplete for authors and books.
- [ ] Deploy on Railway / Vercel.
