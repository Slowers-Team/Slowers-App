import { useTranslation } from "react-i18next";
import { Container } from 'react-bootstrap'


const BusinessInfo = ({ business }) => {
  const { t, i18n } = useTranslation();


  return (
    <Container>
      <div className="business-name-container">
      <h2>{ business.BusinessName }</h2>
      <br></br>
      <h5>{t('businessform.fieldname.businesstype')}: { business.Type }</h5>
      <h5>{t('businessform.fieldname.phonenumber')}: { business.PhoneNumber }</h5>
      <h5>{t('businessform.fieldname.email')}: { business.Email }</h5>
      <h5>{t('businessform.fieldname.postalcode')}: { business.PostalCode }</h5>
      <h5>{t('businessform.fieldname.city')}: { business.City }</h5> 
      <h5>{t('businessform.fieldname.address')}: { business.Address }</h5>
      <h5>{t('businessform.fieldname.businessidcode')}: { business.BusinessIdCode }</h5>
      <h5>{t('businessform.fieldname.created')}: { business.CreatedAt }</h5>

      <br></br>
      <br></br>

      <h4>{t('businessform.fieldname.additionalinfo')}:</h4>
        <br></br>
        <div class="about-business-container">{ business.AdditionalInfo }</div>
        <br></br>
      </div>
      
      <br></br>
    </Container>
  );
};

export default BusinessInfo;