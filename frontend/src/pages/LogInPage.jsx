import { useNavigate, useFetcher } from "react-router-dom";
import LogInForm from "../components/LogInForm";
import { useTranslation } from "react-i18next";
import { Authenticator } from "../Authenticator";
import userService from "../services/users"
import businessService from "../services/business"

const LogInPage = () => {
  const navigate = useNavigate();
  const { t, i18n } = useTranslation();
  let fetcher = useFetcher();

  const handleLogin = async (data) => {
    Authenticator.login(data);
    try {
      const membership = await userService.getDesignation();
  
      if (membership && Object.keys(membership).length > 0) {
        Authenticator.setDesignation(membership.Designation);
      }
      } catch (error) {
      console.error("Error fetching designation:", error.response ? error.response.data : error.message);
    }
    try {
      const business = await businessService.get();
  
      if (business && Object.keys(business).length > 0) {
        Authenticator.setBusinessType(business.Type);
      }
      } catch (error) {
      console.error("Error fetching business type:", error.response ? error.response.data : error.message);
    }
    fetcher.submit({ data: data }, { action: "/login", method: "post" });
  };

  return (
    <div className="logged-out-container">
      <div className="container">
        <div className="row justify-content-center">
          <div className="col-12 col-md-8 col-lg-6 col-xl-4">
            <div className="card" style={{ borderRadius: "1rem" }}>
              <div className="card-body p-5">
                <h2 className="mb-5 text-center">{t("title.login")}</h2>
                <LogInForm onLogin={handleLogin} />
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default LogInPage;
