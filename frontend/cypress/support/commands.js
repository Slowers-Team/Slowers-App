Cypress.Commands.add('register' , ( {username, email, password, role} ) => {
    cy.request('POST', '/api/register', {
        username,
        email,
        password,
        role
    })
})

Cypress.Commands.add('login' , ( {email, password} ) => {
    cy.visit('/login')
    cy.get('#emailInput').type(email)
    cy.get('#passwordInput').type(password)
    cy.get('#loginButton').click()
})

Cypress.Commands.add('registerAndLogin' , ( {username, email, password, role} ) => {
    cy.register({username, email, password, role})
    cy.login({email, password})
})