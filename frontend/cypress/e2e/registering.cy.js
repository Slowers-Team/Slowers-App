describe('Slowers ', function() {
    beforeEach(function() {
      cy.request('GET', '/api/reset')
    })
    after(function() {
      cy.request('GET', '/api/reset')
    })

  it('can register a new user', function() {
    cy.visit('/register')
    cy.get('#newUsernameInput').type('testuser')
    cy.get('#newEmailInput').type('test@email.com')
    cy.get('#newPasswordInput').type('testpassword')
    cy.contains('Grower').click()
    cy.get('#termsCheckbox').check()
    cy.get('#createNewUserButton').click()
    cy.contains('Log into Slowers')
  })

  it('cannot register with existing email address', function() {
    cy.register({username: 'testuser', email: 'test@email.com', password: 'testpassword', role: 'grower'})
    cy.visit('/register')
    cy.get('#newUsernameInput').type('testuser2')
    cy.get('#newEmailInput').type('test@email.com')
    cy.get('#newPasswordInput').type('testpassword2')
    cy.contains('Retailer').click()
    cy.get('#termsCheckbox').check()
    cy.get('#createNewUserButton').click()
    cy.contains('An error occurred. Please try again.')
   })
  

  it('cannot register with malformatted email address', function() {
    cy.visit('/register')
    cy.get('#newUsernameInput').type('testuser')
    cy.get('#newEmailInput').type('testemail')
    cy.get('#newPasswordInput').type('testpassword')
    cy.contains('Grower').click()
    cy.get('#termsCheckbox').check()
    cy.get('#createNewUserButton').click()
    cy.get('input:invalid').should('have.length', 1)
  })

  it('cannot register without accepting terms', function() {
    cy.visit('/register')
    cy.get('#newUsernameInput').type('testuser')
    cy.get('#newEmailInput').type('test@email.com')
    cy.get('#newPasswordInput').type('testpassword')
    cy.contains('Grower').click()
    cy.get('#createNewUserButton').click()
    cy.on('window:alert',(t)=>{
      expect(t).to.contains('You must accept the terms');
   })
  })

  it('cannot register with an empty field', function() {
    cy.visit('/register')
    cy.get('#newEmailInput').type('test@email.com')
    cy.get('#newPasswordInput').type('testpassword')
    cy.contains('Grower').click()
    cy.get('#termsCheckbox').check()
    cy.get('#createNewUserButton').click()
    cy.get('input:invalid').should('have.length', 1)
  })
})