import { useState, useEffect } from 'react'
import userService from '../services/users'
import UserInfo from '../components/UserInfo' 

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

  const switchRole = () => (user.role === 'grower' ? 'retailer' : 'grower')

  return (
    <div className='mx-3'>
      <UserInfo user={user} handleRoleSwitch={handleRoleSwitch} switchedRole={switchRole(user.role)} />
    </div>
  )
}

export default UserPage
