import { useTranslation } from "react-i18next"


const BusinessInfo = ({ business }) => {
  const { t, i18n } = useTranslation()
  const dateOnly = business?.CreatedAt ? new Date(business.CreatedAt).toISOString().split("T")[0] : "N/A";

  return (
    <div className="container">
      <div className="row justify-content-center">
        <div className="col-12 col-md-12 col-lg-12 col-xl-8">
          <div className="card" style={{ borderRadius: "1rem" }}>
            <div className="card-body p-5">
              <h2>{ business.BusinessName }</h2>
              <br/>
              <h5>{t('businessform.fieldname.businesstype')}: { business.Type }</h5>
              <h5>{t('businessform.fieldname.phonenumber')}: { business.PhoneNumber }</h5>
              <h5>{t('businessform.fieldname.email')}: { business.Email }</h5>
              <h5>{t('businessform.fieldname.postalcode')}: { business.PostalCode }</h5>
              <h5>{t('businessform.fieldname.city')}: { business.City }</h5>
              <h5>{t('businessform.fieldname.address')}: { business.Address }</h5>
              <h5>{t('businessform.fieldname.businessidcode')}: { business.BusinessIdCode }</h5>
              <h5>{t('businessform.fieldname.created')}: { dateOnly }</h5>
              <br/>
              <h4>{t('businessform.fieldname.additionalinfo')}:</h4>
              <h5>{ business.AdditionalInfo }</h5>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default BusinessInfo