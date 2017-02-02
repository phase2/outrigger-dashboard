import { EventEmitter } from 'events'
import 'whatwg-fetch'

const store = new EventEmitter()

export default store

/**
 * Fetch dns records.
 */
store.fetchDnsRecords = () => {
  return fetch('/api/dnsrecords')
    .then(function(response) {
      console.log(response)
      return response.json()
    })
    .then(function(json) {
      console.log('parsed json', json)
      return store.processDnsRecords(json)
    })
    .catch(function(ex) {
      console.log('parsing failed', ex)
      return {};
  })
}

store.processDnsRecords = (records) => {
  let entries = []
  for (let prop in records) {
    if( records.hasOwnProperty(prop) ) {
      let record = records[prop]

      entries.push({
        ips: record.IPs,
        name: record.Name + '.' + record.Image + '.vm'
      })

      record.Aliases.forEach(alias => {
        entries.push({
          ips: record.IPs,
          name: alias
        })
      })
    }
  }
  return entries;
}