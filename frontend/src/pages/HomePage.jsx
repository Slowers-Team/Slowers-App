import { useTranslation } from "react-i18next";
import { Container } from 'react-bootstrap'


const HomePage = () => {
  const { t, i18n } = useTranslation();


  return (
    <Container>
      <h2>{t('title.welcome')}</h2>
    </Container>
  );
};

export default HomePage;
