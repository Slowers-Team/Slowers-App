import { useTranslation } from "react-i18next";
import { Container } from 'react-bootstrap'


const HomePage = () => {
  const { t, i18n } = useTranslation();


  return (
    <Container>
      <h2>{t('title.home')}</h2>
      <br/>
      <p>Welcome to Slowers App</p>
    </Container>
  );
};

export default HomePage;
