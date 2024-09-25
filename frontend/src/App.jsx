import './App.css'
import RegisterPage from './pages/RegisterPage'
import HomePage from './pages/HomePage'
import SitePage from './pages/SitePage'
import { BrowserRouter as Router, Routes, Route, Link } from 'react-router-dom'

const App = () => {

  const padding = {
    padding: 5
  }

  return (
    <div>
      <Router>
        <div>
          <nav>
            <Link style={padding} to="/">Home</Link>
            <Link style={padding} to="/register">Register</Link>
            <Link style={padding} to="/site">Site</Link>
          </nav>
          <Routes>
            <Route path="/" element={<HomePage />} />
            <Route path="/register" element={<RegisterPage />} />
            <Route path="/site" element={<SitePage />} />
          </Routes>
        </div>
      </Router>
    </div>
  )
}

export default App
