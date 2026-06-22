# i18n

The `i18n` provider plugin for [togo](https://github.com/togo-framework/togo).

## Install

```bash
togo install togo-framework/i18n
```

On import it self-registers with the kernel (priority-ordered provider). Access it
via the app container in your handlers/actions (e.g. `a.I18N`). Swap the default by
registering another provider for the same capability.
