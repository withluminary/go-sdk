# Changelog

## 0.4.0 (2026-03-28)

Full Changelog: [v0.3.0...v0.4.0](https://github.com/withluminary/go-sdk/compare/v0.3.0...v0.4.0)

### Features

* **internal:** support comma format in multipart form encoding ([a1e7304](https://github.com/withluminary/go-sdk/commit/a1e73041f22085b2831dd0d12768298d76101680))


### Bug Fixes

* prevent duplicate ? in query params ([d04f903](https://github.com/withluminary/go-sdk/commit/d04f9039bbedf89989498cb93c8572f339512f65))
* **tests:** prevent tests failing due to making unnecessary OAuth token requests ([4b2b000](https://github.com/withluminary/go-sdk/commit/4b2b000e6cb1a1d680fee61a45bbff2e0605ca81))


### Chores

* **ci:** skip lint on metadata-only changes ([20b7fb1](https://github.com/withluminary/go-sdk/commit/20b7fb19fe5c7d77b2e0338d86710f5bf0686170))
* **ci:** skip uploading artifacts on stainless-internal branches ([1cd1dfb](https://github.com/withluminary/go-sdk/commit/1cd1dfb29fc909d827b2d682edc07f8df3855d55))
* **ci:** support opting out of skipping builds on metadata-only commits ([2b5595c](https://github.com/withluminary/go-sdk/commit/2b5595c2d8ce869a08b688e0d94abbc7362af06d))
* **client:** fix multipart serialisation of Default() fields ([51c2566](https://github.com/withluminary/go-sdk/commit/51c2566a3ed9bd6250014b37246994df5c36f074))
* **internal:** codegen related update ([d052c29](https://github.com/withluminary/go-sdk/commit/d052c29016722ef940f5838b8b78cb8f2d4e6363))
* **internal:** minor cleanup ([6cbe3d3](https://github.com/withluminary/go-sdk/commit/6cbe3d37fafa1512c5abdd990b46762d17dbe4e2))
* **internal:** support default value struct tag ([9dd077a](https://github.com/withluminary/go-sdk/commit/9dd077ab66f3a747ae045c895bc64050b52dc605))
* **internal:** tweak CI branches ([41888c5](https://github.com/withluminary/go-sdk/commit/41888c52ba2c6847e25111b7478b6d21b1b47fcd))
* **internal:** update gitignore ([1f5298b](https://github.com/withluminary/go-sdk/commit/1f5298bed195884d48cb1572ce9319dd1d45f6d7))
* **internal:** use explicit returns ([54eaa5a](https://github.com/withluminary/go-sdk/commit/54eaa5ad2b0db6d1055ccd6e5288d9e46bc8a556))
* **internal:** use explicit returns in more places ([8fcf61b](https://github.com/withluminary/go-sdk/commit/8fcf61b8f3c5268476aa0c25503b85d65670cb90))
* remove unnecessary error check for url parsing ([9304cc8](https://github.com/withluminary/go-sdk/commit/9304cc838cfc671291563222fb20d80b350cfeed))
* update docs for api:"required" ([017a902](https://github.com/withluminary/go-sdk/commit/017a902c3704db9479bf24bedfbc580a6e49c091))
* update placeholder string ([093b778](https://github.com/withluminary/go-sdk/commit/093b778660db9656a4f45431da42c2bbc84da869))

## 0.3.0 (2026-03-04)

Full Changelog: [v0.2.0...v0.3.0](https://github.com/withluminary/go-sdk/compare/v0.2.0...v0.3.0)

### Features

* **client:** add a convenient param.SetJSON helper ([d9c5d6e](https://github.com/withluminary/go-sdk/commit/d9c5d6efb926e6edf1b0bf273f7749a0722fa389))


### Bug Fixes

* allow canceling a request while it is waiting to retry ([0d4d4ea](https://github.com/withluminary/go-sdk/commit/0d4d4eaffe3b62719f83500bbcafabb13b2c1e9c))
* **docs:** add missing pointer prefix to api.md return types ([1bab893](https://github.com/withluminary/go-sdk/commit/1bab893822a6dd6ffa53391b8ef24ef2413c249c))
* **encoder:** correctly serialize NullStruct ([491585c](https://github.com/withluminary/go-sdk/commit/491585ca245d20382c2aa4c5abd3b022e55fb9d7))
* **internal:** skip tests that depend on mock server ([ff40f4b](https://github.com/withluminary/go-sdk/commit/ff40f4b3e2059e3b03db07e1969f3c37bae2b49d))
* move oauth2 grant type to body of request ([542a83f](https://github.com/withluminary/go-sdk/commit/542a83fe1517ed3c6dce9635c491bfede098320f))
* various public API fixes ([5cdc43a](https://github.com/withluminary/go-sdk/commit/5cdc43a48c0f291ab63c3422b2324142bedc7efb))


### Chores

* add entity in-estate status to API ([da9f477](https://github.com/withluminary/go-sdk/commit/da9f4770fbaa800e68c8e829e322b2bd9e79c1e3))
* **internal:** codegen related update ([90e61cf](https://github.com/withluminary/go-sdk/commit/90e61cfbcd12aeb5e2be8bb5ec7d5b45cce7666c))
* **internal:** move custom custom `json` tags to `api` ([4dd9718](https://github.com/withluminary/go-sdk/commit/4dd9718d29f9c2884fdb7353c4e2484fec729603))
* **internal:** remove mock server code ([0598a24](https://github.com/withluminary/go-sdk/commit/0598a244fc8cf8e52b74228ec69159f75c350e78))
* **internal:** update `actions/checkout` version ([319cd2b](https://github.com/withluminary/go-sdk/commit/319cd2b7e137243fd58903a2788c4644e313bef3))
* remove custom code ([74977aa](https://github.com/withluminary/go-sdk/commit/74977aa136959e16642e94a7c9f591b4d2b64dcb))
* update mock server docs ([2b404b4](https://github.com/withluminary/go-sdk/commit/2b404b4d17bf8f50280c3c50a85c02f6ece732e4))

## 0.2.0 (2026-01-07)

Full Changelog: [v0.1.0...v0.2.0](https://github.com/withluminary/go-sdk/compare/v0.1.0...v0.2.0)

### Features

* **api:** revert token change ([b942c35](https://github.com/withluminary/go-sdk/commit/b942c35cfd651c1898d79d1aec70b548595990f5))

## 0.1.0 (2026-01-07)

Full Changelog: [v0.0.1...v0.1.0](https://github.com/withluminary/go-sdk/compare/v0.0.1...v0.1.0)

### Features

* **api:** add cursor config ([49882be](https://github.com/withluminary/go-sdk/commit/49882becab0c83133b8caac8641fe1d07a1472ad))
* **api:** configurable auth domain ([3103d48](https://github.com/withluminary/go-sdk/commit/3103d485c26e90023257cc0513c0666939694600))
* **api:** manual updates ([6cd2169](https://github.com/withluminary/go-sdk/commit/6cd216947413c9e86a6b1cf536bf4b57ecc95440))


### Chores

* configure new SDK language ([5d5c72d](https://github.com/withluminary/go-sdk/commit/5d5c72dcb168e1173b5e0148049c7aebbe9935d2))
* update SDK settings ([ae097f8](https://github.com/withluminary/go-sdk/commit/ae097f8f4af60e1584b48f7014cabebeb3413ca0))
* user endpoint and pagination ([bf142ea](https://github.com/withluminary/go-sdk/commit/bf142ead541d8cf6b26692d36bd6b4b97604090f))
