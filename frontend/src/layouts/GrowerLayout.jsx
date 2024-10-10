import { Outlet, NavLink } from 'react-router-dom'
//import './Grower.css'

const navigationBar = () => {
  return (
    <div className="nav-container">
      <nav>
        <ul>
          <li>
            <NavLink to="/grower" end>
              Home
            </NavLink>
          </li>
          <li>
            <NavLink to="/grower/flowers">Flowers</NavLink>
          </li>
        </ul>
      </nav>
    </div>
  )
}

const GrowerLayout = () => {
  return (
    <div className="layout-container">
      <header className="header">
        <h1>Slowers</h1>
      </header>
      <div className="content">
        <aside className="side-container">
          <h1>Navigation</h1>
          {navigationBar()}
        </aside>
        <main className="main-container">
          <Outlet />
        </main>
      </div>
    </div>
  )
}

export default GrowerLayout
