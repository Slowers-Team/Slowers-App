import { useState } from 'react'
import { useTranslation } from 'react-i18next'

const RegisterForm = ({ createNewUser }) => {
    const [newUsername, setNewUsername] = useState('')
    const [newPassword, setNewPassword] = useState('')
    const [newEmail, setNewEmail] = useState('')
    const [newRole, setNewRole] = useState('')
    const [termsAccepted, setTermsAccepted] = useState(false)
    const [errorMessage, setErrorMessage] = useState('')
    const { t, i18n } = useTranslation()

    const addUser = async event => {
        event.preventDefault()
        if (!termsAccepted) {
          alert(t('error.acceptterms'))
          return
        }
        
        const userObject = {
          username: newUsername,
          password: newPassword,
          email: newEmail,
          role: newRole,
        }

        try {
          await createNewUser(userObject)
          setNewUsername('')
          setNewPassword('')
          setNewEmail('')
          setNewRole('')
          setTermsAccepted(false)
        } catch (error) {
          //täällä tapahtuu palvelimen lähettämän errorin näyttäminen
          //console.log(t('error.erroroccured'))
          setErrorMessage(t('error.erroroccured'))
        }
    }

    return (
        <div className='text-left'>
          {errorMessage && <p style={{ color: 'red'}}>{errorMessage}</p>}
            <form onSubmit={addUser}>
                <div className="input-group mb-4">
                  <span className="input-group-text">
                    <i className="bi bi-person-circle"></i>
                  </span>
                  <input id="newUsernameInput" value={newUsername} placeholder={t('user.input.username')} onChange={event => setNewUsername(event.target.value)} className='form-control' required/>
                </div>
                <div className='input-group mb-4'>
                  <span className="input-group-text">
                    <i className="bi bi-envelope-fill"></i>
                  </span>
                  <input type="email" id="newEmailInput" value={newEmail} placeholder={t('user.input.email')} onChange={event => setNewEmail(event.target.value)} className='form-control' required/>
                </div>
                <div className='input-group mb-4'>
                  <span className="input-group-text">
                    <i className="bi bi-lock-fill"></i>
                  </span>
                  <input type="password" id="newPasswordInput" value={newPassword} placeholder={t('user.input.password')} onChange={event => setNewPassword(event.target.value)} className='form-control' required/>
                </div>
                <div className='form-group mb-4'>
                    <label htmlFor="roleSelector">{t('label.defaultrole')}:</label>
                    <div>
                      <input type="radio" className='btn-check' name="roleSelector" id="growerSelector" autoComplete="off" value="grower" onChange={event => setNewRole(event.target.value)} required />
                      <label className='btn btn-light' htmlFor="growerSelector" style={{ marginRight: "0.5rem" }}>{t('role.grower')}</label>
                      <input type="radio" className='btn-check' name="roleSelector" id="retailerSelector" autoComplete='off' value="retailer" onChange={event => setNewRole(event.target.value)} />
                      <label className='btn btn-light' htmlFor="retailerSelector">{t('role.retailer')}</label>
                    </div>
                </div>
                <div className='form-check form-group mb-4'>
                  <input type='checkbox' className='form-check-input' id='termsCheckbox' checked={termsAccepted} onChange={() => setTermsAccepted(!termsAccepted)}/>
                  <label className='form-check-label' htmlFor="termsCheckbox">{t('label.iagreeto')} <a href='/terms' target="_blank" rel="noopener noreferrer">{t('label.terms')}</a></label>
                </div>
                <div>
                  <button type="submit" id="createNewUserButton" className='custom-button'>{t('button.register')}</button>
                </div>
            </form>
        </div>
    )
}

export default RegisterForm
