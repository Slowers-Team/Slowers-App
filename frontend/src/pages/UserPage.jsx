import { useState, useEffect } from "react"
import userService from "../services/users"
import UserInfo from "../components/UserInfo"
import { useTranslation } from 'react-i18next'


const UserPage = () => {
  const [user, setUser] = useState({})
  const { t, i18n } = useTranslation()

  useEffect(() => {
    userService.get().then((user) => setUser(user))
  }, []);
  

  return (
    <div className="m-3">
      <div className="row justify-content-center">
        <div className="col-12 col-md-12 col-lg-12 col-xl-8">
          <div className="card" style={{ borderRadius: "1rem" }}>
            <div className="card-body p-5">
              <h2>{t('menu.profile')}</h2>
              <UserInfo user={user} />
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default UserPage;