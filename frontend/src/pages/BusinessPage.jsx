import { useState, useEffect } from "react"
//import userService from "../services/users"
import businessService from "../services/business"
import BusinessInfo from "../components/BusinessInfo"
import { Container } from "react-bootstrap"
import { useTranslation } from 'react-i18next'


const BusinessPage = () => {
  const [business, setBusiness] = useState({})
  const { t, i18n } = useTranslation()
  const [errorMessage, setErrorMessage] = useState("")

  useEffect(() => {
    businessService.get().then((business) => setBusiness(business))
  }, []);

  return (
    <div className="m-3">
      <div className="row justify-content-center">
        <div className="col-12 col-md-12 col-lg-12 col-xl-8">
          <div className="card" style={{ borderRadius: "1rem" }}>
            <div className="card-body p-5">
              <BusinessInfo business={business} />
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default BusinessPage;