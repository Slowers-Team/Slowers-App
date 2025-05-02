import { useState, useEffect } from "react"
import userService from "../services/users"
import UserInfo from "../components/UserInfo"
import { useTranslation } from 'react-i18next'
import CenteredCard from "../components/CenteredCard"


const UserPage = () => {
  const [user, setUser] = useState({})
  const { t, i18n } = useTranslation()

  useEffect(() => {
    userService.get().then((user) => setUser(user))
  }, []);
  

  return (
    <CenteredCard>
      <h2>{t('menu.profile')}</h2>
      <UserInfo user={user} />
    </CenteredCard>
  )
}

export default UserPage;