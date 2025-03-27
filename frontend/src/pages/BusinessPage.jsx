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
    <Container className="m-3">
      <BusinessInfo business={business} />
    </Container>
  )
}

export default BusinessPage;