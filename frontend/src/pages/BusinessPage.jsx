import { useState, useEffect } from "react"
import userService from "../services/users"
import businessService from "../services/business"
import BusinessInfo from "../components/BusinessInfo"
import { useTranslation } from 'react-i18next'
import { Authenticator } from "../Authenticator"
import CreateBusinessForm from "../components/CreateBusinessForm"


const BusinessPage = () => {
  const [user, setUser] = useState({})
  const [business, setBusiness] = useState({})
  const { t, i18n } = useTranslation()
  const [errorMessage, setErrorMessage] = useState("")
  const [ designation, setDesignation ] = useState(Authenticator.designation)

  useEffect(() => {
    if (designation === 'owner' || designation === 'employee') {
      businessService.get().then((business) => setBusiness(business))
    }
  }, [designation]);

  useEffect(() => {
    userService.get().then((user) => setUser(user))
  }, []);

  const createNewBusiness = async (businessObject) => {
    try {
      await businessService.create(businessObject, user.Email)
      console.log("creating business successful")
      const membership = await userService.getDesignation()
      Authenticator.setDesignation(membership.Designation)
      const updatedBusiness = await businessService.get();
      setBusiness(updatedBusiness);
    } catch (error) {
      const key = "error." + error.response.data.toLowerCase().replace(/[^a-z]/g, "");
      console.log(error.response.data)
      setErrorMessage(i18n.exists(key) ? t(key) : error.response.data);
    }
    setDesignation(Authenticator.designation)
    setBusiness(business)
  }


  return (
    <div className="m-3">
      <div className="row justify-content-center">
        <div className="col-12 col-md-12 col-lg-12 col-xl-8">
          <div className="card" style={{ borderRadius: "1rem" }}>
            <div className="card-body p-5">
              {(designation === 'owner' || designation === 'employee')  
                ? (<BusinessInfo business={business} />)
                : (<CreateBusinessForm createNewBusiness={createNewBusiness} />)}
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default BusinessPage