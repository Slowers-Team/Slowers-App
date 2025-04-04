import { useTranslation } from 'react-i18next'

const UserInfo = ({ user }) => {
  const { t, i18n } = useTranslation()

  
  return (
    <div id="userinfo">
      <table>
        <tbody>
        <tr>
          <td>{t('user.data.username')}</td>
          <td>{ user.Username }</td>
        </tr>
        <tr>
          <td>{t('user.data.email')}</td>
          <td>{user.Email}</td>
        </tr>
        {/* <tr>
          <td>{t('user.data.role')}</td>
          <td id="roleValue">{t(`role.${user.role}`)}</td>
        </tr> */}
        </tbody>
      </table>
    </div>
  )
}

export default UserInfo
