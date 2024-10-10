import { useState, useEffect } from 'react'
import userService from '../services/users'
import '../Misc.css'

const UserPage = ({setDefaultRole}) => {
  const [user, setUser] = useState({})

  useEffect(() => {
    userService.get().then(user => setUser(user))
  }, [])

  const handleRoleSwitch = () => {
    const newRole = switchRole();
    userService.setRole(newRole).then(_ => {
      setUser({...user, role: newRole})
      localStorage.setItem('role', newRole);
      setDefaultRole(newRole)
    })  }

  const switchRole = () => {
    if (user.role == "grower") {
      return "retailer"
    } else {
      return "grower"
    }
  }

  return (
    <div>
      <table>
        <tbody>
        <tr>
          <td>Username</td>
          <td>{ user.username }</td>
        </tr>
        <tr>
          <td>Email</td>
          <td>{user.email}</td>
        </tr>
        <tr>
          <td>Role</td>
          <td>{user.role}</td>
          <td><button onClick={handleRoleSwitch}>Switch to {switchRole()}</button></td>
        </tr>
        </tbody>
      </table>
    </div>
  )
}

export default UserPage
