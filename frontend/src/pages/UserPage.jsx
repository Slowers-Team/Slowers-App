import { useState, useEffect } from "react"
import userService from "../services/users"
import businessService from "../services/business"
import UserInfo from "../components/UserInfo"
import { Authenticator } from "../Authenticator"
import { Container } from "react-bootstrap"
import { useTranslation } from 'react-i18next'
import CreateBusinessForm from "../components/CreateBusinessForm"


const UserPage = () => {
  const [user, setUser] = useState({})
  const { t, i18n } = useTranslation()
  const [errorMessage, setErrorMessage] = useState("")

  useEffect(() => {
    userService.get().then((user) => setUser(user))
  }, []);

  const createNewBusiness = async (businessObject) => {
    const updatedRole = businessObject.type === "grower" ? "growerowner" : "retailerowner";
    try {
      await businessService.create(businessObject, user.email)
        .then(
          console.log("creating business successful")
        )
      userService.setRole(updatedRole).then((_) => {
        setUser({ ...user, role: updatedRole })
        Authenticator.setRole(updatedRole)
      })
    } catch (error) {
      const key = "error." + error.response.data.toLowerCase().replace(/[^a-z]/g, "");

      setErrorMessage(i18n.exists(key) ? t(key) : error.response.data);
      throw error;
    }
  }
  

  return (
    <Container className="m-3">
      <h2>{t('menu.profile')}</h2>
      <UserInfo user={user} />
      <br/>
      <CreateBusinessForm createNewBusiness={createNewBusiness} />
    </Container>
  )
}

export default UserPage;