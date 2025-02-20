import { useState, useEffect } from "react";
import userService from "../services/users";
import UserInfo from "../components/UserInfo";
import { Authenticator } from "../Authenticator";
import { Container } from "react-bootstrap";
import CreateBusinessForm from "../components/CreateBusinessForm";
import { useTranslation } from 'react-i18next';

const UserPage = () => {
  const [user, setUser] = useState({});
  const { t, i18n } = useTranslation()

  useEffect(() => {
    userService.get().then((user) => setUser(user));
  }, []);

  const switchRole = () => (user.role === "grower" ? "retailer" : "grower");

  return (
    <Container className="m-3">
      <h2>{t('menu.profile')}</h2>
      <UserInfo
        user={user}
      />
      <br/>
      <h2>{t('title.createbusiness')}</h2>
      <CreateBusinessForm />
    </Container>
  );
};

export default UserPage;
