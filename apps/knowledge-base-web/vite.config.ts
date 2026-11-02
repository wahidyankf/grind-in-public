import { tanstackStart } from "@tanstack/react-start/plugin/vite";
import react from "@vitejs/plugin-react";
import { defineConfig } from "vite";

// The API runs on its own port. The dev and preview servers forward /api to it,
// so the browser calls one origin and no CORS setup is needed.
const proxy = {
  "/api": { target: "http://localhost:8200" },
};

export default defineConfig({
  // TanStack Start runs in SPA mode, so pages render in the browser only, as CSR.
  plugins: [tanstackStart({ spa: { enabled: true } }), react()],
  server: { port: 3200, proxy },
  preview: { port: 3200, proxy },
});
