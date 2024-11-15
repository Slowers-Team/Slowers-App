describe('Slowers ', function() {
  beforeEach(function() {
    cy.request('GET', '/api/reset')
  })
  after(function() {
    cy.request('GET', '/api/reset')
  })
  it('can open front page', function() {
    cy.visit('')
    cy.contains('Slowers')
  })

  it('can open register page', function() {
    cy.visit('/register')
    cy.contains('Slowers registration')
  })

  it('can open terms page when not logged in', function() {
    cy.visit('/terms')
    cy.contains('Terms and Conditions')
  })

  it('can open side menu by clicking hamburger button', function() {
    cy.visit('')
    cy.get('#offcanvasButton').click()
    cy.contains('Home')
    cy.contains('Login')
    cy.contains('Register')
    cy.contains('Terms')
  })

  it('can register a new user', function() {
    cy.visit('/register')
    cy.get('#newUsernameInput').type('testuser')
    cy.get('#newEmailInput').type('test@email.com')
    cy.get('#newPasswordInput').type('testpassword')
    cy.contains('Grower').click()
    cy.get('#termsCheckbox').check()
    cy.get('#createNewUserButton').click()
    cy.contains('Homepage')
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

  it('can login a user', function() {
    cy.register({username: 'testuser', email: 'test@email.com', password: 'testpassword', role: 'grower'})
    cy.visit('/login')
    cy.get('#emailInput').type('test@email.com')
    cy.get('#passwordInput').type('testpassword')
    cy.get('#loginButton').click()
    cy.contains('Homepage')
  })

  it('cannot login with incorrect credentials', function() {
    cy.register({username: 'testuser', email: 'test@email.com', password: 'testpassword', role: 'grower'})
    cy.visit('/login')
    cy.get('#emailInput').type('test@email.com')
    cy.get('#passwordInput').type('wrongpassword')
    cy.get('#loginButton').click()
    cy.contains('An error occurred. Please try again.')
  })
})