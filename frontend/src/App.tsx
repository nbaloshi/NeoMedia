import { BrowserRouter, Navigate, Route, Routes } from "react-router-dom"
import LoginPage from "./pages/Login"
import RegisterPage from "./pages/Register"
import { SessionProvider } from "./context/SessionProvider"
import { PrivateRoute } from "./components/PrivateRoute"
import HomePage from "./pages/Home"

function App() {
  return (
    <SessionProvider>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<Navigate to="/login" replace />} />
          <Route path="/login" element={<LoginPage />} />
          <Route path="/register" element={<RegisterPage />} />
          <Route path="/home" element={
            <PrivateRoute>
              <HomePage />
            </PrivateRoute>
          } />
        </Routes>
      </BrowserRouter>
    </SessionProvider>
  )
}

export default App