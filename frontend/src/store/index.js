import { EventEmitter } from 'events'
import 'whatwg-fetch'

const store = new EventEmitter()

export default store

/**
 * Fetch dns records.
 */

store.fetchDnsRecords = () => {
  fetch('http://dnsdock.devtools.vm/services')
    .then(function(response) {
      console.log(response)
      return response.json()
    })
    .then(function(json) {
      console.log('parsed json', json)
    })
    .catch(function(ex) {
      console.log('parsing failed', ex)
  })
}