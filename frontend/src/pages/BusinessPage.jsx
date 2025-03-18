import { useTranslation } from "react-i18next";
import { Container } from 'react-bootstrap'


const BusinessPage = () => {
  const { t, i18n } = useTranslation();


  return (
    <Container>
      <h2>{t('title.businesspage')}</h2>
      <br/>
      <p>Here comes the content of the business page</p>
    </Container>
  );
};

export default BusinessPage;