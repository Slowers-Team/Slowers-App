describe('Slowers ', function() {
  beforeEach(function() {
    cy.request('GET', '/api/reset')
  })
  after(function() {
    cy.request('GET', '/api/reset')
  })
  describe('when logged in as a retailer', function() {
    beforeEach(function() {
      cy.registerAndLogin({username: 'testuser', email: 'test@email.com', password: 'testpassword', role: 'retailer'})
      cy.contains('Welcome to Slowers')
    })

    describe('page visibility', function() {
      beforeEach(function(){
        cy.visit('/user')
        cy.contains('test@email.com')
      })
      // it('retailer user can not access grower page', function() {
      //   cy.contains('Retailer')
      //   cy.visit('grower')
      //   cy.contains('Welcome to Slowers')
      // })
      // it('retailer can see authorized pages in side menu', function() {
      //   cy.contains('Retailer')
      //   cy.visit('home')
      //   cy.get('#offcanvasButton').click()
      //   cy.contains('Home')
      //   cy.contains('Marketplace')
      //   cy.contains('Retailer')
      //   cy.contains('Terms')
      // })	
    })
    describe('business page visible', function() {
      beforeEach(function(){
        cy.visit('/business')
      })
      it('contains business form before business form is created', function() {
        cy.contains('Create a business')
      })
    })
    describe('when business exists', function() {
      beforeEach(function() {
        cy.visit('/business')
        cy.get('#businessNameInput').type('Test business')
        cy.get('#businessIdCodeInput').type('1234567-8')
        cy.get('#retailerSelector').check({ force: true })
        cy.get('#businessPhoneNumberInput').type('0400123456')
        cy.get('#businessEmailInput').type('testi@email.com')
        cy.get('#businessAddressInput').type('Testikuja 1 A')
        cy.get('#businessPostalCodeInput').type('00100')
        cy.get('#businessCityInput').type('Helsinki')
        cy.get('#businessAddInfoInput').type('This is a test business')
        cy.get('#businessFormSubmit').click()
      })

      it('displays correct business information', function() {
        cy.visit('/business')
        cy.contains('Test business')
        cy.contains('1234567-8')
        cy.contains('0400123456')
        cy.contains('testi@email.com')
        cy.contains('Testikuja 1 A')
        cy.contains('00100')
        cy.contains('Helsinki')
        cy.contains('This is a test business')
      })
    })
  })
  describe('when logged in as a grower', function() {
    beforeEach(function() {
      cy.registerAndLogin({username: 'testuser', email: 'test@email.com', password: 'testpassword', role: 'grower'})
      cy.contains('Welcome to Slowers')
    })

    describe('page visibility', function() {
      beforeEach(function() {
        cy.visit('/user')
        cy.contains('test@email.com')
      })
      // it('grower user can not access retailer page', function() {
      //   cy.contains('Grower')
      //   cy.visit('retailer')
      //   cy.contains('Welcome to Slowers')
      // })
      // it('grower user can not access business owner page', function() {
      //   cy.contains('Grower')
      //   cy.visit('business_owner')
      //   cy.contains('Welcome to Slowers')
      // })		
      // it('grower can see authorized pages in side menu', function() {
      //   cy.contains('Grower')
      //   cy.visit('home')
      //   cy.get('#offcanvasButton').click()
      //   cy.contains('Home')
      //   cy.contains('Marketplace')
      //   cy.contains('Grower page')
      //   cy.contains('Terms')
      // })
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