import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import apiConfig from './src/services/apiConfig'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    host: true,
    strictPort: true,
    port: process.env.REACT_APP_PORT
  }
})
