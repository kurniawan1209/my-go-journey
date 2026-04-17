Excellent choice. A high-performance REST API is the perfect playground for Go because concurrency and web serving are built directly into its DNA. [cite_start]We are going to bypass the heavy frameworks you might be used to, like Laravel or FastAPI[cite: 56], and see how Go handles web traffic almost entirely with its standard library. 

[cite_start]Since you have experience maintaining and restructuring transactional databases[cite: 58], our MVP will be a **Transactional Order API**. It will handle incoming orders, validate them, and write them to a PostgreSQL database—fast.

Let's kick off Sprint 1. We are keeping the feedback loop tight so you don't have time to get bored.

---

### 🎫 TICKET: GO-001
**Title:** Initialize Project and Create Health Check Endpoint
**Story Points:** 1 (Should take you less than 30 minutes)
**Objective:** Set up the basic foundation of the Go application and prove it can serve web traffic.

**Acceptance Criteria:**
1. Initialize a new Go module (name it something like `order-api`).
2. Create an HTTP server that listens on port `8080`.
3. Create a single `GET` endpoint at the path `/health`.
4. The endpoint must return a JSON response: `{"status": "ok", "message": "API is running"}`.

**PM Notes & Constraints (Read Before Coding):**
* **No External Frameworks:** Do *not* use heavy frameworks like Gin or Fiber just yet. I want you to use Go's built-in `net/http` and `encoding/json` standard libraries. You need to understand how Go natively handles requests before we add abstractions.
* **Hints to Google:** Look up `go mod init`, `http.HandleFunc`, and `http.ListenAndServe`. 

---

### Your Next Step

Your task is to write the code to satisfy this ticket, test it locally (hit `localhost:8080/health` in your browser or Postman), and then paste your `main.go` file here for my code review. 

**Drop your code below when you are ready for the code review, or let me know if you hit a blocker setting up the server?**