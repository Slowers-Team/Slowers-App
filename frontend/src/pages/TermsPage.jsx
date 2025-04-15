import { useNavigate } from "react-router-dom"
import TermsWindow from "../components/TermsWindow"
import { useTranslation } from "react-i18next"


const TermsPage = () => {
  //const navigate = useNavigate()
  const { t, i18n } = useTranslation()

  return (
    <div className="m-3">
      <div className="row justify-content-center">
        <div className="col-12 col-md-12 col-lg-12 col-xl-8">
          <div className="card" style={{ borderRadius: "1rem" }}>
            <div className="card-body p-5">
              <h2 className="mx-3 my-3">{t("title.terms")}</h2>
              <TermsWindow />
              <br />
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default TermsPage

console.log("jee")
