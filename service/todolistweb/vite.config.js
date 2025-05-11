import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

export default defineConfig({
  plugins: [react()],
  server: {
    port: 8888,
    proxy: {
      '/v1': 'http://localhost:8000'
    }
  }
})