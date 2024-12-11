import { useNavigate } from "react-router-dom"
import TermsWindow from "../components/TermsWindow"
import { useTranslation } from "react-i18next"
import { Container } from "react-bootstrap"

const TermsPage = () => {
  const navigate = useNavigate()
  const { t, i18n } = useTranslation()

  const handleBack = () => {
    navigate(-1)
  }

  return (
    <Container>
      <h2 className="mx-3 my-3">{t("title.terms")}</h2>
      <TermsWindow />
      <br />
      <button onClick={handleBack} className="mx-2 btn btn-light">{t("button.back")}</button>
    </Container>
  )
}

export default TermsPage
