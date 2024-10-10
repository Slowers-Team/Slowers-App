import { useState } from 'react'

const RegisterForm = ({ createNewUser }) => {
    const [newUsername, setNewUsername] = useState('')
    const [newPassword, setNewPassword] = useState('')
    const [newEmail, setNewEmail] = useState('')
    const [newRole, setNewRole] = useState('')
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
            email: newEmail,
            role: newRole
        })

        setNewUsername('')
        setNewPassword('')
        setNewEmail('')
        setNewRole('')
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
                    <label htmlFor="roleSelector">Default role:</label>
                    <div>
                        <input type="radio" name="roleSelector" id="growerSelector" value="grower" onChange={event => setNewRole(event.target.value)}/>
                        <label htmlFor="growerSelector">Grower</label>
                    </div>
                    <div>
                        <input type="radio" name="roleSelector" id="retailerSelector" value="retailer" onChange={event => setNewRole(event.target.value)}/>
                        <label htmlFor="retailerSelector">Retailer</label>
                    </div>
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
