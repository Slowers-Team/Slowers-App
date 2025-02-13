import { useTranslation } from "react-i18next";
import { Container } from 'react-bootstrap'


const BusinessOwnerPage = () => {
  const { t, i18n } = useTranslation();


  return (
    <Container>
      <h2>{t('title.businessowner')}</h2>
      <br/>
      <p>Here comes the content of the business page</p>
    </Container>
  );
};

export default BusinessOwnerPage;