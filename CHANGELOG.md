# Changelog

## [0.9.1](https://github.com/Nexters/book/compare/v0.9.0...v0.9.1) (2023-02-19)


### Bug Fixes

* import path ([24febb1](https://github.com/Nexters/book/commit/24febb131f3ad7a3adf88c4c7a376e5665130dc9))
* potential memory leak ([d57a919](https://github.com/Nexters/book/commit/d57a919bdd0d3cb5a90473caad17679cba088574))
* remove deferring ([c2db513](https://github.com/Nexters/book/commit/c2db5138532fde2fecc98e7edb864852181929ce))
* wrap error inside defer ([ac5baf6](https://github.com/Nexters/book/commit/ac5baf617c999953c3dd62b37c01145623c823c8))

## [0.9.0](https://github.com/Nexters/book/compare/v0.8.3...v0.9.0) (2023-02-11)


### Features

* 메모 수정 삭제 도입 ([638ccdd](https://github.com/Nexters/book/commit/638ccdd1b88b34ee8d8030b0f9b3d85a4b1deb4c))


### Bug Fixes

* returns total memos when user stat queries ([16792cc](https://github.com/Nexters/book/commit/16792cce59b13b558743e552314b34fd7e39c09d))
* 전체 메모 개수를 조회하게 user api 수정 ([98789ae](https://github.com/Nexters/book/commit/98789ae136a0914c1afa82e6e968e6f8962d634e))

## [0.8.3](https://github.com/Nexters/book/compare/v0.8.2...v0.8.3) (2023-02-09)


### Bug Fixes

* include new domains to cors allow origins ([467edd8](https://github.com/Nexters/book/commit/467edd80bcefb4dc2fa25f00c54c1d6e18ad8e00))

## [0.8.2](https://github.com/Nexters/book/compare/v0.8.1...v0.8.2) (2023-02-09)


### Bug Fixes

* return memos from books all api ([8c9d2ff](https://github.com/Nexters/book/commit/8c9d2ffb11576638199e9b7cc8da55a775e1d7e4))

## [0.8.1](https://github.com/Nexters/book/compare/v0.8.0...v0.8.1) (2023-02-08)


### Bug Fixes

* add charset and loc ([2fe67a0](https://github.com/Nexters/book/commit/2fe67a014e76784098693102bd40ad1331f12703))
* add parseTime true ([8d1712e](https://github.com/Nexters/book/commit/8d1712e2122b098c859939f6c591fbd18144938a))

## [0.8.0](https://github.com/Nexters/book/compare/v0.7.1...v0.8.0) (2023-02-08)


### Features

* implement book delete ([7f23ce8](https://github.com/Nexters/book/commit/7f23ce8105d13be7bf18c549f813fe85db44f376))
* implement book update ([8bad81f](https://github.com/Nexters/book/commit/8bad81f8725d76c5f22398c8a871bf947967e8db))
* implement user stat api ([fc64d35](https://github.com/Nexters/book/commit/fc64d3544293d5ec1e22346f9fa2e1b1a484e7a3))


### Bug Fixes

* ad allow credentials ([89e326f](https://github.com/Nexters/book/commit/89e326f81426215b57d0898ab6cad498047a3b4a))
* allow options method ([885cf7c](https://github.com/Nexters/book/commit/885cf7ccdd5ca44ec9e4001b9090113d7d741bdd))
* cors ([e3ef059](https://github.com/Nexters/book/commit/e3ef059e932fc441d9750ed1e5d036f2db3bc3c9))
* make explicit about allow headers ([99b9e31](https://github.com/Nexters/book/commit/99b9e31982de192fafcd3006e176a3b9f4572b49))
* unused var ([2a39f1c](https://github.com/Nexters/book/commit/2a39f1c1acaf0de0398a45d17f9ddb75bb0b39a2))
* update cors middleware ([7ab2d19](https://github.com/Nexters/book/commit/7ab2d19364ffd1a5e5c6a3326442f9f636b057d9))

## [0.7.1](https://github.com/Nexters/book/compare/v0.7.0...v0.7.1) (2023-02-02)


### Bug Fixes

* remove unused mysql declaration ([8356064](https://github.com/Nexters/book/commit/8356064e56d353073f9641bc02d7e42cde908003))

## [0.7.0](https://github.com/Nexters/book/compare/v0.6.0...v0.7.0) (2023-02-02)


### Features

* add count to book ([0120979](https://github.com/Nexters/book/commit/0120979b2d07be944fe046eecdfc0dab0d8ae28b))


### Bug Fixes

* book api ([2720b4b](https://github.com/Nexters/book/commit/2720b4b850ad170255fd2c8d514d25ac34aa1fe2))

## [0.6.0](https://github.com/Nexters/book/compare/v0.5.3...v0.6.0) (2023-01-31)


### Features

* add isReading flag ([7a866e7](https://github.com/Nexters/book/commit/7a866e7d49574c729758f05755095e99dcd3c10b))
* add max-len to memo text ([172563e](https://github.com/Nexters/book/commit/172563e2e79b09fddf84f86f2e9b9a902bdb33d8))

## [0.5.3](https://github.com/Nexters/book/compare/v0.5.2...v0.5.3) (2023-01-28)


### Bug Fixes

* api host 변경 ([cbf7f5f](https://github.com/Nexters/book/commit/cbf7f5fca5bc6d63d1b8d7abd50a1e475c3dbf30))
* 잘못된 swagger 문서 주석 수정정 ([f184d3a](https://github.com/Nexters/book/commit/f184d3a2f292d253ab3531ee6878a5407ad39f42))

## [0.5.2](https://github.com/Nexters/book/compare/v0.5.1...v0.5.2) (2023-01-28)


### Bug Fixes

* swagger endpoint mismatch fixed ([50a5e30](https://github.com/Nexters/book/commit/50a5e30e266810e1cf2a57476c3f4ca0b76d58c4))

## [0.5.1](https://github.com/Nexters/book/compare/v0.5.0...v0.5.1) (2023-01-25)


### Bug Fixes

* column property ([d5db8c5](https://github.com/Nexters/book/commit/d5db8c5537c16901dc275fc3bfa6309e6fe9e31d))

## [0.5.0](https://github.com/Nexters/book/compare/v0.4.0...v0.5.0) (2023-01-24)


### Features

* add swagger ([dba09e6](https://github.com/Nexters/book/commit/dba09e62709413b6ecc44a3c582933467e2a9cab))

## [0.4.0](https://github.com/Nexters/book/compare/v0.3.0...v0.4.0) (2023-01-24)


### Features

* add memo controller and boilerplate service ([ef4ff3f](https://github.com/Nexters/book/commit/ef4ff3f472e33d0c26d3487ce81b057b1e5affad))
* implement memo api ([d923f42](https://github.com/Nexters/book/commit/d923f425c85507da9efce5b356a212e4311a2351))

## [0.3.0](https://github.com/Nexters/book/compare/v0.2.0...v0.3.0) (2023-01-24)


### Features

* implement book api ([f541f18](https://github.com/Nexters/book/commit/f541f18a3e3aa27d85f81e505b105262dff9bf90))
* implement naver book search api ([7880387](https://github.com/Nexters/book/commit/7880387f73d5e55a61e0433dcfcbbcf162aa642c))
* search and create book ([aa13b13](https://github.com/Nexters/book/commit/aa13b139f109f98fd5147213ca17890f1f202612))

## [0.2.0](https://github.com/Nexters/book/compare/v0.1.0...v0.2.0) (2023-01-23)


### Features

* entity 및 user repository 추가 ([#3](https://github.com/Nexters/book/issues/3)) ([b39cdc0](https://github.com/Nexters/book/commit/b39cdc0c56945569d470e2c489d8bb7c92fcbe63))

## 0.1.0 (2023-01-20)


### Features

* add repository and entity ([4720670](https://github.com/Nexters/book/commit/47206702fccaadc420cbaf3ffaece92cbcf976cc))
* add sqlite ([08c70c6](https://github.com/Nexters/book/commit/08c70c6a07e6cb1faa7aea71bf47087414016f0e))
* implement SAST with securego/gosec ([b791ddb](https://github.com/Nexters/book/commit/b791ddbe381851f9f5cb04ccd41f96dbd20f180f))


### Bug Fixes

* add install script for gosec ([8ed2033](https://github.com/Nexters/book/commit/8ed20334da7c56034b6ba5d45be05c602ae67cf5))
* add personal access token ([3fb0b6f](https://github.com/Nexters/book/commit/3fb0b6fde35bd768b212f1778e0011df7b50b2b3))
* add setup go to SAST job ([64f320a](https://github.com/Nexters/book/commit/64f320a4cc4b780a042ab429fab42a90237fd7ff))
* handle errors ([7d726b7](https://github.com/Nexters/book/commit/7d726b76ec51dfb90a49bd6571ab818985b261f9))
* use securego/gosec to run action ([c58d54d](https://github.com/Nexters/book/commit/c58d54d87bdb63b6ace819c808f878fc036b0d40))


### Miscellaneous Chores

* release 0.1.0 ([62418f2](https://github.com/Nexters/book/commit/62418f220dede2c77d3bfa4c8808130065978789))

## [1.1.0](https://github.com/chaewonkong/go-template/compare/v1.0.0...v1.1.0) (2023-01-18)


### Features

* add sqlite ([08c70c6](https://github.com/chaewonkong/go-template/commit/08c70c6a07e6cb1faa7aea71bf47087414016f0e))

## 1.0.0 (2023-01-15)


### Features

* implement SAST with securego/gosec ([b791ddb](https://github.com/chaewonkong/go-template/commit/b791ddbe381851f9f5cb04ccd41f96dbd20f180f))


### Bug Fixes

* add install script for gosec ([8ed2033](https://github.com/chaewonkong/go-template/commit/8ed20334da7c56034b6ba5d45be05c602ae67cf5))
* add personal access token ([3fb0b6f](https://github.com/chaewonkong/go-template/commit/3fb0b6fde35bd768b212f1778e0011df7b50b2b3))
* add setup go to SAST job ([64f320a](https://github.com/chaewonkong/go-template/commit/64f320a4cc4b780a042ab429fab42a90237fd7ff))
* handle errors ([7d726b7](https://github.com/chaewonkong/go-template/commit/7d726b76ec51dfb90a49bd6571ab818985b261f9))
* use securego/gosec to run action ([c58d54d](https://github.com/chaewonkong/go-template/commit/c58d54d87bdb63b6ace819c808f878fc036b0d40))
