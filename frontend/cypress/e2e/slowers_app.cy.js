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

  it('can login a user', function() {
    cy.register({username: 'testuser', email: 'test@email.com', password: 'testpassword', role: 'grower'})
    cy.visit('/login')
    cy.get('#emailInput').type('test@email.com')
    cy.get('#passwordInput').type('testpassword')
    cy.get('#loginButton').click()
    cy.contains('Homepage')
  })

  describe('when logged in', function() {
    beforeEach(function() {
      cy.registerAndLogin({username: 'testuser', email: 'test@email.com', password: 'testpassword', role: 'grower'})
    })

    it('changes role value on profile page when role switching button is clicked', function() {
      cy.visit('/user')
      cy.get('#roleValue').contains('Grower')
      cy.contains('Switch to Retailer').click()
      cy.get('#roleValue').contains('Retailer')
    })

    describe('when a site has been added', function() {
      beforeEach(function() {
        cy.visit('/grower/sites')
        cy.get('#addNewSiteButton').click()
        cy.get('#newSiteNameInput').type('Test site')
        cy.get('#newSiteNoteInput').type('Test note')
        cy.get('#saveNewSiteButton').click()
      })

      it('shows a site on the Sites tab after adding it', function() {
        cy.visit('/grower/sites')
        cy.contains('Test site')
        cy.contains('Test note')
      })

      it('does not show a site after deleting it', function() {
        cy.visit('/grower/sites')
        cy.contains('Test site').click()
        cy.get('#sitesTab').click()
        cy.get('#deleteSiteButton').click()
        cy.on('window:confirm', (confirmText) => {
          return true
        })

        cy.visit('/grower/sites')
        cy.contains('Test site').should('not.exist')
      })
    })
  })
})
