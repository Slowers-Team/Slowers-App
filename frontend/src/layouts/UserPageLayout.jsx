import { Outlet } from 'react-router-dom'


const UserPageLayout = () => {
  return (
    <div>
      <div>
        <main className="main-container">
          <Outlet />
        </main>
      </div>
    </div>
  )
}

export default UserPageLayout
