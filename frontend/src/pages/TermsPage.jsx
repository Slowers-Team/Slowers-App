import { useNavigate } from "react-router-dom"
import TermsWindow from "../components/TermsWindow"
import { useTranslation } from "react-i18next"

const TermsPage = () => {
  const navigate = useNavigate()
  const { t, i18n } = useTranslation()

  const handleBack = () => {
    navigate(-1)
  }

  return (
    <div>
      <h1>{t("title.terms")}</h1>
      <TermsWindow />
      <br />
      <button onClick={handleBack}>{t("button.back")}</button>
    </div>
  )
}

export default TermsPage
