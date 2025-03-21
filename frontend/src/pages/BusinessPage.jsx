import { useTranslation } from "react-i18next";
import { Container } from 'react-bootstrap'


const BusinessPage = () => {
  const { t, i18n } = useTranslation();


  return (
    <Container>
      <center>
      <div className="business-name-container">
      <h2>Business name</h2>
      <br></br>
      <h5>{t('businessform.fieldname.businesstype')}: Grower</h5>
      <h5>{t('businessform.fieldname.phonenumber')}: +358 45 2789654</h5>
      <h5>{t('businessform.fieldname.email')}: business@email.com</h5>
      <h5>{t('businessform.fieldname.postalcode')}: 00000</h5>
      <h5>{t('businessform.fieldname.city')}: Kerava</h5> 
      <h5>{t('businessform.fieldname.address')}: Mustikkatie 8 A3</h5>
      <h5>business id code: 23998</h5>
      <h5>{t('businessform.fieldname.created')}: 23.5.2021</h5>

      <br></br>
      <br></br>

      <h4>{t('businessform.fieldname.additionalinfo')}:</h4>
        <br></br>
        <div class="about-business-container">About business</div>
        <br></br>
      </div>
      </center>
      <br></br>


    </Container>
  );
};

export default BusinessPage;