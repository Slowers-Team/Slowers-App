import { useState, useEffect } from 'react'
import userService from '../services/users'
import UserInfo from '../components/UserInfo' 
import { useSubmit } from 'react-router-dom'

const UserPage = () => {
  const [user, setUser] = useState({})
  const submit = useSubmit()

  useEffect(() => {
    userService.get().then(user => setUser(user))
  }, [])

  const handleRoleSwitch = () => {
    const newRole = switchRole();
    userService.setRole(newRole).then(_ => {
      setUser({...user, role: newRole})
      submit({role: newRole}, {action: "/user", method: "post"})
    })  }

  const switchRole = () => (user.role === 'grower' ? 'retailer' : 'grower')

  return (
    <div className='mx-3'>
      <UserInfo user={user} handleRoleSwitch={handleRoleSwitch} switchedRole={switchRole(user.role)} />
    </div>
  )
}

export default UserPage
