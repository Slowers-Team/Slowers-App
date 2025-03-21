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

  const handleCreateBusiness = async (props) => {
    const updatedRole = props.type === "grower" ? "growerowner" : "retailerowner";
    userService.setRole(updatedRole).then((_) => {
      setUser({ ...user, role: updatedRole })
      Authenticator.setRole(updatedRole)
    })
    businessService.create(props)
      .then(
        console.log("creating business successful")
      )
  }
  
  useEffect(() => {
    userService.get().then((user) => setUser(user))
  }, []);

  return (
    <Container className="m-3">
      <h2>{t('menu.profile')}</h2>
      <UserInfo user={user} />
      <br/>
      <CreateBusinessForm onSubmit={handleCreateBusiness} />
    </Container>
  )
}

export default UserPage;