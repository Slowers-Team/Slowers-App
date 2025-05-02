import TermsWindow from "../components/TermsWindow"
import { useTranslation } from "react-i18next"
import CenteredCard from "../components/CenteredCard"


const TermsPage = () => {
  const { t, i18n } = useTranslation()

  return (
    <CenteredCard>
      <h2 className="mx-3 my-3">{t("title.terms")}</h2>
      <TermsWindow />
      <br />
    </CenteredCard>
  )
}

export default TermsPage
