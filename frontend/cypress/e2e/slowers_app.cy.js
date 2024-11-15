describe('Slowers ', function() {
  beforeEach(function() {
    cy.request('GET', '/api/reset')
  })
  it('front page can be opened', function() {
    cy.visit('')
    cy.contains('Slowers')
  })

  it('register page can be opened', function() {
    cy.visit('/register')
    cy.contains('Slowers registration')
  })

  it('user can register', function() {
    cy.visit('/register')
    cy.get('#newUsernameInput').type('testuser')
    cy.get('#newEmailInput').type('test@email.com')
    cy.get('#newPasswordInput').type('testpassword')
    cy.contains('Grower').click()
    cy.get('#termsCheckbox').check()
    cy.get('#createNewUserButton').click()
  })

  it('user can login', function() {
    cy.visit('/login')
    cy.get('#emailInput').type('test@email.com')
    cy.get('#passwordInput').type('testpassword')
    cy.get('#loginButton').click()
    cy.contains('Homepage')
  })
})