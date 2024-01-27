# How to Contribute

## Git Commit Message

A good git commit message should include action symbol and start with verb, e.g.

`* update form styles in registration page`

Action symbols includes:
- `+`: adding features
- `-`: removing features
- `*`: updating features
- `!`: fixing bugs
- `$`: refactoring codes
- `~`: chores, e.g. updating docs, dependencies, etc.

## Git Branch Naming

### Feature Branch

The branch name should be: `feat/{go/react}/{feature-name}`

e.g.

- `feat/go/add-login-service`
- `feat/react/update-form-styles-in-registration-page`

### Bug Fixing Branch

The branch name should be: `fix/{go/react}/{bug-name}`

e.g.

- `fix/go/unable-to-login`
- `fix/react/invalid-registration-form-validation`

### Code Refactoring Branch

The branch name should be: `refactor/{go/react}/{refactor-part}`

e.g.

- `refactor/go/rename-login-function`
- `refactor/react/extract-registration-form`

### Chore Branch

The branch name should be: `chore/{action}`

e.g.

- `chore/update-contributing-doc`
- `chore/update-go-dependencies`
