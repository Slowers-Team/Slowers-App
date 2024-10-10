import { useState, useEffect } from 'react'
import userService from '../services/users'
import '../Misc.css'

const UserPage = () => {
  const [user, setUser] = useState({})

  useEffect(() => {
    userService.get().then(user => setUser(user))
  }, [])

  /* const deleteFlower = flowerObject => {
    if (window.confirm(`Are you sure you want to delete flower ${flowerObject.name}?`)) {
      flowerService.remove(flowerObject._id).then(response => {
        console.log(response)
        setFlowers(l => l.filter(item => item._id !== flowerObject._id))
      })
    }
  } */

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
          <td><button>Switch to {switchRole()}</button></td>
        </tr>
        </tbody>
      </table>
    </div>
  )
}

export default UserPage
