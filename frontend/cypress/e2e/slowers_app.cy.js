/*
 * WARNING: The code that follows may make you cry:
 *           A Safety Pig has been provided below for your benefit
 *                              _
 *      _._ _..._ .-',     _.._(`))
 *     '-. `     '  /-._.-'    ',/
 *        )         \            '.
 *       / _    _    |             \
 *      |  a    a    /              |
 *      \   .-.                     ;
 *       '-('' ).-'       ,'       ;
 *          '-;           |      .'
 *             \           \    /
 *             | 7  .__  _.-\   \
 *             | |  |  ``/  /`  /
 *            /,_|  |   /,_/   /
 *               /,_/      '`-'
 */


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

  it('can change language', function() {
    cy.visit('')
    cy.contains('Log into Slowers')
    cy.get('#languageButton').click()
    cy.contains('🇬🇧 English')
    cy.contains('🇫🇮 Suomi')
    cy.get('#fi').click()
    cy.contains('Kirjaudu Slowersiin')
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

  it('cannot register with existing email address', function() {
    cy.register({username: 'testuser', email: 'test@email.com', password: 'testpassword', role: 'grower'})
    cy.visit('/register')
    cy.get('#newUsernameInput').type('testuser2')
    cy.get('#newEmailInput').type('test@email.com')
    cy.get('#newPasswordInput').type('testpassword2')
    cy.contains('Retailer').click()
    cy.get('#termsCheckbox').check()
    cy.get('#createNewUserButton').click()
    cy.on('window:alert',(t)=>{
      expect(t).to.contains('Error: Email already exists');
   })
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

  describe('when logged in', function() {
    beforeEach(function() {
      cy.registerAndLogin({username: 'testuser', email: 'test@email.com', password: 'testpassword', role: 'grower'})
      cy.contains('Homepage')
    })

    it('changes role value on profile page when role switching button is clicked', function() {
      cy.visit('/user')
      cy.contains('test@email.com')
      cy.get('#roleValue').contains('Grower')
      cy.contains('Switch to Retailer').click()
      cy.get('#roleValue').contains('Retailer')
    })

    describe('when a site has been added', function() {
      beforeEach(function() {
        cy.visit('/grower')
        cy.get('#addNewSiteButton').click()
        cy.get('#newSiteNameInput').type('Test site')
        cy.get('#newSiteNoteInput').type('Test note')
        cy.get('#saveNewSiteButton').click()

        cy.visit('/grower')
        cy.get('#addNewSiteButton').click()
        cy.get('#newSiteNameInput').type('Greenhouse')
        cy.get('#newSiteNoteInput').type('Something')
        cy.get('#saveNewSiteButton').click()

        cy.visit('/grower')
        cy.contains('Test site').click()
        cy.get('#homeTab').click()
        cy.get('#addNewSiteButton').click()
        cy.get('#newSiteNameInput').type('Field')
        cy.get('#newSiteNoteInput').type('Stuff')
        cy.get('#saveNewSiteButton').click()
      })

      it('shows a site on the Home tab after adding it', function() {
        cy.visit('/grower')
        cy.contains('Test site')
        cy.contains('Test note')
      })

      it('does not show a site after deleting it', function() {
        cy.visit('/grower')
        cy.contains('Test site').click()
        cy.get('#homeTab').click()
        cy.contains('Field')
        cy.get('#deleteSiteButton').click()
        cy.on('window:confirm', (confirmText) => {
          return true
        })

        cy.visit('/grower')
        cy.contains('Greenhouse')
        cy.contains('Test site').should('not.exist')
      })

      describe('when a flower has been added', function() {
        beforeEach(function() {
          cy.visit('/grower')
          cy.contains('Test site').click()
          cy.contains('Test site homepage')
          cy.get('#flowersTab').click()
          cy.get('#showFlowerAddingFormButton').click()
          cy.get('#newFlowerNameInput').type('Test flower')
          cy.get('#newFlowerLatinNameInput').type('Test latin name')
          cy.get('#newFlowerQtyInput').clear()
          cy.get('#newFlowerQtyInput').type('10')
          cy.get('#saveNewFlowerButton').click()
        })

        it('can add flower to a site', function() {
          cy.contains('Test flower')
        })

        it('can delete flower from a site', function() {
          cy.visit('/grower')
          cy.contains('Test site').click()
          cy.contains('Test site homepage')
          cy.get('#flowersTab').click()
          cy.contains('Test flower').click()
          cy.get('#deleteFlowerButton').click()
          cy.on('window:confirm', (confirmText) => {
            return true
          })
          cy.contains('Test flower').should('not.exist')
        })

        it('does not show flower without an image on retailer flower page', function() {
          cy.visit('/retailer/flowers')
          cy.contains('Test flower').should('not.exist')
        })
      })
    })
  })
})
