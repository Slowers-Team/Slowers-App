import userService from "../services/users";
import RegisterForm from "../components/RegisterForm";
import { useTranslation } from "react-i18next";
import { useNavigate } from "react-router-dom";
import { useState } from "react";
import Notification from "../components/Notification";

const RegisterPage = () => {
  const { t, i18n } = useTranslation();
  const [errorMessage, setErrorMessage] = useState("")
  const navigate = useNavigate();
  const [message, setMessage] = useState("");

  const createNewUser = (userObject) => {
    return userService
      .create(userObject)
      .then(() => {
        setMessage(t("message.registersuccessful"));
        setTimeout(() => {
          setMessage(null);
          navigate("/login");
        }, 3000)
      })
      .catch((error) => {
        const key =
          "error." + error.response.data.toLowerCase().replace(/[^a-z]/g, "");
        
        setErrorMessage(i18n.exists(key) ? t(key) : error.response.data);
        throw error;
      });
  };

  return (
    <div className="logged-out-container">
      <div className="container">
        <Notification message={message}/>
        <br />
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
