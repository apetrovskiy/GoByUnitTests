Feature: Wallet

    In order to use wallet
    As a user
    I want to be able to deposit and withdraw funds

    Background: wallet initial state
        Given there is 20 in wallet

    Scenario: deposit
        When I deposit 10
        Then amount is 30

    Scenario: withdrawal
        When I withdraw 10
        Then amount is 10
