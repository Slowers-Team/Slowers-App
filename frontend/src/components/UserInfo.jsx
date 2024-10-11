const UserInfo = ({ user, handleRoleSwitch, switchedRole }) => {
  return (
    <div id="userinfo">
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
          <td><button onClick={handleRoleSwitch}>Switch to {switchedRole}</button></td>
        </tr>
        </tbody>
      </table>
    </div>
  )
}

export default UserInfo
