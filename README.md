# go-couchdb-query-engine
CouchDB query implementation in Golang for in-memory object querying

**Build status**: [![Build Status](https://travis-ci.org/wearetheledger/go-couchdb-query-engine.svg?branch=master)](https://travis-ci.org/wearetheledger/go-couchdb-query-engine)

## Features
- skip
- limit
- selector

## Selector
The selector query accepts following couchdb operators:
- all
- and
- elematch
- eq
- exists
- gt
- gte
- in
- lt
- lte
- mod
- ne
- nin
- nor
- not
- or
- regex
- size
- type
