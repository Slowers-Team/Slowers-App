import { useState } from "react";
import { useTranslation } from "react-i18next";
import { validateEmail, validateBusinessIdCode, validatePostalCode, validatePhoneNumber } from "../utils";


const CreateBusinessForm = ({ createNewBusiness }) => {
  const { t, i18n } = useTranslation()
  const [businessName, setBusinessName] = useState('')
  const [businessIdCode, setBusinessIdCode] = useState('')
  const [type, setType] = useState ('')
  const [phoneNumber, setPhoneNumber] = useState('')
  const [email, setEmail] = useState('')
  const [additionalInfo, setAdditionalInfo] = useState('')
  const [address, setAddress] = useState('')
  const [postalCode, setPostalCode] = useState('')
  const [city, setCity] = useState('')
  const [delivery, setDelivery] = useState('no')
  const [errorMessage, setErrorMessage] = useState('')


  const validateBusiness = (businessObject) => {
    if (!validateBusinessIdCode(businessObject.businessIdCode)) {
      return (t('error.erroroccured'))
    }
    if (!validateEmail(businessObject.email)) {
      return (t('error.erroroccured'))
    }
    if (!validatePostalCode(businessObject.postalCode)) {
      return (t('error.erroroccured'))
    }
    if (!validatePhoneNumber(businessObject.phoneNumber)) {
      return (t('error.erroroccured'))
    }
    return 'ok'
  }

  const createBusiness = async (event) => {
    event.preventDefault()

    const businessObject = {
      businessName,
      businessIdCode,
      type,
      phoneNumber,
      email,
      additionalInfo,
      address,
      postalCode,
      city,
      delivery
    }

    setErrorMessage('')
    const validationResult = validateBusiness(businessObject)
    if (validationResult != 'ok') {
      setErrorMessage(validationResult)
      return
    }

    try {
      await createNewBusiness(businessObject)
      setBusinessName('')
      setBusinessIdCode('')
      setType('')
      setPhoneNumber('')
      setEmail('')
      setAdditionalInfo('')
      setAddress('')
      setPostalCode('')
      setCity('')
      setDelivery('')
    } catch (error) {
        setErrorMessage(t('error.erroroccured'))
    }
  }

  return (
    <div>
      <h2>{t('title.createbusiness')}</h2>
      <p>
        {t('businessform.instructions.note')}
      </p>
      <div>
        <form onSubmit={createBusiness}>
          <div>
            <table>
              <tbody>
                <tr>
                  <td>{t('businessform.fieldname.businessname')}</td>
                  <td>
                    <input 
                      className="form-control"
                      type="text"
                      value={businessName}
                      placeholder={t('businessform.input.businessname')}
                      onChange={event => setBusinessName(event.target.value)}
                      required
                    />
                  </td>
                </tr>
                <tr>
                  <td>{t('businessform.fieldname.businessidcode')}</td>
                  <td>
                    <input
                      className="form-control"
                      type="text"
                      value={businessIdCode}
                      maxLength={9}
                      placeholder={t('businessform.input.businessidcode')}
                      onChange={event => setBusinessIdCode(event.target.value)}
                      required
                    />
                  </td>
                </tr>
                <tr>
                  <td>{t('businessform.fieldname.businesstype')}</td>
                  <td>
                    <input
                      type="radio"
                      id="growerSelector"
                      className='btn-check'
                      name="typeSelector"
                      value="grower"
                      checked={type === "grower"}
                      onChange={event => {
                        setType(event.target.value)
                        setDelivery("no")
                      }}
                      required
                    />
                    <label
                      className='btn btn-outline-secondary'
                      style={{ marginRight: "0.5rem" }}
                      htmlFor="growerSelector" >
                      {t('button.grower')}
                    </label>
                    <input
                      type="radio"
                      id="retailerSelector"
                      className='btn-check'
                      name="typeSelector"
                      value="retailer"
                      checked={type === "retailer"}
                      onChange={event => {
                        setType(event.target.value)
                        setDelivery("no")
                      }}
                    />
                    <label
                      className='btn btn-outline-secondary'
                      style={{ marginRight: "0.5rem" }}
                      htmlFor="retailerSelector" >
                      {t('button.retailer')}
                    </label>
                  </td>
                </tr>
                <tr>
                  <td>{t('businessform.fieldname.phonenumber')}</td>
                  <td>
                    <input
                      className="form-control"
                      type="tel"
                      value={phoneNumber}
                      placeholder={t('businessform.input.phonenumber')}
                      minLength={10}
                      maxLength={13}
                      onChange={event => setPhoneNumber(event.target.value)}
                      required
                    />
                  </td>
                </tr>
                <tr>
                  <td>{t('businessform.fieldname.email')}</td>
                  <td>
                    <input
                      className="form-control"
                      type="email"
                      value={email}
                      placeholder={t('businessform.input.email')}
                      onChange={event => setEmail(event.target.value)}
                      required
                    />
                  </td>
                </tr>
                <tr>
                  <td style={{verticalAlign: "top"}}>{t('businessform.fieldname.address')}</td>
                  <td>
                    <input
                      className="form-control"
                      type="text"
                      value={address}
                      placeholder={t('businessform.input.address')}
                      onChange={event => setAddress(event.target.value)}
                    />
                    <small className="w-100">{t('businessform.instructions.address')}</small>
                  </td>
                </tr>
                <tr>
                  <td>{t('businessform.fieldname.postalcode')}</td>
                  <td>
                    <input 
                      className="form-control"
                      type="tel"
                      value={postalCode}
                      minLength={5}
                      maxLength={5}
                      placeholder={t('businessform.input.postalcode')}
                      onChange={event => setPostalCode(event.target.value)}
                      style={{ width: "200px"}}
                      required
                    />
                  </td>
                </tr>
                <tr>
                  <td>{t('businessform.fieldname.city')}</td>
                  <td>
                    <input
                      className="form-control"
                      type="text"
                      value={city}
                      placeholder={t('businessform.input.city')}
                      onChange={event => setCity(event.target.value)}
                      required
                    />
                  </td>
                </tr>
                <tr>
                  <td style={{verticalAlign: "top"}}>{t('businessform.fieldname.additionalinfo')}</td>
                  <td>
                    <textarea
                      className="form-control"
                      type="text"
                      value={additionalInfo}
                      placeholder={t('businessform.input.additionalinfo')}
                      rows={5}
                      maxLength={1500}
                      onChange={event => setAdditionalInfo(event.target.value)}
                      required>
                    </textarea>
                    <small className="w-100">{t('businessform.instructions.additionalinfo')}</small>
                  </td>
                </tr>
                {type === "grower" && (
                  <tr>
                    <td>{t('businessform.fieldname.delivery')}</td>
                    <td>
                      <input
                        type="radio"
                        id="yesSelector"
                        className='btn-check'
                        name="deliverySelector"
                        value="yes"
                        checked={delivery === "yes"}
                        onChange={event => setDelivery(event.target.value)}
                        required
                      />
                      <label
                        className='btn btn-outline-secondary'
                        style={{ marginRight: "0.5rem" }}
                        htmlFor="yesSelector" >
                        {t('businessform.input.yesdelivery')}
                      </label>
                      <input
                        type="radio"
                        id="noSelector"
                        className='btn-check'
                        name="deliverySelector"
                        value="no"
                        checked={delivery === "no"}
                        onChange={event => setDelivery(event.target.value)}
                      />
                      <label
                        className='btn btn-outline-secondary'
                        style={{ marginRight: "0.5rem" }}
                        htmlFor="noSelector" >
                        {t('businessform.input.nodelivery')}
                      </label>
                    </td>
                  </tr>
                )}
              <tr>
                <td>
                <button type="submit" className='custom-button'>{t('button.createbusiness')}</button>
                </td>
                <td>
                  {errorMessage && <p style={{ color: 'red'}}>{errorMessage}</p>}
                </td>
              </tr>
              </tbody>
            </table>
          </div>
        </form>
      </div> 
    </div>
  )
}

export default CreateBusinessForm
