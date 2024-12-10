import { useTranslation } from 'react-i18next'

const UserInfo = ({ user, handleRoleSwitch, switchedRole }) => {
  const { t, i18n } = useTranslation()

  const translateRoleButton = (role) => {
    if (role == "grower") {
      return t('button.switchToGrower')
    } else {
      return t('button.switchToRetailer')
    }
  }
  
  return (
    <div id="userinfo">
      <table>
        <tbody>
        <tr>
          <td>{t('user.data.username')}</td>
          <td>{ user.username }</td>
        </tr>
        <tr>
          <td>{t('user.data.email')}</td>
          <td>{user.email}</td>
        </tr>
        <tr>
          <td>{t('user.data.role')}</td>
          <td>{t(`role.${user.role}`)}</td>
          <td><button onClick={handleRoleSwitch} className='flower-button'>{translateRoleButton(switchedRole)}</button></td>
        </tr>
        </tbody>
      </table>
    </div>
  )
}

export default UserInfo
