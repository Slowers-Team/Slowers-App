import { useNavigate } from "react-router-dom"
import TermsWindow from "../components/TermsWindow"

const TermsPage = () => {
  const navigate = useNavigate()

  const handleBack = () => {
    navigate(-1)
  }

  return (
    <div>
      <h1>Terms and Conditions</h1>
      <TermsWindow />
      <button onClick={handleBack}>Back</button>
    </div>
  )
}

export default TermsPage
