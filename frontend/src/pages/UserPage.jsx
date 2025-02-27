import { useState, useEffect } from "react"
import userService from "../services/users"
import UserInfo from "../components/UserInfo"
import { Authenticator } from "../Authenticator"
import { Container } from "react-bootstrap"
import { useTranslation } from 'react-i18next'
// import CreateBusinessForm from "../components/CreateBusinessForm"


const UserPage = () => {
  const [user, setUser] = useState({})
  const { t, i18n } = useTranslation()
  const [businessName, setBusinessName] = useState('')
  const [type, setType] = useState ('')

  const handleCreateBusiness = async (event) => {
    event.preventDefault()
    const updatedRole = type === "grower" ? "growerowner" : "retailerowner";
    userService.setRole(updatedRole).then((_) => {
      setUser({ ...user, role: updatedRole })
      Authenticator.setRole(updatedRole)
    })
   }

  useEffect(() => {
    userService.get().then((user) => setUser(user))
  }, []);

  return (
    <Container className="m-3">
      <h2>{t('menu.profile')}</h2>
      <UserInfo
        user={user}
      />
      <br/>
      <h2>{t('title.createbusiness')}</h2>
      <div>
        <form onSubmit={handleCreateBusiness}>
          <div>
            <table>
              <tbody>
                <tr>
                  <td>{t('businessform.fieldname.businessname')}</td>
                  <td>
                    <input 
                      className="form-control"
                      value={businessName}
                      placeholder={t('businessform.input.businessname')}
                      onChange={event => setBusinessName(event.target.value)}
                      style={{ width: "400px"}}
                    />
                  </td>
                </tr>
                <tr>
                  <td>{t('businessform.fieldname.businesstype')}</td>
                  <td>
                    <input type="radio" id="growerSelector" className='btn-check' name="typeSelector" value="grower" checked={type === "grower"} onChange={event => setType(event.target.value)} />
                    <label className='btn btn-outline-secondary' style={{ marginRight: "0.5rem" }} htmlFor="growerSelector" >{t('button.grower')}</label>
                    <input type="radio" id="retailerSelector" className='btn-check' name="typeSelector" value="retailer" checked={type === "retailer"}  onChange={event => setType(event.target.value)} />
                    <label className='btn btn-outline-secondary' style={{ marginRight: "0.5rem" }} htmlFor="retailerSelector" >{t('button.retailer')}</label>
                  </td>
                </tr>
              </tbody>
            </table>
            <br/>
            <button type="submit" className='custom-button'>{t('button.createbusiness')}</button>
          </div>
        </form>
      </div> 
    </Container>
  )
}

export default UserPage;
