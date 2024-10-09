import { useState } from 'react'

const RegisterForm = ({ createNewUser }) => {
    const [newUsername, setNewUsername] = useState('')
    const [newPassword, setNewPassword] = useState('')
    const [newEmail, setNewEmail] = useState('')
    const [termsAccepted, setTermsAccepted] = useState(false)

    const addUser = event => {
        event.preventDefault()
        if (!termsAccepted) {
            alert('You must accept the terms')
            return
        }
        createNewUser({ 
            username: newUsername,
            password: newPassword,
            email: newEmail 
        })

        setNewUsername('')
        setNewPassword('')
        setNewEmail('')
        setTermsAccepted(false)
    }

    return (
        <div>
            <form onSubmit={addUser}>
                <div>
                  <label htmlFor="newUsernameInput">Username:</label>
                  <input id="newUsernameInput" value={newUsername} onChange={event => setNewUsername(event.target.value)} />
                </div>
                <div>
                  <label htmlFor="newPasswordInput">Password:</label>
                  <input type="password" id="newPasswordInput" value={newPassword} onChange={event => setNewPassword(event.target.value)} />
                </div>
                <div>
                  <label htmlFor="newEmailInput">Email:</label>
                  <input type="email" id="newEmailInput" value={newEmail} onChange={event => setNewEmail(event.target.value)} />
                </div>
                <div>
                  <input type='checkbox' id='termsCheckbox' checked={termsAccepted} onChange={() => setTermsAccepted(!termsAccepted)}/>
                  <label htmlFor="termsCheckbox">I agree to the <a href='/terms' target="_blank" rel="noopener noreferrer">terms and conditions</a></label>
                </div>
                <div>
                  <button id="createNewUserButton" type="submit">Register</button>
                </div>
            </form>
        </div>
    )
}

export default RegisterForm
