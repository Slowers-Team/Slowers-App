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
    cy.contains('Login')
    cy.contains('Register')
    cy.contains('Terms')
  })

  it('can change language', function() {
    cy.visit('')
    cy.contains('Log into Slowers')
    cy.get('#languageButton').click()
    cy.contains('🇬🇧 English')
    cy.contains('🇫🇮 Suomi')
    cy.contains('🇸🇪 Svenska')
    cy.get('#fi').click()
    cy.contains('Kirjaudu Slowersiin')
    cy.get('#languageButton').click()
    cy.get('#sv').click()
    cy.contains('Logga in i Slowers')
  })
})