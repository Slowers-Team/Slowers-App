describe('Slowers ', function() {
    beforeEach(function() {
      cy.request('GET', '/api/reset')
			cy.register({username: 'testuser', email: 'test@email.com', password: 'testpassword', role: 'grower'})
    })
    after(function() {
      cy.request('GET', '/api/reset')
    })

  it('can login a user', function() {
    cy.visit('/login')
    cy.get('#emailInput').type('test@email.com')
    cy.get('#passwordInput').type('testpassword')
    cy.get('#loginButton').click()
    cy.contains('Homepage')
  })

  it('cannot login with incorrect credentials', function() {
    cy.visit('/login')
    cy.get('#emailInput').type('test@email.com')
    cy.get('#passwordInput').type('wrongpassword')
    cy.get('#loginButton').click()
    cy.contains('Invalid email or password')
  })
})
