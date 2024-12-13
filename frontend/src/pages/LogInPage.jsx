import { useNavigate, useFetcher } from "react-router-dom";
import LogInForm from "../components/LogInForm";
import { useTranslation } from "react-i18next";
import { Authenticator } from "../Authenticator";

const LogInPage = () => {
  const navigate = useNavigate();
  const { t, i18n } = useTranslation();
  let fetcher = useFetcher();

  const handleLogin = (data) => {
    Authenticator.login(data);
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
