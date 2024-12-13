import userService from "../services/users";
import RegisterForm from "../components/RegisterForm";
import { Authenticator } from "../Authenticator";
import { useTranslation } from "react-i18next";
import { useFetcher } from "react-router-dom";

const RegisterPage = () => {
  const { t, i18n } = useTranslation();
  const fetcher = useFetcher();

  const createNewUser = (userObject) => {
    userService
      .create(userObject)
      .then((data) => {
        Authenticator.login(data);
        alert(t("alert.usercreated"));
        fetcher.submit({ data: data }, { action: "/login", method: "post" });
      })
      .catch((error) => {
        const key =
          "error." + error.response.data.toLowerCase().replace(/[^a-z]/g, "");
        alert(
          t("error.error") +
            ": " +
            (i18n.exists(key) ? t(key) : error.response.data),
        );
      });
  };

  return (
    <div className="logged-out-container">
      <div className="container">
        <div className="row justify-content-center">
          <div className="col-12 col-md-8 col-lg-6 col-xl-5">
            <div className="card" style={{ borderRadius: "1rem" }}>
              <div className="card-body p-5">
                <h2 className="mb-5 text-center">{t("title.register")}</h2>
                <RegisterForm createNewUser={createNewUser} />
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default RegisterPage;
