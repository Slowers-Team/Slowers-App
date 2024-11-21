import { Navigate, Outlet, useRouteLoaderData } from 'react-router-dom'

export const ProtectedRoute = () => {
  let { isLoggedIn } = useRouteLoaderData("root")

  if (!isLoggedIn) {
    return <Navigate to="/login" replace />
  }

  return <Outlet />
}

export default ProtectedRoute
