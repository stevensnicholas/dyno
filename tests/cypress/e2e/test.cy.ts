describe('end to end', () => {
  it('end to end', () => {
    cy.visit('/');
    cy.get('input[type="text"]')
      .type('val');
    cy.get('input[type="submit"]')
      .click();
    cy.get('h2').should(
      "have.text",
      "val"
    );
  });
});
