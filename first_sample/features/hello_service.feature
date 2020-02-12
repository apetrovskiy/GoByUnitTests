Feature: Hello

    Says Hello %username% in several languages

    Scenario Outline: Saying hello
        Given I greet user "<username>"
        When I say Hello in "<language>"
        Then the phrase is "<result phrase>"

        Examples:
            | username | language | result phrase   |
            | Chris    | English  | Hello, Chris    |
            | Elodie   | Spanish  | Hola, Elodie    |
            | Elodie   | French   | Bonjour, Elodie |
